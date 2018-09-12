package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/cockroachdb/cockroach/pkg/internal/rsg"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
)

func main() {
	if err := compare(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func compare() error {
	y, err := ioutil.ReadFile("gram.y")
	if err != nil {
		return err
	}
	r, err := rsg.NewRSG(0, "gram.y", string(y), false)
	if err != nil {
		return err
	}

	cr, err := pgx.Connect(pgx.ConnConfig{
		Host: "127.0.0.1",
		Port: 26257,
		User: "root",
	})
	if err != nil {
		return errors.Wrap(err, "cr connection")
	}

	pg, err := pgx.Connect(pgx.ConnConfig{
		Host: "127.0.0.1",
		Port: 5432,
		User: "postgres",
	})
	if err != nil {
		return errors.Wrap(err, "pg connection")
	}

	var wg sync.WaitGroup
	var tagCR, tagPG pgx.CommandTag
	var errCR, errPG error
	for i := 0; i < 10000; i++ {
		q := r.Generate("stmt", 200)

		wg.Add(2)
		go func() {
			tagCR, errCR = cr.Exec(q)
			wg.Done()
		}()
		go func() {
			tagPG, errPG = pg.Exec(q)
			wg.Done()
		}()
		wg.Wait()

		s := fmt.Sprintf("%s\n%v:\n\tcr: %s: %v\n\tpg: %s: %v\n\n",
			q,
			tagCR == tagPG,
			tagCR, errCR,
			tagPG, errPG,
		)
		if tagCR == "" && tagPG != "" {
			fmt.Print(s)
		}
	}
	return nil
}

// todo: filter out cr errors with a github issue in unimplemented. mark others that we need to add a github issue.
// buckets: things we actually want to do (i.e., alter role); not doing in near term, but maybe someday if enough care; probably never going to do something (like setting an isolation level we won't support).
