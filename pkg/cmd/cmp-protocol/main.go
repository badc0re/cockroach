package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"reflect"

	"github.com/cockroachdb/cockroach/pkg/sql/pgwire/pgwirebase"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/jackc/pgx/pgproto3"
	"github.com/pkg/errors"
)

var (
	pgAddr = flag.String("pg", "localhost:5432", "postgres address")
	pgUser = flag.String("pg-user", "postgres", "postgres user")
	crAddr = flag.String("cr", "localhost:26257", "cockroach address")
	crUser = flag.String("cr-user", "root", "cockroach user")
)

func main() {
	flag.Parse()

	rng := rand.New(rand.NewSource(0)) // time.Now().Unix()))

	for {
		typ := sqlbase.RandColumnType(rng)
		sem := typ.SemanticType
		switch sem {
		case sqlbase.ColumnType_COLLATEDSTRING,
			sqlbase.ColumnType_BIT,
			sqlbase.ColumnType_DECIMAL,
			sqlbase.ColumnType_INTERVAL,
			sqlbase.ColumnType_ARRAY,
			sqlbase.ColumnType_TUPLE:
			continue
		}
		datum := sqlbase.RandDatum(rng, typ, true)
		input := fmt.Sprintf("SELECT ARRAY[%s]", datum)
		fmt.Print("\n------\n", input, "\n\n")
		if err := compare(os.Stdout, input, *pgAddr, *crAddr, *pgUser, *crUser); err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
		fmt.Println()
	}
}

func compare(w io.Writer, input, pg, cr, pgUser, crUser string) error {
	for _, code := range []pgwirebase.FormatCode{pgwirebase.FormatText, pgwirebase.FormatBinary} {
		results := map[string][]byte{}
		for addr, user := range map[string]string{
			pg: pgUser,
			cr: crUser,
		} {
			res, err := exec(input, addr, user, code)
			if err != nil {
				return errors.Wrapf(err, "addr: %s, code: %s", addr, code)
			}
			for k, v := range results {
				if !bytes.Equal(res, v) {
					return errors.Errorf("%q (%[1]v, %s) != %q (%[3]v, %s)", v, k, res, addr)
				}
			}
			results[addr] = res
		}
	}
	return nil
}

func exec(input, addr, user string, code pgwirebase.FormatCode) ([]byte, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Printf("connecting to %s as %s with code %s\n", addr, user, code)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, errors.Wrap(err, "dail")
	}
	defer conn.Close()

	fe, err := pgproto3.NewFrontend(conn, conn)
	if err != nil {
		return nil, errors.Wrap(err, "new frontend")
	}

	send := make(chan pgproto3.FrontendMessage)
	recv := make(chan pgproto3.BackendMessage)
	logRecv := false

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(send)
				return
			case msg := <-send:
				//fmt.Printf("sending: %#v\n", msg)
				err := fe.Send(msg)
				if err != nil {
					panic(err)
				}
			}
		}
	}()
	go func() {
		for {
			msg, err := fe.Receive()
			if err != nil {
				panic(err)
			}
			if logRecv {
				//fmt.Printf("received: %#v\n", msg)
			}

			x := reflect.ValueOf(msg)
			starX := x.Elem()
			y := reflect.New(starX.Type())
			starY := y.Elem()
			starY.Set(starX)
			dup := y.Interface().(pgproto3.BackendMessage)

			select {
			case <-ctx.Done():
				close(recv)
				return
			case recv <- dup:
			}
		}
	}()

	send <- &pgproto3.StartupMessage{
		ProtocolVersion: 196608,
		Parameters: map[string]string{
			"user": user,
		},
	}
	{
		r := <-recv
		if msg, ok := r.(*pgproto3.Authentication); !ok || msg.Type != 0 {
			return nil, errors.Errorf("unexpected: %#v\n", r)
		}
	}
WaitConnLoop:
	for {
		msg := <-recv
		switch msg.(type) {
		case *pgproto3.ReadyForQuery:
			break WaitConnLoop
		}
	}
	send <- &pgproto3.Parse{
		Query: input,
	}
	send <- &pgproto3.Describe{
		ObjectType: 'S',
	}
	send <- &pgproto3.Sync{}
	r := <-recv
	if _, ok := r.(*pgproto3.ParseComplete); !ok {
		return nil, errors.Errorf("unexpected: %#v", r)
	}
	send <- &pgproto3.Bind{
		ResultFormatCodes: []int16{int16(code)},
	}
	send <- &pgproto3.Execute{}
	logRecv = true
	send <- &pgproto3.Sync{}
	var res []byte
WaitExecuteLoop:
	for {
		msg := <-recv
		switch msg := msg.(type) {
		case *pgproto3.DataRow:
			if res != nil {
				return nil, errors.New("already saw a row")
			}
			if len(msg.Values) != 1 {
				return nil, errors.Errorf("unexpected: %#v\n", msg)
			}
			res = msg.Values[0]
			//fmt.Printf("res: %q\n", res)
		case *pgproto3.CommandComplete,
			*pgproto3.EmptyQueryResponse,
			*pgproto3.ErrorResponse:
			break WaitExecuteLoop
		}
	}
	//fmt.Println()

	return res, nil
}
