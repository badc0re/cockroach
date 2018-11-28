package parser

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
	
	"github.com/cockroachdb/cockroach/pkg/sql/lex"
)

func (sc *Scanner) scan(lval *sqlSymType) {

	%% machine scanner;
	%% write data;

	data := sc.in
	cs, p, pe, eof := sc.cs, sc.pos, len(data), len(data)
	_ = eof

	var (
		mark int
		s string
		uval uint64
		err error
		isFconst bool
		isUpper bool
		isNotASCII bool
		numQuote int
		b []byte
		ch byte
		rn rune
		buf *bytes.Buffer
	)
	str := func() { s = string(data[mark:p]) }
	emit := func(tok int, s string) {
		lval.id = tok
		lval.str = s
		lval.pos = p
		sc.pos = p
		sc.cs = cs
	}
	emitToken := func(tok int) {
		lval.id = tok
		lval.str = data[mark:p]
		lval.pos = p
		sc.pos = p
		sc.cs = cs
	}
	emitError := func(err string) {
		lval.id = ERROR
		lval.str = err
	}

	%%{
		action mark { mark = p }
		action str { str() }

		action placeholder {
			mark++
			str()
			uval, err = strconv.ParseUint(s, 10, 64)
			if err != nil {
				emitError(err.Error())
				return
			}
			if err == nil && uval > 1<<63 {
				emitError(fmt.Sprintf("integer value out of range: %d", uval))
				return
			}
			emit(lex.PLACEHOLDER, s); return
		}
		action number {
			// Get zeros.
			if mark == p {
				mark--
			}
			str()
			if isFconst {
				emit(lex.FCONST, s); return
			} else {
				emit(lex.ICONST, s); return
			}
			isFconst = false
		}
		action markZero {
			if mark == p && data[p] == '0' {
				mark++
			}
		}
		action fconst {
			isFconst = true
		}
		action hex {
			emit(lex.ICONST, data[mark-1:p]); return
		}
		action ident {
			if isNotASCII {
				str()
				s = lex.NormalizeName(s)
			} else if isUpper {
				b := make([]byte, p-mark)
				for i, c := range data[mark:p] {
					if c >= 'A' && c <= 'Z' {
						c += 'a' - 'A'
					}
					b[i] = byte(c)
				}
				s = string(b)
			} else {
				str()
			}

			if id, ok := lex.Keywords[s]; ok {
				emit(int(id.Tok), s); return
			} else {
				emit(lex.IDENT, s); return
			}
			isUpper = false
			isNotASCII = false
		}
		int = digit+;
		pn = ('-' | '+')?;
		number =
			pn
			(int $markZero ('.' >fconst)? digit* | '.' >fconst int)
			([eE] >fconst pn int)?
			;
		hex = '0x'i xdigit+;
		placeholder = '$' int;
		notASCII = 128..255 >{ isNotASCII = true };
		identStart =
			'a'..'z'
			| 'A'..'Z' >{ isUpper = true }
			| '_'
			| notASCII
			;
		ident =
			identStart
			(identStart | digit | '$')*
			;
		action identQuote {
			if numQuote != 0 {
				b = make([]byte, p-mark-1-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
						i++
					}
				}
				numQuote = 0
			} else {
				b = []byte(data[mark+1:p])
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					emitError(errInvalidUTF8)
					return
				}
				isNotASCII = false
			}
			s += string(b)
		}
		identQuote =
			'"' >mark >{fmt.Println("mark quote", string(data[p]))}
			(
				'""' %{ numQuote++ }
				| notASCII
				| /[^"]/
			)* %identQuote
			'"'
			;
		action emitIdent {
			emit(lex.IDENT, s); return
		}
		singleQuote =
			"'"
			(
				"''" %{ numQuote++ }
				| notASCII
				| /[^']/
			)*
			"'"
			;
		action singleQuote {
			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '\'' {
						i++
					}
				}
				s = string(b)
				numQuote = 0
			} else {
				b = []byte(data[mark+1:p-1])
			}
			if isNotASCII {
				if !utf8.Valid(b) {
					emitError(errInvalidUTF8)
					return
				}
				isNotASCII = false
			}
			emit(lex.SCONST, string(b)); return
		}
		escape =
			'a' %{ buf.WriteByte('\a') }
			| 'b' %{ buf.WriteByte('\b') }
			| 'f' %{ buf.WriteByte('\f') }
			| 'n' %{ buf.WriteByte('\n') }
			| 'r' %{ buf.WriteByte('\r') }
			| 't' %{ buf.WriteByte('\t') }
			| 'v' %{ buf.WriteByte('\v') }
			;
		slashHex =
			'x'i xdigit {2}
			>{ ch = 0 }
			${ ch = (ch << 4) | unhex(data[p]) }
			%{ buf.WriteByte(ch) }
			;
		slashUnicode =
			((
				'u' xdigit {4}
				${ rn = (rn << 4) | rune(unhex(data[p])) }
			) | (
				'U' xdigit {8}
				${ rn = (rn << 4) | rune(unhex(data[p])) }
			))
			>{ rn = 0 }
			%{ buf.WriteRune(rn) }
			;
		slashOctal =
			('0'..'7') {3}
			>{ ch = 0 }
			${ ch = (ch << 3) | data[p] - '0' }
			%{ buf.WriteByte(ch) }
			;
		stringLiteral = 
			"'" %{ buf = new(bytes.Buffer) }
			(
				(
					"''"
					| /[^'\\]/
				) @{ buf.WriteByte(data[p]) }
				| "\\" (
					escape
					| slashHex
					| slashUnicode
					| slashOctal
					| ^(escape | 'x'i | 'u'i | '0'..'7') ${ buf.WriteByte(data[p]) }
				)
			)*
			"'"
			;
		bytes = "b" stringLiteral;
		action bytes {
			emit(lex.BCONST, buf.String()); return
		}
		escapedString = "e"i stringLiteral;
		action escapedString {
			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		}
		hexString =
			'x'i %{ buf = new(bytes.Buffer) }
			"'"
			(
				xdigit {2}
				>{ ch = 0 }
				${ ch = (ch << 4) | unhex(data[p]) }
				%{ buf.WriteByte(ch) }
			)*
			"'"
			;
		bitArray =
			"B'"
			(/[01]/)*
			"'"
			;
		action bitArray {
			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		}
		top =
			  space
			| '--' [^\n]*
			| '/*' any* '*/' # TODO: support nesting
			| hex >mark %hex @err{
				emitError(errInvalidHexNumeric)
				return
			}
			| number >mark %number @err{
				emitError("invalid floating point literal")
				return
			}
			| placeholder >mark %placeholder
			| ident >mark %ident
			| identQuote >{s = ""}
				(space* '\n' space* identQuote)*
				%emitIdent
			| singleQuote >mark %singleQuote
			| bytes %bytes
			| escapedString %escapedString
			| hexString %bytes @err{
				emitError("invalid hexadecimal bytes literal")
				return
			}
			| bitArray >mark %bitArray @err{
				emitError(fmt.Sprintf(`"%c" is not a valid binary digit`, data[p]))
				return
			}

			| punct >mark %{ emitToken(int(data[p-1])); return }

			| '..' >mark %{ emitToken(lex.DOT_DOT); return }

			| '!=' >mark %{ emitToken(lex.NOT_EQUALS); return }
			| '!~*' >mark %{ emitToken(lex.NOT_REGIMATCH); return }
			| '!~' >mark %{ emitToken(lex.NOT_REGMATCH); return }

			| '??' >mark %{ emitToken(lex.HELPTOKEN); return }
			| '?|' >mark %{ emitToken(lex.JSON_SOME_EXISTS); return }
			| '?&' >mark %{ emitToken(lex.JSON_ALL_EXISTS); return }

			| '<<=' >mark %{ emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return }
			| '<<' >mark %{ emitToken(lex.LSHIFT); return }
			| '<>' >mark %{ emitToken(lex.NOT_EQUALS); return }
			| '<=' >mark %{ emitToken(lex.LESS_EQUALS); return }
			| '<@' >mark %{ emitToken(lex.CONTAINED_BY); return }

			| '>>=' >mark %{ emitToken(lex.INET_CONTAINS_OR_EQUALS); return }
			| '>>' >mark %{ emitToken(lex.RSHIFT); return }
			| '>=' >mark %{ emitToken(lex.GREATER_EQUALS); return }

			| ':::' >mark %{ emitToken(lex.TYPEANNOTATE); return }
			| '::' >mark %{ emitToken(lex.TYPECAST); return }

			| '||' >mark %{ emitToken(lex.CONCAT); return }

			| '//' >mark %{ emitToken(lex.FLOORDIV); return }

			| '~*' >mark %{ emitToken(lex.REGIMATCH); return }

			| '@>' >mark %{ emitToken(lex.CONTAINS); return }

			| '&&' >mark %{ emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return }

			| '->>' >mark %{ emitToken(lex.FETCHTEXT); return }
			| '->' >mark %{ emitToken(lex.FETCHVAL); return }

			| '#>>' >mark %{ emitToken(lex.FETCHTEXT_PATH); return }
			| '#>' >mark %{ emitToken(lex.FETCHVAL_PATH); return }
			| '#-' >mark %{ emitToken(lex.REMOVE_PATH); return }

			;
		main :=
			top** @err{
				emitError("invalid syntax")
				return
			}
			%{emit(0, "EOF"); return}
			;

		write init;
		write exec;
	}%%
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
