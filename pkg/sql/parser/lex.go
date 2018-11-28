
//line lex.rl:1
package parser

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
	
	"github.com/cockroachdb/cockroach/pkg/sql/lex"
)

func (sc *Scanner) scan(lval *sqlSymType) {

	
//line lex.rl:15
	
//line lex.go:20
const scanner_start int = 71
const scanner_first_final int = 71
const scanner_error int = 0

const scanner_en_main int = 71


//line lex.rl:16

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

	
//line lex.go:69
	{
	cs = scanner_start
	}

//line lex.go:74
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 71:
		goto st_case_71
	case 0:
		goto st_case_0
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 1:
		goto st_case_1
	case 74:
		goto st_case_74
	case 2:
		goto st_case_2
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 3:
		goto st_case_3
	case 81:
		goto st_case_81
	case 4:
		goto st_case_4
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 5:
		goto st_case_5
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 6:
		goto st_case_6
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 9:
		goto st_case_9
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 116:
		goto st_case_116
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 125:
		goto st_case_125
	case 70:
		goto st_case_70
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	}
	goto st_out
tr188:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
	goto st71
tr220:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
	goto st71
tr256:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
	goto st71
tr287:
//line lex.rl:186

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
		
	goto st71
tr322:
//line lex.rl:73

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
		
	goto st71
tr351:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
	goto st71
tr383:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
	goto st71
tr417:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
	goto st71
tr448:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
	goto st71
tr483:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
	goto st71
tr515:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
	goto st71
tr548:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
	goto st71
tr582:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
	goto st71
tr614:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
	goto st71
tr645:
//line lex.rl:97

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
		
	goto st71
tr667:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
	goto st71
tr699:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
	goto st71
tr730:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
	goto st71
tr761:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
	goto st71
tr780:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
	goto st71
tr812:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
	goto st71
tr843:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
	goto st71
tr874:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
	goto st71
tr905:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
	goto st71
tr937:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
	goto st71
tr968:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
	goto st71
tr999:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
	goto st71
tr1030:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
	goto st71
tr1063:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
	goto st71
tr1088:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
	goto st71
tr1119:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
	goto st71
tr1150:
//line lex.rl:59

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
		
	goto st71
tr1179:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
	goto st71
tr1210:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
	goto st71
tr1241:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
	goto st71
tr1272:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
	goto st71
tr1303:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
	goto st71
tr1335:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
	goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line lex.go:648
		switch data[p] {
		case 32:
			goto st71
		case 33:
			goto tr158
		case 34:
			goto tr159
		case 35:
			goto tr160
		case 36:
			goto tr161
		case 38:
			goto tr163
		case 39:
			goto tr164
		case 43:
			goto tr165
		case 45:
			goto tr166
		case 46:
			goto tr167
		case 47:
			goto tr168
		case 48:
			goto tr169
		case 58:
			goto tr171
		case 60:
			goto tr172
		case 62:
			goto tr173
		case 63:
			goto tr174
		case 64:
			goto tr175
		case 66:
			goto tr177
		case 69:
			goto tr178
		case 88:
			goto tr179
		case 95:
			goto tr180
		case 98:
			goto tr182
		case 101:
			goto tr183
		case 120:
			goto tr184
		case 124:
			goto tr185
		case 126:
			goto tr186
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr162
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto st71
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr176
					}
				case data[p] >= 59:
					goto tr162
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr162
					}
				case data[p] >= 97:
					goto tr181
				}
			default:
				goto tr162
			}
		default:
			goto tr170
		}
		goto tr187
tr44:
//line lex.rl:366

				emitError("invalid syntax")
				return
			
	goto st0
tr12:
//line lex.rl:303

				emitError("invalid floating point literal")
				return
			
//line lex.rl:366

				emitError("invalid syntax")
				return
			
	goto st0
tr14:
//line lex.rl:317

				emitError(fmt.Sprintf(`"%c" is not a valid binary digit`, data[p]))
				return
			
//line lex.rl:366

				emitError("invalid syntax")
				return
			
	goto st0
tr23:
//line lex.rl:313

				emitError("invalid hexadecimal bytes literal")
				return
			
//line lex.rl:366

				emitError("invalid syntax")
				return
			
	goto st0
tr154:
//line lex.rl:299

				emitError(errInvalidHexNumeric)
				return
			
//line lex.rl:366

				emitError("invalid syntax")
				return
			
	goto st0
tr156:
//line lex.rl:299

				emitError(errInvalidHexNumeric)
				return
			
//line lex.rl:303

				emitError("invalid floating point literal")
				return
			
//line lex.rl:313

				emitError("invalid hexadecimal bytes literal")
				return
			
//line lex.rl:317

				emitError(fmt.Sprintf(`"%c" is not a valid binary digit`, data[p]))
				return
			
	goto st0
//line lex.go:828
st_case_0:
	st0:
		cs = 0
		goto _out
tr158:
//line lex.rl:56
 mark = p 
	goto st72
tr189:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st72
tr221:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st72
tr257:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st72
tr288:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st72
tr323:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st72
tr352:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st72
tr384:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st72
tr418:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st72
tr449:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st72
tr484:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st72
tr516:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st72
tr549:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st72
tr583:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st72
tr615:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st72
tr646:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st72
tr668:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st72
tr700:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st72
tr731:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st72
tr762:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st72
tr781:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st72
tr813:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st72
tr844:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st72
tr875:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st72
tr906:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st72
tr938:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st72
tr969:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st72
tr1000:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st72
tr1031:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st72
tr1064:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st72
tr1089:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st72
tr1120:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st72
tr1151:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st72
tr1180:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st72
tr1211:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st72
tr1242:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st72
tr1273:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st72
tr1304:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st72
tr1336:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line lex.go:1207
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto tr194
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto tr197
		case 46:
			goto tr198
		case 47:
			goto tr199
		case 48:
			goto tr200
		case 58:
			goto tr202
		case 59:
			goto tr193
		case 60:
			goto tr203
		case 61:
			goto st133
		case 62:
			goto tr205
		case 63:
			goto tr206
		case 64:
			goto tr207
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto tr217
		case 126:
			goto st134
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr208
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto tr201
		}
		goto tr219
tr159:
//line lex.rl:56
 mark = p 
	goto st73
tr190:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st73
tr258:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st73
tr289:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st73
tr324:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st73
tr353:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st73
tr385:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st73
tr419:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st73
tr450:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st73
tr485:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st73
tr517:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st73
tr550:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st73
tr584:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st73
tr616:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st73
tr647:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st73
tr669:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st73
tr701:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st73
tr732:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st73
tr763:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st73
tr782:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st73
tr814:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st73
tr845:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st73
tr876:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st73
tr907:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st73
tr939:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st73
tr970:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st73
tr1001:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st73
tr1032:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st73
tr1065:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st73
tr1090:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st73
tr1121:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st73
tr1152:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st73
tr1181:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st73
tr1212:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st73
tr1243:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st73
tr1274:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st73
tr1305:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st73
tr1337:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line lex.go:1652
		if data[p] == 34 {
			goto st74
		}
		if data[p] <= 127 {
			goto st1
		}
		goto tr2
tr2:
//line lex.rl:131
 isNotASCII = true 
	goto st1
tr3:
//line lex.rl:145
 numQuote++ 
	goto st1
tr5:
//line lex.rl:145
 numQuote++ 
//line lex.rl:131
 isNotASCII = true 
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line lex.go:1679
		if data[p] == 34 {
			goto st74
		}
		if data[p] <= 127 {
			goto st1
		}
		goto tr2
tr4:
//line lex.rl:145
 numQuote++ 
	goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line lex.go:1696
		switch data[p] {
		case 32:
			goto tr220
		case 33:
			goto tr221
		case 34:
			goto st2
		case 35:
			goto tr223
		case 36:
			goto tr224
		case 38:
			goto tr226
		case 39:
			goto tr227
		case 43:
			goto tr228
		case 45:
			goto tr229
		case 46:
			goto tr230
		case 47:
			goto tr231
		case 48:
			goto tr232
		case 58:
			goto tr234
		case 60:
			goto tr235
		case 62:
			goto tr236
		case 63:
			goto tr237
		case 64:
			goto tr238
		case 66:
			goto tr240
		case 69:
			goto tr241
		case 88:
			goto tr242
		case 95:
			goto tr243
		case 98:
			goto tr245
		case 101:
			goto tr246
		case 120:
			goto tr247
		case 124:
			goto tr248
		case 126:
			goto tr249
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr225
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr220
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr239
					}
				case data[p] >= 59:
					goto tr225
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr225
					}
				case data[p] >= 97:
					goto tr244
				}
			default:
				goto tr225
			}
		default:
			goto tr233
		}
		goto tr250
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 34 {
			goto tr4
		}
		if data[p] <= 127 {
			goto tr3
		}
		goto tr5
tr160:
//line lex.rl:56
 mark = p 
	goto st75
tr191:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st75
tr223:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st75
tr259:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st75
tr290:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st75
tr325:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st75
tr354:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st75
tr386:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st75
tr420:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st75
tr451:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st75
tr486:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st75
tr518:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st75
tr551:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st75
tr585:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st75
tr617:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st75
tr648:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st75
tr670:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st75
tr702:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st75
tr733:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st75
tr764:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st75
tr783:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st75
tr815:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st75
tr846:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st75
tr877:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st75
tr908:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st75
tr940:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st75
tr971:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st75
tr1002:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st75
tr1033:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st75
tr1066:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st75
tr1091:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st75
tr1122:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st75
tr1153:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st75
tr1182:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st75
tr1213:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st75
tr1244:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st75
tr1275:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st75
tr1306:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st75
tr1338:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line lex.go:2185
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto tr194
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto st130
		case 46:
			goto tr198
		case 47:
			goto tr199
		case 48:
			goto tr200
		case 58:
			goto tr202
		case 60:
			goto tr203
		case 62:
			goto st131
		case 63:
			goto tr206
		case 64:
			goto tr207
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto tr217
		case 126:
			goto tr253
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr208
					}
				case data[p] >= 59:
					goto tr193
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto tr201
		}
		goto tr219
tr161:
//line lex.rl:56
 mark = p 
	goto st76
tr192:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st76
tr224:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st76
tr260:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st76
tr291:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st76
tr326:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st76
tr355:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st76
tr387:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st76
tr421:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st76
tr452:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st76
tr487:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st76
tr519:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st76
tr552:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st76
tr586:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st76
tr618:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st76
tr671:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st76
tr703:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st76
tr734:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st76
tr784:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st76
tr816:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st76
tr847:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st76
tr878:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st76
tr909:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st76
tr941:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st76
tr972:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st76
tr1003:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st76
tr1034:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st76
tr1067:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st76
tr1092:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st76
tr1123:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st76
tr1154:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st76
tr1183:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st76
tr1214:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st76
tr1245:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st76
tr1276:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st76
tr1307:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st76
tr1339:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line lex.go:2600
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto tr194
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto tr197
		case 46:
			goto tr198
		case 47:
			goto tr199
		case 58:
			goto tr202
		case 60:
			goto tr203
		case 62:
			goto tr205
		case 63:
			goto tr206
		case 64:
			goto tr207
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto tr217
		case 126:
			goto tr253
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr208
					}
				case data[p] >= 59:
					goto tr193
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto st129
		}
		goto tr219
tr162:
//line lex.rl:56
 mark = p 
	goto st77
tr193:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st77
tr225:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st77
tr261:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st77
tr292:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st77
tr327:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st77
tr356:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st77
tr388:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st77
tr422:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st77
tr453:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st77
tr488:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st77
tr520:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st77
tr553:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st77
tr587:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st77
tr619:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st77
tr650:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st77
tr672:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st77
tr704:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st77
tr735:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st77
tr765:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st77
tr785:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st77
tr817:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st77
tr848:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st77
tr879:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st77
tr910:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st77
tr942:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st77
tr973:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st77
tr1004:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st77
tr1035:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st77
tr1068:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st77
tr1093:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st77
tr1124:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st77
tr1155:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st77
tr1184:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st77
tr1215:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st77
tr1246:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st77
tr1277:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st77
tr1308:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st77
tr1340:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line lex.go:3075
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto tr194
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto tr197
		case 46:
			goto tr198
		case 47:
			goto tr199
		case 48:
			goto tr200
		case 58:
			goto tr202
		case 60:
			goto tr203
		case 62:
			goto tr205
		case 63:
			goto tr206
		case 64:
			goto tr207
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto tr217
		case 126:
			goto tr253
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr208
					}
				case data[p] >= 59:
					goto tr193
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto tr201
		}
		goto tr219
tr163:
//line lex.rl:56
 mark = p 
	goto st78
tr194:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st78
tr226:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st78
tr262:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st78
tr293:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st78
tr328:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st78
tr357:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st78
tr389:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st78
tr423:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st78
tr454:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st78
tr489:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st78
tr521:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st78
tr554:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st78
tr588:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st78
tr620:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st78
tr651:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st78
tr673:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st78
tr705:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st78
tr736:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st78
tr766:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st78
tr786:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st78
tr818:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st78
tr849:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st78
tr880:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st78
tr911:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st78
tr943:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st78
tr974:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st78
tr1005:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st78
tr1036:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st78
tr1069:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st78
tr1094:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st78
tr1125:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st78
tr1156:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st78
tr1185:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st78
tr1216:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st78
tr1247:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st78
tr1278:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st78
tr1309:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st78
tr1341:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
//line lex.go:3552
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto st79
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto tr197
		case 46:
			goto tr198
		case 47:
			goto tr199
		case 48:
			goto tr200
		case 58:
			goto tr202
		case 60:
			goto tr203
		case 62:
			goto tr205
		case 63:
			goto tr206
		case 64:
			goto tr207
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto tr217
		case 126:
			goto tr253
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr208
					}
				case data[p] >= 59:
					goto tr193
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto tr201
		}
		goto tr219
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 32:
			goto tr256
		case 33:
			goto tr257
		case 34:
			goto tr258
		case 35:
			goto tr259
		case 36:
			goto tr260
		case 38:
			goto tr262
		case 39:
			goto tr263
		case 43:
			goto tr264
		case 45:
			goto tr265
		case 46:
			goto tr266
		case 47:
			goto tr267
		case 48:
			goto tr268
		case 58:
			goto tr270
		case 60:
			goto tr271
		case 62:
			goto tr272
		case 63:
			goto tr273
		case 64:
			goto tr274
		case 66:
			goto tr276
		case 69:
			goto tr277
		case 88:
			goto tr278
		case 95:
			goto tr279
		case 98:
			goto tr281
		case 101:
			goto tr282
		case 120:
			goto tr283
		case 124:
			goto tr284
		case 126:
			goto tr285
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr261
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr256
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr275
					}
				case data[p] >= 59:
					goto tr261
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr261
					}
				case data[p] >= 97:
					goto tr280
				}
			default:
				goto tr261
			}
		default:
			goto tr269
		}
		goto tr286
tr164:
//line lex.rl:56
 mark = p 
	goto st80
tr195:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st80
tr227:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st80
tr263:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st80
tr329:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st80
tr358:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st80
tr390:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st80
tr424:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st80
tr455:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st80
tr490:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st80
tr522:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st80
tr555:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st80
tr589:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st80
tr621:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st80
tr652:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st80
tr674:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st80
tr737:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st80
tr767:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st80
tr787:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st80
tr819:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st80
tr850:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st80
tr881:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st80
tr912:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st80
tr944:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st80
tr975:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st80
tr1006:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st80
tr1037:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st80
tr1070:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st80
tr1095:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st80
tr1126:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st80
tr1157:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st80
tr1186:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st80
tr1217:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st80
tr1248:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st80
tr1279:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st80
tr1310:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st80
tr1342:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line lex.go:4093
		if data[p] == 39 {
			goto st81
		}
		if data[p] <= 127 {
			goto st3
		}
		goto tr8
tr8:
//line lex.rl:131
 isNotASCII = true 
	goto st3
tr9:
//line lex.rl:180
 numQuote++ 
	goto st3
tr11:
//line lex.rl:180
 numQuote++ 
//line lex.rl:131
 isNotASCII = true 
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line lex.go:4120
		if data[p] == 39 {
			goto st81
		}
		if data[p] <= 127 {
			goto st3
		}
		goto tr8
tr10:
//line lex.rl:180
 numQuote++ 
	goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
//line lex.go:4137
		switch data[p] {
		case 32:
			goto tr287
		case 33:
			goto tr288
		case 34:
			goto tr289
		case 35:
			goto tr290
		case 36:
			goto tr291
		case 38:
			goto tr293
		case 39:
			goto st4
		case 43:
			goto tr295
		case 45:
			goto tr296
		case 46:
			goto tr297
		case 47:
			goto tr298
		case 48:
			goto tr299
		case 58:
			goto tr301
		case 60:
			goto tr302
		case 62:
			goto tr303
		case 63:
			goto tr304
		case 64:
			goto tr305
		case 66:
			goto tr307
		case 69:
			goto tr308
		case 88:
			goto tr309
		case 95:
			goto tr310
		case 98:
			goto tr312
		case 101:
			goto tr313
		case 120:
			goto tr314
		case 124:
			goto tr315
		case 126:
			goto tr316
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr292
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr287
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr306
					}
				case data[p] >= 59:
					goto tr292
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr292
					}
				case data[p] >= 97:
					goto tr311
				}
			default:
				goto tr292
			}
		default:
			goto tr300
		}
		goto tr317
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 39 {
			goto tr10
		}
		if data[p] <= 127 {
			goto tr9
		}
		goto tr11
tr165:
//line lex.rl:56
 mark = p 
	goto st82
tr196:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st82
tr228:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st82
tr264:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st82
tr295:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st82
tr330:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st82
tr359:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st82
tr391:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st82
tr425:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st82
tr456:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st82
tr491:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st82
tr523:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st82
tr556:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st82
tr590:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st82
tr622:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st82
tr653:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st82
tr675:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st82
tr706:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st82
tr738:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st82
tr768:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st82
tr788:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st82
tr820:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st82
tr851:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st82
tr882:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st82
tr913:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st82
tr945:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st82
tr976:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st82
tr1007:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st82
tr1038:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st82
tr1071:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st82
tr1096:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st82
tr1127:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st82
tr1158:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st82
tr1187:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st82
tr1218:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st82
tr1249:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st82
tr1280:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st82
tr1312:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st82
tr1343:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line lex.go:4626
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto tr194
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto tr197
		case 46:
			goto tr318
		case 47:
			goto tr199
		case 58:
			goto tr202
		case 60:
			goto tr203
		case 62:
			goto tr205
		case 63:
			goto tr206
		case 64:
			goto tr207
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto tr217
		case 126:
			goto tr253
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr208
					}
				case data[p] >= 59:
					goto tr193
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto tr319
		}
		goto tr219
tr166:
//line lex.rl:56
 mark = p 
	goto st83
tr197:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st83
tr229:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st83
tr265:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st83
tr296:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st83
tr331:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st83
tr360:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st83
tr392:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st83
tr426:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st83
tr457:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st83
tr492:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st83
tr524:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st83
tr557:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st83
tr591:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st83
tr623:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st83
tr654:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st83
tr676:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st83
tr707:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st83
tr739:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st83
tr769:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st83
tr789:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st83
tr821:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st83
tr852:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st83
tr883:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st83
tr914:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st83
tr946:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st83
tr977:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st83
tr1008:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st83
tr1039:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st83
tr1072:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st83
tr1097:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st83
tr1128:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st83
tr1159:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st83
tr1188:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st83
tr1219:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st83
tr1250:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st83
tr1281:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st83
tr1313:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st83
tr1344:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st83
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
//line lex.go:5101
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto tr194
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto st84
		case 46:
			goto tr318
		case 47:
			goto tr199
		case 58:
			goto tr202
		case 60:
			goto tr203
		case 62:
			goto st127
		case 63:
			goto tr206
		case 64:
			goto tr207
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto tr217
		case 126:
			goto tr253
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr208
					}
				case data[p] >= 59:
					goto tr193
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto tr319
		}
		goto tr219
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		if data[p] == 10 {
			goto st71
		}
		goto st84
tr318:
//line lex.rl:91

			isFconst = true
		
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line lex.go:5222
		if 48 <= data[p] && data[p] <= 57 {
			goto st85
		}
		goto tr12
tr414:
//line lex.rl:91

			isFconst = true
		
	goto st85
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
//line lex.go:5238
		switch data[p] {
		case 32:
			goto tr322
		case 33:
			goto tr323
		case 34:
			goto tr324
		case 35:
			goto tr325
		case 36:
			goto tr326
		case 38:
			goto tr328
		case 39:
			goto tr329
		case 43:
			goto tr330
		case 45:
			goto tr331
		case 46:
			goto tr332
		case 47:
			goto tr333
		case 58:
			goto tr334
		case 60:
			goto tr335
		case 62:
			goto tr336
		case 63:
			goto tr337
		case 64:
			goto tr338
		case 66:
			goto tr340
		case 69:
			goto tr341
		case 88:
			goto tr342
		case 95:
			goto tr343
		case 98:
			goto tr345
		case 101:
			goto tr341
		case 120:
			goto tr346
		case 124:
			goto tr347
		case 126:
			goto tr348
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr327
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr322
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr339
					}
				case data[p] >= 59:
					goto tr327
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr327
					}
				case data[p] >= 97:
					goto tr344
				}
			default:
				goto tr327
			}
		default:
			goto st85
		}
		goto tr349
tr167:
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr198:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr230:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr266:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr297:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr332:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr361:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr393:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr427:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr458:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr493:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr525:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr558:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr592:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr624:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr655:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr677:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr708:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr740:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr770:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr790:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr822:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr853:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr884:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr915:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr947:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr978:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr1009:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr1040:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr1073:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr1098:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr1129:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr1160:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr1189:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr1220:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr1251:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr1282:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr1314:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
tr1345:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:91

			isFconst = true
		
	goto st86
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
//line lex.go:5869
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto tr194
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto tr197
		case 46:
			goto st87
		case 47:
			goto tr199
		case 58:
			goto tr202
		case 60:
			goto tr203
		case 62:
			goto tr205
		case 63:
			goto tr206
		case 64:
			goto tr207
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto tr217
		case 126:
			goto tr253
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr208
					}
				case data[p] >= 59:
					goto tr193
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto st85
		}
		goto tr219
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch data[p] {
		case 32:
			goto tr351
		case 33:
			goto tr352
		case 34:
			goto tr353
		case 35:
			goto tr354
		case 36:
			goto tr355
		case 38:
			goto tr357
		case 39:
			goto tr358
		case 43:
			goto tr359
		case 45:
			goto tr360
		case 46:
			goto tr361
		case 47:
			goto tr362
		case 48:
			goto tr363
		case 58:
			goto tr365
		case 60:
			goto tr366
		case 62:
			goto tr367
		case 63:
			goto tr368
		case 64:
			goto tr369
		case 66:
			goto tr371
		case 69:
			goto tr372
		case 88:
			goto tr373
		case 95:
			goto tr374
		case 98:
			goto tr376
		case 101:
			goto tr377
		case 120:
			goto tr378
		case 124:
			goto tr379
		case 126:
			goto tr380
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr356
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr351
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr370
					}
				case data[p] >= 59:
					goto tr356
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr356
					}
				case data[p] >= 97:
					goto tr375
				}
			default:
				goto tr356
			}
		default:
			goto tr364
		}
		goto tr381
tr168:
//line lex.rl:56
 mark = p 
	goto st88
tr199:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st88
tr231:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st88
tr267:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st88
tr298:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st88
tr333:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st88
tr362:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st88
tr394:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st88
tr428:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st88
tr459:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st88
tr494:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st88
tr526:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st88
tr559:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st88
tr593:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st88
tr625:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st88
tr656:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st88
tr678:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st88
tr709:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st88
tr741:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st88
tr771:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st88
tr791:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st88
tr823:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st88
tr854:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st88
tr885:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st88
tr916:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st88
tr948:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st88
tr979:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st88
tr1010:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st88
tr1041:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st88
tr1074:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st88
tr1099:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st88
tr1130:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st88
tr1161:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st88
tr1190:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st88
tr1221:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st88
tr1252:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st88
tr1283:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st88
tr1315:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st88
tr1346:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st88
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
//line lex.go:6451
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto tr194
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto tr197
		case 46:
			goto tr198
		case 47:
			goto st89
		case 48:
			goto tr200
		case 58:
			goto tr202
		case 60:
			goto tr203
		case 62:
			goto tr205
		case 63:
			goto tr206
		case 64:
			goto tr207
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto tr217
		case 126:
			goto tr253
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr208
					}
				case data[p] >= 59:
					goto tr193
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto tr201
		}
		goto tr219
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 32:
			goto tr383
		case 33:
			goto tr384
		case 34:
			goto tr385
		case 35:
			goto tr386
		case 36:
			goto tr387
		case 38:
			goto tr389
		case 39:
			goto tr390
		case 43:
			goto tr391
		case 45:
			goto tr392
		case 46:
			goto tr393
		case 47:
			goto tr394
		case 48:
			goto tr395
		case 58:
			goto tr397
		case 60:
			goto tr398
		case 62:
			goto tr399
		case 63:
			goto tr400
		case 64:
			goto tr401
		case 66:
			goto tr403
		case 69:
			goto tr404
		case 88:
			goto tr405
		case 95:
			goto tr406
		case 98:
			goto tr408
		case 101:
			goto tr409
		case 120:
			goto tr410
		case 124:
			goto tr411
		case 126:
			goto tr412
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr388
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr383
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr402
					}
				case data[p] >= 59:
					goto tr388
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr388
					}
				case data[p] >= 97:
					goto tr407
				}
			default:
				goto tr388
			}
		default:
			goto tr396
		}
		goto tr413
tr169:
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr200:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr232:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr268:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr299:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr363:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr395:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr429:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr460:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr495:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr527:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr560:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr594:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr626:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr679:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr710:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr742:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr792:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr824:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr855:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr886:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr917:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr949:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr980:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1011:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1042:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1100:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1131:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1191:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1222:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1253:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1284:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1316:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
tr1347:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
//line lex.go:7132
		switch data[p] {
		case 32:
			goto tr322
		case 33:
			goto tr323
		case 34:
			goto tr324
		case 35:
			goto tr325
		case 36:
			goto tr326
		case 38:
			goto tr328
		case 39:
			goto tr329
		case 43:
			goto tr330
		case 45:
			goto tr331
		case 46:
			goto tr414
		case 47:
			goto tr333
		case 58:
			goto tr334
		case 60:
			goto tr335
		case 62:
			goto tr336
		case 63:
			goto tr337
		case 64:
			goto tr338
		case 66:
			goto tr340
		case 69:
			goto tr341
		case 88:
			goto st70
		case 95:
			goto tr343
		case 98:
			goto tr345
		case 101:
			goto tr341
		case 120:
			goto st70
		case 124:
			goto tr347
		case 126:
			goto tr348
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr327
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr322
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr339
					}
				case data[p] >= 59:
					goto tr327
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr327
					}
				case data[p] >= 97:
					goto tr344
				}
			default:
				goto tr327
			}
		default:
			goto tr319
		}
		goto tr349
tr170:
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr201:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr233:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr269:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr300:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr319:
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr364:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr396:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr430:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr461:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr496:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr528:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr561:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr595:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr627:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr680:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr711:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr743:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr793:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr825:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr856:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr887:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr918:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr950:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr981:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1012:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1043:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1101:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1132:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1192:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1223:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1254:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1285:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1317:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
tr1348:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:86

			if mark == p && data[p] == '0' {
				mark++
			}
		
	goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
//line lex.go:7712
		switch data[p] {
		case 32:
			goto tr322
		case 33:
			goto tr323
		case 34:
			goto tr324
		case 35:
			goto tr325
		case 36:
			goto tr326
		case 38:
			goto tr328
		case 39:
			goto tr329
		case 43:
			goto tr330
		case 45:
			goto tr331
		case 46:
			goto tr414
		case 47:
			goto tr333
		case 58:
			goto tr334
		case 60:
			goto tr335
		case 62:
			goto tr336
		case 63:
			goto tr337
		case 64:
			goto tr338
		case 66:
			goto tr340
		case 69:
			goto tr341
		case 88:
			goto tr342
		case 95:
			goto tr343
		case 98:
			goto tr345
		case 101:
			goto tr341
		case 120:
			goto tr346
		case 124:
			goto tr347
		case 126:
			goto tr348
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr327
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr322
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr339
					}
				case data[p] >= 59:
					goto tr327
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr327
					}
				case data[p] >= 97:
					goto tr344
				}
			default:
				goto tr327
			}
		default:
			goto tr319
		}
		goto tr349
tr171:
//line lex.rl:56
 mark = p 
	goto st92
tr202:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st92
tr234:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st92
tr270:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st92
tr301:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st92
tr334:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st92
tr365:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st92
tr397:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st92
tr462:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st92
tr497:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st92
tr529:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st92
tr562:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st92
tr596:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st92
tr628:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st92
tr657:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st92
tr681:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st92
tr712:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st92
tr744:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st92
tr772:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st92
tr794:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st92
tr826:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st92
tr857:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st92
tr888:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st92
tr919:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st92
tr951:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st92
tr982:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st92
tr1013:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st92
tr1044:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st92
tr1075:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st92
tr1102:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st92
tr1133:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st92
tr1162:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st92
tr1193:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st92
tr1224:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st92
tr1255:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st92
tr1286:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st92
tr1318:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st92
tr1349:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
//line lex.go:8181
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto tr194
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto tr197
		case 46:
			goto tr198
		case 47:
			goto tr199
		case 48:
			goto tr200
		case 58:
			goto st93
		case 60:
			goto tr203
		case 62:
			goto tr205
		case 63:
			goto tr206
		case 64:
			goto tr207
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto tr217
		case 126:
			goto tr253
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr208
					}
				case data[p] >= 59:
					goto tr193
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto tr201
		}
		goto tr219
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 32:
			goto tr417
		case 33:
			goto tr418
		case 34:
			goto tr419
		case 35:
			goto tr420
		case 36:
			goto tr421
		case 38:
			goto tr423
		case 39:
			goto tr424
		case 43:
			goto tr425
		case 45:
			goto tr426
		case 46:
			goto tr427
		case 47:
			goto tr428
		case 48:
			goto tr429
		case 58:
			goto st94
		case 60:
			goto tr432
		case 62:
			goto tr433
		case 63:
			goto tr434
		case 64:
			goto tr435
		case 66:
			goto tr437
		case 69:
			goto tr438
		case 88:
			goto tr439
		case 95:
			goto tr440
		case 98:
			goto tr442
		case 101:
			goto tr443
		case 120:
			goto tr444
		case 124:
			goto tr445
		case 126:
			goto tr446
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr422
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr417
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr436
					}
				case data[p] >= 59:
					goto tr422
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr422
					}
				case data[p] >= 97:
					goto tr441
				}
			default:
				goto tr422
			}
		default:
			goto tr430
		}
		goto tr447
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 32:
			goto tr448
		case 33:
			goto tr449
		case 34:
			goto tr450
		case 35:
			goto tr451
		case 36:
			goto tr452
		case 38:
			goto tr454
		case 39:
			goto tr455
		case 43:
			goto tr456
		case 45:
			goto tr457
		case 46:
			goto tr458
		case 47:
			goto tr459
		case 48:
			goto tr460
		case 58:
			goto tr462
		case 60:
			goto tr463
		case 62:
			goto tr464
		case 63:
			goto tr465
		case 64:
			goto tr466
		case 66:
			goto tr468
		case 69:
			goto tr469
		case 88:
			goto tr470
		case 95:
			goto tr471
		case 98:
			goto tr473
		case 101:
			goto tr474
		case 120:
			goto tr475
		case 124:
			goto tr476
		case 126:
			goto tr477
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr453
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr448
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr467
					}
				case data[p] >= 59:
					goto tr453
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr453
					}
				case data[p] >= 97:
					goto tr472
				}
			default:
				goto tr453
			}
		default:
			goto tr461
		}
		goto tr478
tr172:
//line lex.rl:56
 mark = p 
	goto st95
tr203:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st95
tr235:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st95
tr271:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st95
tr302:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st95
tr335:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st95
tr366:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st95
tr398:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st95
tr432:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st95
tr463:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st95
tr498:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st95
tr530:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st95
tr563:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st95
tr597:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st95
tr629:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st95
tr658:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st95
tr682:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st95
tr713:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st95
tr745:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st95
tr773:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st95
tr795:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st95
tr827:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st95
tr858:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st95
tr889:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st95
tr920:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st95
tr952:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st95
tr983:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st95
tr1014:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st95
tr1045:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st95
tr1076:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st95
tr1103:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st95
tr1134:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st95
tr1163:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st95
tr1194:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st95
tr1225:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st95
tr1256:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st95
tr1287:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st95
tr1319:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st95
tr1350:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
//line lex.go:8872
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto tr194
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto tr197
		case 46:
			goto tr198
		case 47:
			goto tr199
		case 48:
			goto tr200
		case 58:
			goto tr202
		case 59:
			goto tr193
		case 60:
			goto st96
		case 61:
			goto st122
		case 62:
			goto st123
		case 63:
			goto tr206
		case 64:
			goto st124
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto tr217
		case 126:
			goto tr253
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr208
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto tr201
		}
		goto tr219
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 32:
			goto tr483
		case 33:
			goto tr484
		case 34:
			goto tr485
		case 35:
			goto tr486
		case 36:
			goto tr487
		case 38:
			goto tr489
		case 39:
			goto tr490
		case 43:
			goto tr491
		case 45:
			goto tr492
		case 46:
			goto tr493
		case 47:
			goto tr494
		case 48:
			goto tr495
		case 58:
			goto tr497
		case 59:
			goto tr488
		case 60:
			goto tr498
		case 61:
			goto st97
		case 62:
			goto tr500
		case 63:
			goto tr501
		case 64:
			goto tr502
		case 66:
			goto tr504
		case 69:
			goto tr505
		case 88:
			goto tr506
		case 95:
			goto tr507
		case 98:
			goto tr509
		case 101:
			goto tr510
		case 120:
			goto tr511
		case 124:
			goto tr512
		case 126:
			goto tr513
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr488
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr483
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr503
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr488
					}
				case data[p] >= 97:
					goto tr508
				}
			default:
				goto tr488
			}
		default:
			goto tr496
		}
		goto tr514
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch data[p] {
		case 32:
			goto tr515
		case 33:
			goto tr516
		case 34:
			goto tr517
		case 35:
			goto tr518
		case 36:
			goto tr519
		case 38:
			goto tr521
		case 39:
			goto tr522
		case 43:
			goto tr523
		case 45:
			goto tr524
		case 46:
			goto tr525
		case 47:
			goto tr526
		case 48:
			goto tr527
		case 58:
			goto tr529
		case 60:
			goto tr530
		case 62:
			goto tr531
		case 63:
			goto tr532
		case 64:
			goto tr533
		case 66:
			goto tr535
		case 69:
			goto tr536
		case 88:
			goto tr537
		case 95:
			goto tr538
		case 98:
			goto tr540
		case 101:
			goto tr541
		case 120:
			goto tr542
		case 124:
			goto tr543
		case 126:
			goto tr544
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr520
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr515
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr534
					}
				case data[p] >= 59:
					goto tr520
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr520
					}
				case data[p] >= 97:
					goto tr539
				}
			default:
				goto tr520
			}
		default:
			goto tr528
		}
		goto tr545
tr173:
//line lex.rl:56
 mark = p 
	goto st98
tr205:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st98
tr236:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st98
tr272:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st98
tr303:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st98
tr336:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st98
tr367:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st98
tr399:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st98
tr433:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st98
tr464:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st98
tr500:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st98
tr531:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st98
tr564:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st98
tr598:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st98
tr630:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st98
tr659:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st98
tr683:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st98
tr714:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st98
tr746:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st98
tr774:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st98
tr796:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st98
tr828:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st98
tr859:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st98
tr890:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st98
tr922:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st98
tr953:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st98
tr984:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st98
tr1015:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st98
tr1046:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st98
tr1077:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st98
tr1135:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st98
tr1164:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st98
tr1195:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st98
tr1257:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st98
tr1288:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st98
tr1320:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st98
tr1351:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st98
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
//line lex.go:9549
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto tr194
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto tr197
		case 46:
			goto tr198
		case 47:
			goto tr199
		case 48:
			goto tr200
		case 58:
			goto tr202
		case 59:
			goto tr193
		case 60:
			goto tr203
		case 61:
			goto st99
		case 62:
			goto st120
		case 63:
			goto tr206
		case 64:
			goto tr207
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto tr217
		case 126:
			goto tr253
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr208
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto tr201
		}
		goto tr219
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch data[p] {
		case 32:
			goto tr548
		case 33:
			goto tr549
		case 34:
			goto tr550
		case 35:
			goto tr551
		case 36:
			goto tr552
		case 38:
			goto tr554
		case 39:
			goto tr555
		case 43:
			goto tr556
		case 45:
			goto tr557
		case 46:
			goto tr558
		case 47:
			goto tr559
		case 48:
			goto tr560
		case 58:
			goto tr562
		case 60:
			goto tr563
		case 62:
			goto tr564
		case 63:
			goto tr565
		case 64:
			goto tr566
		case 66:
			goto tr568
		case 69:
			goto tr569
		case 88:
			goto tr570
		case 95:
			goto tr571
		case 98:
			goto tr573
		case 101:
			goto tr574
		case 120:
			goto tr575
		case 124:
			goto tr576
		case 126:
			goto tr577
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr553
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr548
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr567
					}
				case data[p] >= 59:
					goto tr553
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr553
					}
				case data[p] >= 97:
					goto tr572
				}
			default:
				goto tr553
			}
		default:
			goto tr561
		}
		goto tr578
tr174:
//line lex.rl:56
 mark = p 
	goto st100
tr206:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st100
tr237:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st100
tr273:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st100
tr304:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st100
tr337:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st100
tr368:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st100
tr400:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st100
tr434:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st100
tr465:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st100
tr501:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st100
tr532:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st100
tr565:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st100
tr599:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st100
tr631:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st100
tr660:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st100
tr684:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st100
tr715:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st100
tr747:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st100
tr775:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st100
tr797:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st100
tr829:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st100
tr860:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st100
tr891:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st100
tr923:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st100
tr954:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st100
tr985:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st100
tr1016:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st100
tr1047:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st100
tr1078:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st100
tr1105:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st100
tr1136:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st100
tr1165:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st100
tr1196:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st100
tr1227:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st100
tr1258:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st100
tr1289:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st100
tr1321:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st100
tr1352:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st100
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
//line lex.go:10132
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto st101
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto tr197
		case 46:
			goto tr198
		case 47:
			goto tr199
		case 48:
			goto tr200
		case 58:
			goto tr202
		case 60:
			goto tr203
		case 62:
			goto tr205
		case 63:
			goto st118
		case 64:
			goto tr207
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto st119
		case 126:
			goto tr253
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr208
					}
				case data[p] >= 59:
					goto tr193
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto tr201
		}
		goto tr219
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 32:
			goto tr582
		case 33:
			goto tr583
		case 34:
			goto tr584
		case 35:
			goto tr585
		case 36:
			goto tr586
		case 38:
			goto tr588
		case 39:
			goto tr589
		case 43:
			goto tr590
		case 45:
			goto tr591
		case 46:
			goto tr592
		case 47:
			goto tr593
		case 48:
			goto tr594
		case 58:
			goto tr596
		case 60:
			goto tr597
		case 62:
			goto tr598
		case 63:
			goto tr599
		case 64:
			goto tr600
		case 66:
			goto tr602
		case 69:
			goto tr603
		case 88:
			goto tr604
		case 95:
			goto tr605
		case 98:
			goto tr607
		case 101:
			goto tr608
		case 120:
			goto tr609
		case 124:
			goto tr610
		case 126:
			goto tr611
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr587
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr582
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr601
					}
				case data[p] >= 59:
					goto tr587
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr587
					}
				case data[p] >= 97:
					goto tr606
				}
			default:
				goto tr587
			}
		default:
			goto tr595
		}
		goto tr612
tr175:
//line lex.rl:56
 mark = p 
	goto st102
tr207:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st102
tr238:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st102
tr274:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st102
tr305:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st102
tr338:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st102
tr369:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st102
tr401:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st102
tr435:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st102
tr466:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st102
tr502:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st102
tr533:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st102
tr566:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st102
tr600:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st102
tr632:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st102
tr661:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st102
tr685:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st102
tr716:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st102
tr748:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st102
tr776:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st102
tr798:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st102
tr830:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st102
tr861:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st102
tr892:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st102
tr924:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st102
tr955:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st102
tr986:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st102
tr1017:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st102
tr1048:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st102
tr1079:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st102
tr1106:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st102
tr1137:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st102
tr1166:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st102
tr1197:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st102
tr1228:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st102
tr1259:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st102
tr1290:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st102
tr1322:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st102
tr1353:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st102
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
//line lex.go:10716
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto tr194
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto tr197
		case 46:
			goto tr198
		case 47:
			goto tr199
		case 48:
			goto tr200
		case 58:
			goto tr202
		case 60:
			goto tr203
		case 62:
			goto st103
		case 63:
			goto tr206
		case 64:
			goto tr207
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto tr217
		case 126:
			goto tr253
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr208
					}
				case data[p] >= 59:
					goto tr193
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto tr201
		}
		goto tr219
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 32:
			goto tr614
		case 33:
			goto tr615
		case 34:
			goto tr616
		case 35:
			goto tr617
		case 36:
			goto tr618
		case 38:
			goto tr620
		case 39:
			goto tr621
		case 43:
			goto tr622
		case 45:
			goto tr623
		case 46:
			goto tr624
		case 47:
			goto tr625
		case 48:
			goto tr626
		case 58:
			goto tr628
		case 60:
			goto tr629
		case 62:
			goto tr630
		case 63:
			goto tr631
		case 64:
			goto tr632
		case 66:
			goto tr634
		case 69:
			goto tr635
		case 88:
			goto tr636
		case 95:
			goto tr637
		case 98:
			goto tr639
		case 101:
			goto tr640
		case 120:
			goto tr641
		case 124:
			goto tr642
		case 126:
			goto tr643
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr619
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr614
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr633
					}
				case data[p] >= 59:
					goto tr619
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr619
					}
				case data[p] >= 97:
					goto tr638
				}
			default:
				goto tr619
			}
		default:
			goto tr627
		}
		goto tr644
tr665:
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr181:
//line lex.rl:56
 mark = p 
	goto st104
tr176:
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr187:
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr213:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st104
tr208:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr219:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr244:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st104
tr239:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr250:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr280:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st104
tr275:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr286:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr311:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st104
tr306:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr317:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr344:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st104
tr339:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr349:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr375:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st104
tr370:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr381:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr407:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st104
tr402:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr413:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr441:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st104
tr436:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr447:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr472:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st104
tr467:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr478:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr508:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st104
tr503:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr514:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr539:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st104
tr534:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr545:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr572:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st104
tr567:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr578:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr606:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st104
tr601:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr612:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr638:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st104
tr633:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr644:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr662:
//line lex.rl:134
 isUpper = true 
	goto st104
tr691:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st104
tr686:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr697:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr722:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st104
tr717:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr728:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr754:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st104
tr749:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr760:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr804:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st104
tr799:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr810:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr836:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st104
tr831:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr842:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr867:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st104
tr862:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr873:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr898:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st104
tr893:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr904:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr930:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st104
tr925:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr936:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr961:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st104
tr956:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr967:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr992:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st104
tr987:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr998:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr1023:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st104
tr1018:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr1029:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr1054:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st104
tr1049:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr1060:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr1083:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st104
tr1080:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr1087:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr1112:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st104
tr1107:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr1118:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr1143:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st104
tr1138:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr1149:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr1172:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st104
tr1167:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr1178:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr1203:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st104
tr1198:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr1209:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr1234:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st104
tr1229:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr1240:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr1265:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st104
tr1260:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr1271:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr1296:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st104
tr1291:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr1302:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr1328:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st104
tr1323:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr1334:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
tr1359:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st104
tr1354:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st104
tr1365:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:131
 isNotASCII = true 
	goto st104
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
//line lex.go:12008
		switch data[p] {
		case 32:
			goto tr645
		case 33:
			goto tr646
		case 34:
			goto tr647
		case 35:
			goto tr648
		case 36:
			goto st104
		case 38:
			goto tr651
		case 39:
			goto tr652
		case 43:
			goto tr653
		case 45:
			goto tr654
		case 46:
			goto tr655
		case 47:
			goto tr656
		case 58:
			goto tr657
		case 60:
			goto tr658
		case 62:
			goto tr659
		case 63:
			goto tr660
		case 64:
			goto tr661
		case 95:
			goto st104
		case 124:
			goto tr663
		case 126:
			goto tr664
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr650
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr645
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr662
					}
				case data[p] >= 59:
					goto tr650
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr650
					}
				case data[p] >= 97:
					goto st104
				}
			default:
				goto tr650
			}
		default:
			goto st104
		}
		goto tr665
tr185:
//line lex.rl:56
 mark = p 
	goto st105
tr217:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st105
tr248:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st105
tr284:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st105
tr315:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st105
tr347:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st105
tr379:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st105
tr411:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st105
tr445:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st105
tr476:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st105
tr512:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st105
tr543:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st105
tr576:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st105
tr610:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st105
tr642:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st105
tr663:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st105
tr695:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st105
tr726:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st105
tr758:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st105
tr777:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st105
tr808:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st105
tr840:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st105
tr871:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st105
tr902:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st105
tr934:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st105
tr965:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st105
tr996:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st105
tr1027:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st105
tr1058:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st105
tr1085:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st105
tr1116:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st105
tr1147:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st105
tr1176:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st105
tr1207:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st105
tr1238:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st105
tr1269:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st105
tr1300:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st105
tr1332:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st105
tr1363:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st105
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
//line lex.go:12471
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto tr194
		case 39:
			goto tr195
		case 43:
			goto tr196
		case 45:
			goto tr197
		case 46:
			goto tr198
		case 47:
			goto tr199
		case 48:
			goto tr200
		case 58:
			goto tr202
		case 60:
			goto tr203
		case 62:
			goto tr205
		case 63:
			goto tr206
		case 64:
			goto tr207
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto st117
		case 126:
			goto tr253
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr208
					}
				case data[p] >= 59:
					goto tr193
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto tr201
		}
		goto tr219
tr177:
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr209:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr240:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr276:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr307:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr340:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr371:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr403:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr437:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr468:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr504:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr535:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr568:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr602:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr634:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr687:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr718:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr750:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr800:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr832:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr863:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr894:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr926:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr957:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr988:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr1019:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr1050:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr1108:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr1139:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr1168:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr1199:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr1230:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr1261:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr1292:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr1324:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
tr1355:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st106
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
//line lex.go:12950
		switch data[p] {
		case 32:
			goto tr645
		case 33:
			goto tr646
		case 34:
			goto tr647
		case 35:
			goto tr648
		case 36:
			goto st104
		case 38:
			goto tr651
		case 39:
			goto st6
		case 43:
			goto tr653
		case 45:
			goto tr654
		case 46:
			goto tr655
		case 47:
			goto tr656
		case 58:
			goto tr657
		case 60:
			goto tr658
		case 62:
			goto tr659
		case 63:
			goto tr660
		case 64:
			goto tr661
		case 95:
			goto st104
		case 124:
			goto tr663
		case 126:
			goto tr664
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr650
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr645
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr662
					}
				case data[p] >= 59:
					goto tr650
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr650
					}
				case data[p] >= 97:
					goto st104
				}
			default:
				goto tr650
			}
		default:
			goto st104
		}
		goto tr665
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 39 {
			goto st107
		}
		if 48 <= data[p] && data[p] <= 49 {
			goto st6
		}
		goto tr14
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 32:
			goto tr667
		case 33:
			goto tr668
		case 34:
			goto tr669
		case 35:
			goto tr670
		case 36:
			goto tr671
		case 38:
			goto tr673
		case 39:
			goto tr674
		case 43:
			goto tr675
		case 45:
			goto tr676
		case 46:
			goto tr677
		case 47:
			goto tr678
		case 48:
			goto tr679
		case 58:
			goto tr681
		case 60:
			goto tr682
		case 62:
			goto tr683
		case 63:
			goto tr684
		case 64:
			goto tr685
		case 66:
			goto tr687
		case 69:
			goto tr688
		case 88:
			goto tr689
		case 95:
			goto tr690
		case 98:
			goto tr692
		case 101:
			goto tr693
		case 120:
			goto tr694
		case 124:
			goto tr695
		case 126:
			goto tr696
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr672
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr667
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr686
					}
				case data[p] >= 59:
					goto tr672
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr672
					}
				case data[p] >= 97:
					goto tr691
				}
			default:
				goto tr672
			}
		default:
			goto tr680
		}
		goto tr697
tr183:
//line lex.rl:56
 mark = p 
	goto st108
tr178:
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr215:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st108
tr210:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr246:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st108
tr241:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr282:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st108
tr277:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr313:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st108
tr308:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr1062:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st108
tr1061:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr377:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st108
tr372:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr409:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st108
tr404:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr443:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st108
tr438:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr474:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st108
tr469:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr510:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st108
tr505:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr541:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st108
tr536:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr574:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st108
tr569:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr608:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st108
tr603:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr640:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st108
tr635:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr693:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st108
tr688:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr724:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st108
tr719:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr756:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st108
tr751:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr806:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st108
tr801:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr838:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st108
tr833:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr869:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st108
tr864:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr900:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st108
tr895:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr932:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st108
tr927:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr963:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st108
tr958:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr994:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st108
tr989:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr1025:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st108
tr1020:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr1056:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st108
tr1051:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr1114:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st108
tr1109:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr1145:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st108
tr1140:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr1174:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st108
tr1169:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr1205:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st108
tr1200:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr1236:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st108
tr1231:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr1267:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st108
tr1262:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr1298:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st108
tr1293:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr1330:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st108
tr1325:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
tr1361:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st108
tr1356:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
//line lex.go:13833
		switch data[p] {
		case 32:
			goto tr645
		case 33:
			goto tr646
		case 34:
			goto tr647
		case 35:
			goto tr648
		case 36:
			goto st104
		case 38:
			goto tr651
		case 39:
			goto st7
		case 43:
			goto tr653
		case 45:
			goto tr654
		case 46:
			goto tr655
		case 47:
			goto tr656
		case 58:
			goto tr657
		case 60:
			goto tr658
		case 62:
			goto tr659
		case 63:
			goto tr660
		case 64:
			goto tr661
		case 95:
			goto st104
		case 124:
			goto tr663
		case 126:
			goto tr664
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr650
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr645
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr662
					}
				case data[p] >= 59:
					goto tr650
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr650
					}
				case data[p] >= 97:
					goto st104
				}
			default:
				goto tr650
			}
		default:
			goto st104
		}
		goto tr665
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 39:
			goto tr18
		case 92:
			goto tr19
		}
		goto tr17
tr17:
//line lex.rl:245
 buf = new(bytes.Buffer) 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st8
tr20:
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st8
tr94:
//line lex.rl:256
 buf.WriteByte(data[p]) 
	goto st8
tr108:
//line lex.rl:242
 buf.WriteByte(ch) 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st8
tr119:
//line lex.rl:236
 buf.WriteRune(rn) 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st8
tr124:
//line lex.rl:225
 buf.WriteByte(ch) 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st8
tr127:
//line lex.rl:213
 buf.WriteByte('\a') 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st8
tr130:
//line lex.rl:214
 buf.WriteByte('\b') 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st8
tr133:
//line lex.rl:215
 buf.WriteByte('\f') 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st8
tr136:
//line lex.rl:216
 buf.WriteByte('\n') 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st8
tr139:
//line lex.rl:217
 buf.WriteByte('\r') 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st8
tr142:
//line lex.rl:218
 buf.WriteByte('\t') 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st8
tr149:
//line lex.rl:219
 buf.WriteByte('\v') 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line lex.go:14013
		switch data[p] {
		case 39:
			goto st109
		case 92:
			goto st41
		}
		goto tr20
tr18:
//line lex.rl:245
 buf = new(bytes.Buffer) 
	goto st109
tr109:
//line lex.rl:242
 buf.WriteByte(ch) 
	goto st109
tr120:
//line lex.rl:236
 buf.WriteRune(rn) 
	goto st109
tr125:
//line lex.rl:225
 buf.WriteByte(ch) 
	goto st109
tr128:
//line lex.rl:213
 buf.WriteByte('\a') 
	goto st109
tr131:
//line lex.rl:214
 buf.WriteByte('\b') 
	goto st109
tr134:
//line lex.rl:215
 buf.WriteByte('\f') 
	goto st109
tr137:
//line lex.rl:216
 buf.WriteByte('\n') 
	goto st109
tr140:
//line lex.rl:217
 buf.WriteByte('\r') 
	goto st109
tr143:
//line lex.rl:218
 buf.WriteByte('\t') 
	goto st109
tr150:
//line lex.rl:219
 buf.WriteByte('\v') 
	goto st109
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
//line lex.go:14070
		switch data[p] {
		case 32:
			goto tr699
		case 33:
			goto tr700
		case 34:
			goto tr701
		case 35:
			goto tr702
		case 36:
			goto tr703
		case 38:
			goto tr705
		case 39:
			goto tr20
		case 43:
			goto tr706
		case 45:
			goto tr707
		case 46:
			goto tr708
		case 47:
			goto tr709
		case 48:
			goto tr710
		case 58:
			goto tr712
		case 60:
			goto tr713
		case 62:
			goto tr714
		case 63:
			goto tr715
		case 64:
			goto tr716
		case 66:
			goto tr718
		case 69:
			goto tr719
		case 88:
			goto tr720
		case 95:
			goto tr721
		case 98:
			goto tr723
		case 101:
			goto tr724
		case 120:
			goto tr725
		case 124:
			goto tr726
		case 126:
			goto tr727
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr704
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr699
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr717
					}
				case data[p] >= 59:
					goto tr704
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr704
					}
				case data[p] >= 97:
					goto tr722
				}
			default:
				goto tr704
			}
		default:
			goto tr711
		}
		goto tr728
tr184:
//line lex.rl:56
 mark = p 
	goto st110
tr179:
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr216:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st110
tr211:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr247:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st110
tr242:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr283:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st110
tr278:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr314:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st110
tr309:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr346:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st110
tr342:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr378:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st110
tr373:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr410:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st110
tr405:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr444:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st110
tr439:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr475:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st110
tr470:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr511:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st110
tr506:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr542:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st110
tr537:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr575:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st110
tr570:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr609:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st110
tr604:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr641:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st110
tr636:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr694:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st110
tr689:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr725:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st110
tr720:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr757:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st110
tr752:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr807:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st110
tr802:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr839:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st110
tr834:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr870:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st110
tr865:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr901:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st110
tr896:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr933:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st110
tr928:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr964:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st110
tr959:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr995:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st110
tr990:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr1026:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st110
tr1021:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr1057:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st110
tr1052:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr1084:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st110
tr1081:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr1115:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st110
tr1110:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr1146:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st110
tr1141:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr1175:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st110
tr1170:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr1206:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st110
tr1201:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr1237:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st110
tr1232:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr1268:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st110
tr1263:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr1299:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st110
tr1294:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr1331:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st110
tr1326:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
tr1362:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st110
tr1357:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
//line lex.rl:134
 isUpper = true 
	goto st110
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
//line lex.go:14866
		switch data[p] {
		case 32:
			goto tr645
		case 33:
			goto tr646
		case 34:
			goto tr647
		case 35:
			goto tr648
		case 36:
			goto st104
		case 38:
			goto tr651
		case 39:
			goto tr729
		case 43:
			goto tr653
		case 45:
			goto tr654
		case 46:
			goto tr655
		case 47:
			goto tr656
		case 58:
			goto tr657
		case 60:
			goto tr658
		case 62:
			goto tr659
		case 63:
			goto tr660
		case 64:
			goto tr661
		case 95:
			goto st104
		case 124:
			goto tr663
		case 126:
			goto tr664
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr650
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr645
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr662
					}
				case data[p] >= 59:
					goto tr650
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr650
					}
				case data[p] >= 97:
					goto st104
				}
			default:
				goto tr650
			}
		default:
			goto st104
		}
		goto tr665
tr729:
//line lex.rl:278
 buf = new(bytes.Buffer) 
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line lex.go:14964
		if data[p] == 39 {
			goto st111
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr25
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto tr23
tr92:
//line lex.rl:284
 buf.WriteByte(ch) 
	goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
//line lex.go:14990
		switch data[p] {
		case 32:
			goto tr730
		case 33:
			goto tr731
		case 34:
			goto tr732
		case 35:
			goto tr733
		case 36:
			goto tr734
		case 38:
			goto tr736
		case 39:
			goto tr737
		case 43:
			goto tr738
		case 45:
			goto tr739
		case 46:
			goto tr740
		case 47:
			goto tr741
		case 48:
			goto tr742
		case 58:
			goto tr744
		case 60:
			goto tr745
		case 62:
			goto tr746
		case 63:
			goto tr747
		case 64:
			goto tr748
		case 66:
			goto tr750
		case 69:
			goto tr751
		case 88:
			goto tr752
		case 95:
			goto tr753
		case 98:
			goto tr755
		case 101:
			goto tr756
		case 120:
			goto tr757
		case 124:
			goto tr758
		case 126:
			goto tr759
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr735
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr730
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr749
					}
				case data[p] >= 59:
					goto tr735
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr735
					}
				case data[p] >= 97:
					goto tr754
				}
			default:
				goto tr735
			}
		default:
			goto tr743
		}
		goto tr760
tr180:
//line lex.rl:56
 mark = p 
	goto st112
tr212:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st112
tr243:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st112
tr279:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st112
tr310:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st112
tr343:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st112
tr374:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st112
tr406:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st112
tr440:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st112
tr471:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st112
tr507:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st112
tr538:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st112
tr571:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st112
tr605:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st112
tr637:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st112
tr690:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st112
tr721:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st112
tr753:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st112
tr803:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st112
tr835:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st112
tr866:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st112
tr897:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st112
tr929:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st112
tr960:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st112
tr991:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st112
tr1022:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st112
tr1053:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st112
tr1082:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st112
tr1111:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st112
tr1142:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st112
tr1171:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st112
tr1202:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st112
tr1233:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st112
tr1264:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st112
tr1295:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st112
tr1327:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st112
tr1358:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st112
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
//line lex.go:15405
		switch data[p] {
		case 32:
			goto tr761
		case 33:
			goto tr762
		case 34:
			goto tr763
		case 35:
			goto tr764
		case 36:
			goto st104
		case 38:
			goto tr766
		case 39:
			goto tr767
		case 43:
			goto tr768
		case 45:
			goto tr769
		case 46:
			goto tr770
		case 47:
			goto tr771
		case 58:
			goto tr772
		case 60:
			goto tr773
		case 62:
			goto tr774
		case 63:
			goto tr775
		case 64:
			goto tr776
		case 95:
			goto st104
		case 124:
			goto tr777
		case 126:
			goto tr778
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr765
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr761
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr662
					}
				case data[p] >= 59:
					goto tr765
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr765
					}
				case data[p] >= 97:
					goto st104
				}
			default:
				goto tr765
			}
		default:
			goto st104
		}
		goto tr665
tr186:
//line lex.rl:56
 mark = p 
	goto st113
tr253:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st113
tr249:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st113
tr285:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st113
tr316:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st113
tr348:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st113
tr380:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st113
tr412:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st113
tr446:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st113
tr477:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st113
tr513:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st113
tr544:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st113
tr577:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st113
tr611:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st113
tr643:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st113
tr664:
//line lex.rl:97

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
		
//line lex.rl:56
 mark = p 
	goto st113
tr696:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st113
tr727:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st113
tr759:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st113
tr778:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st113
tr809:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st113
tr841:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st113
tr872:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st113
tr903:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st113
tr935:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st113
tr966:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st113
tr997:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st113
tr1028:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st113
tr1059:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st113
tr1086:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:56
 mark = p 
	goto st113
tr1117:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st113
tr1148:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st113
tr1177:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st113
tr1208:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st113
tr1239:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st113
tr1270:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st113
tr1301:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st113
tr1333:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st113
tr1364:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st113
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
//line lex.go:15868
		switch data[p] {
		case 32:
			goto tr188
		case 33:
			goto tr189
		case 34:
			goto tr190
		case 35:
			goto tr191
		case 36:
			goto tr192
		case 38:
			goto tr194
		case 39:
			goto tr195
		case 42:
			goto st114
		case 43:
			goto tr196
		case 45:
			goto tr197
		case 46:
			goto tr198
		case 47:
			goto tr199
		case 48:
			goto tr200
		case 58:
			goto tr202
		case 60:
			goto tr203
		case 62:
			goto tr205
		case 63:
			goto tr206
		case 64:
			goto tr207
		case 66:
			goto tr209
		case 69:
			goto tr210
		case 88:
			goto tr211
		case 95:
			goto tr212
		case 98:
			goto tr214
		case 101:
			goto tr215
		case 120:
			goto tr216
		case 124:
			goto tr217
		case 126:
			goto tr253
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr193
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr188
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr208
					}
				case data[p] >= 59:
					goto tr193
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr193
					}
				case data[p] >= 97:
					goto tr213
				}
			default:
				goto tr193
			}
		default:
			goto tr201
		}
		goto tr219
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 32:
			goto tr780
		case 33:
			goto tr781
		case 34:
			goto tr782
		case 35:
			goto tr783
		case 36:
			goto tr784
		case 38:
			goto tr786
		case 39:
			goto tr787
		case 43:
			goto tr788
		case 45:
			goto tr789
		case 46:
			goto tr790
		case 47:
			goto tr791
		case 48:
			goto tr792
		case 58:
			goto tr794
		case 60:
			goto tr795
		case 62:
			goto tr796
		case 63:
			goto tr797
		case 64:
			goto tr798
		case 66:
			goto tr800
		case 69:
			goto tr801
		case 88:
			goto tr802
		case 95:
			goto tr803
		case 98:
			goto tr805
		case 101:
			goto tr806
		case 120:
			goto tr807
		case 124:
			goto tr808
		case 126:
			goto tr809
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr785
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr780
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr799
					}
				case data[p] >= 59:
					goto tr785
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr785
					}
				case data[p] >= 97:
					goto tr804
				}
			default:
				goto tr785
			}
		default:
			goto tr793
		}
		goto tr810
tr182:
//line lex.rl:56
 mark = p 
	goto st115
tr214:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:56
 mark = p 
	goto st115
tr245:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:56
 mark = p 
	goto st115
tr281:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st115
tr312:
//line lex.rl:186

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
		
//line lex.rl:56
 mark = p 
	goto st115
tr345:
//line lex.rl:73

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
		
//line lex.rl:56
 mark = p 
	goto st115
tr376:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:56
 mark = p 
	goto st115
tr408:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:56
 mark = p 
	goto st115
tr442:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:56
 mark = p 
	goto st115
tr473:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:56
 mark = p 
	goto st115
tr509:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st115
tr540:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st115
tr573:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st115
tr607:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st115
tr639:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:56
 mark = p 
	goto st115
tr692:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:56
 mark = p 
	goto st115
tr723:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st115
tr755:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:56
 mark = p 
	goto st115
tr805:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st115
tr837:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:56
 mark = p 
	goto st115
tr868:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:56
 mark = p 
	goto st115
tr899:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:56
 mark = p 
	goto st115
tr931:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:56
 mark = p 
	goto st115
tr962:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st115
tr993:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st115
tr1024:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st115
tr1055:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:56
 mark = p 
	goto st115
tr1113:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:56
 mark = p 
	goto st115
tr1144:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:56
 mark = p 
	goto st115
tr1173:
//line lex.rl:59

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
		
//line lex.rl:56
 mark = p 
	goto st115
tr1204:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:56
 mark = p 
	goto st115
tr1235:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:56
 mark = p 
	goto st115
tr1266:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:56
 mark = p 
	goto st115
tr1297:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:56
 mark = p 
	goto st115
tr1329:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:56
 mark = p 
	goto st115
tr1360:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:56
 mark = p 
	goto st115
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
//line lex.go:16384
		switch data[p] {
		case 32:
			goto tr645
		case 33:
			goto tr646
		case 34:
			goto tr647
		case 35:
			goto tr648
		case 36:
			goto st104
		case 38:
			goto tr651
		case 39:
			goto st10
		case 43:
			goto tr653
		case 45:
			goto tr654
		case 46:
			goto tr655
		case 47:
			goto tr656
		case 58:
			goto tr657
		case 60:
			goto tr658
		case 62:
			goto tr659
		case 63:
			goto tr660
		case 64:
			goto tr661
		case 95:
			goto st104
		case 124:
			goto tr663
		case 126:
			goto tr664
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr650
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr645
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr662
					}
				case data[p] >= 59:
					goto tr650
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr650
					}
				case data[p] >= 97:
					goto st104
				}
			default:
				goto tr650
			}
		default:
			goto st104
		}
		goto tr665
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 39:
			goto tr27
		case 92:
			goto tr28
		}
		goto tr26
tr26:
//line lex.rl:245
 buf = new(bytes.Buffer) 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st11
tr29:
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st11
tr32:
//line lex.rl:256
 buf.WriteByte(data[p]) 
	goto st11
tr47:
//line lex.rl:242
 buf.WriteByte(ch) 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st11
tr58:
//line lex.rl:236
 buf.WriteRune(rn) 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st11
tr63:
//line lex.rl:225
 buf.WriteByte(ch) 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st11
tr66:
//line lex.rl:213
 buf.WriteByte('\a') 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st11
tr69:
//line lex.rl:214
 buf.WriteByte('\b') 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st11
tr72:
//line lex.rl:215
 buf.WriteByte('\f') 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st11
tr75:
//line lex.rl:216
 buf.WriteByte('\n') 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st11
tr78:
//line lex.rl:217
 buf.WriteByte('\r') 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st11
tr81:
//line lex.rl:218
 buf.WriteByte('\t') 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st11
tr88:
//line lex.rl:219
 buf.WriteByte('\v') 
//line lex.rl:250
 buf.WriteByte(data[p]) 
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line lex.go:16564
		switch data[p] {
		case 39:
			goto st116
		case 92:
			goto st12
		}
		goto tr29
tr27:
//line lex.rl:245
 buf = new(bytes.Buffer) 
	goto st116
tr48:
//line lex.rl:242
 buf.WriteByte(ch) 
	goto st116
tr59:
//line lex.rl:236
 buf.WriteRune(rn) 
	goto st116
tr64:
//line lex.rl:225
 buf.WriteByte(ch) 
	goto st116
tr67:
//line lex.rl:213
 buf.WriteByte('\a') 
	goto st116
tr70:
//line lex.rl:214
 buf.WriteByte('\b') 
	goto st116
tr73:
//line lex.rl:215
 buf.WriteByte('\f') 
	goto st116
tr76:
//line lex.rl:216
 buf.WriteByte('\n') 
	goto st116
tr79:
//line lex.rl:217
 buf.WriteByte('\r') 
	goto st116
tr82:
//line lex.rl:218
 buf.WriteByte('\t') 
	goto st116
tr89:
//line lex.rl:219
 buf.WriteByte('\v') 
	goto st116
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
//line lex.go:16621
		switch data[p] {
		case 32:
			goto tr730
		case 33:
			goto tr731
		case 34:
			goto tr732
		case 35:
			goto tr733
		case 36:
			goto tr734
		case 38:
			goto tr736
		case 39:
			goto tr29
		case 43:
			goto tr738
		case 45:
			goto tr739
		case 46:
			goto tr740
		case 47:
			goto tr741
		case 48:
			goto tr742
		case 58:
			goto tr744
		case 60:
			goto tr745
		case 62:
			goto tr746
		case 63:
			goto tr747
		case 64:
			goto tr748
		case 66:
			goto tr750
		case 69:
			goto tr751
		case 88:
			goto tr752
		case 95:
			goto tr753
		case 98:
			goto tr755
		case 101:
			goto tr756
		case 120:
			goto tr757
		case 124:
			goto tr758
		case 126:
			goto tr759
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr735
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr730
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr749
					}
				case data[p] >= 59:
					goto tr735
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr735
					}
				case data[p] >= 97:
					goto tr754
				}
			default:
				goto tr735
			}
		default:
			goto tr743
		}
		goto tr760
tr28:
//line lex.rl:245
 buf = new(bytes.Buffer) 
	goto st12
tr49:
//line lex.rl:242
 buf.WriteByte(ch) 
	goto st12
tr60:
//line lex.rl:236
 buf.WriteRune(rn) 
	goto st12
tr65:
//line lex.rl:225
 buf.WriteByte(ch) 
	goto st12
tr68:
//line lex.rl:213
 buf.WriteByte('\a') 
	goto st12
tr71:
//line lex.rl:214
 buf.WriteByte('\b') 
	goto st12
tr74:
//line lex.rl:215
 buf.WriteByte('\f') 
	goto st12
tr77:
//line lex.rl:216
 buf.WriteByte('\n') 
	goto st12
tr80:
//line lex.rl:217
 buf.WriteByte('\r') 
	goto st12
tr83:
//line lex.rl:218
 buf.WriteByte('\t') 
	goto st12
tr90:
//line lex.rl:219
 buf.WriteByte('\v') 
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line lex.go:16773
		switch data[p] {
		case 85:
			goto tr34
		case 88:
			goto st25
		case 97:
			goto st28
		case 98:
			goto st29
		case 102:
			goto st30
		case 110:
			goto st31
		case 114:
			goto st32
		case 116:
			goto st33
		case 117:
			goto tr42
		case 118:
			goto st38
		case 120:
			goto st25
		}
		if 48 <= data[p] && data[p] <= 55 {
			goto tr33
		}
		goto tr32
tr33:
//line lex.rl:240
 ch = 0 
//line lex.rl:241
 ch = (ch << 3) | data[p] - '0' 
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line lex.go:16813
		if 48 <= data[p] && data[p] <= 55 {
			goto tr45
		}
		goto tr44
tr45:
//line lex.rl:241
 ch = (ch << 3) | data[p] - '0' 
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line lex.go:16827
		if 48 <= data[p] && data[p] <= 55 {
			goto tr46
		}
		goto tr44
tr46:
//line lex.rl:241
 ch = (ch << 3) | data[p] - '0' 
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line lex.go:16841
		switch data[p] {
		case 39:
			goto tr48
		case 92:
			goto tr49
		}
		goto tr47
tr34:
//line lex.rl:235
 rn = 0 
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line lex.go:16858
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr50
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr50
			}
		default:
			goto tr50
		}
		goto tr44
tr50:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line lex.go:16881
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr51
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr51
			}
		default:
			goto tr51
		}
		goto tr44
tr51:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line lex.go:16904
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr52
			}
		default:
			goto tr52
		}
		goto tr44
tr52:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line lex.go:16927
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr53
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr53
			}
		default:
			goto tr53
		}
		goto tr44
tr53:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line lex.go:16950
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr54
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr54
			}
		default:
			goto tr54
		}
		goto tr44
tr54:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line lex.go:16973
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr55
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr55
			}
		default:
			goto tr55
		}
		goto tr44
tr55:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line lex.go:16996
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr56
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr56
			}
		default:
			goto tr56
		}
		goto tr44
tr56:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line lex.go:17019
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr57
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr57
			}
		default:
			goto tr57
		}
		goto tr44
tr57:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st24
tr87:
//line lex.rl:230
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line lex.go:17046
		switch data[p] {
		case 39:
			goto tr59
		case 92:
			goto tr60
		}
		goto tr58
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr61
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr61
			}
		default:
			goto tr61
		}
		goto tr44
tr61:
//line lex.rl:223
 ch = 0 
//line lex.rl:224
 ch = (ch << 4) | unhex(data[p]) 
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line lex.go:17083
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr62
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr62
			}
		default:
			goto tr62
		}
		goto tr44
tr62:
//line lex.rl:224
 ch = (ch << 4) | unhex(data[p]) 
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line lex.go:17106
		switch data[p] {
		case 39:
			goto tr64
		case 92:
			goto tr65
		}
		goto tr63
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 39:
			goto tr67
		case 92:
			goto tr68
		}
		goto tr66
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 39:
			goto tr70
		case 92:
			goto tr71
		}
		goto tr69
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 39:
			goto tr73
		case 92:
			goto tr74
		}
		goto tr72
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 39:
			goto tr76
		case 92:
			goto tr77
		}
		goto tr75
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 39:
			goto tr79
		case 92:
			goto tr80
		}
		goto tr78
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 39:
			goto tr82
		case 92:
			goto tr83
		}
		goto tr81
tr42:
//line lex.rl:235
 rn = 0 
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line lex.go:17195
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr84
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr84
			}
		default:
			goto tr84
		}
		goto tr44
tr84:
//line lex.rl:230
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line lex.go:17218
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr85
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr85
			}
		default:
			goto tr85
		}
		goto tr44
tr85:
//line lex.rl:230
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line lex.go:17241
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr86
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr86
			}
		default:
			goto tr86
		}
		goto tr44
tr86:
//line lex.rl:230
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line lex.go:17264
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr87
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr87
			}
		default:
			goto tr87
		}
		goto tr44
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 39:
			goto tr89
		case 92:
			goto tr90
		}
		goto tr88
tr25:
//line lex.rl:282
 ch = 0 
//line lex.rl:283
 ch = (ch << 4) | unhex(data[p]) 
	goto st39
tr93:
//line lex.rl:284
 buf.WriteByte(ch) 
//line lex.rl:282
 ch = 0 
//line lex.rl:283
 ch = (ch << 4) | unhex(data[p]) 
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line lex.go:17309
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr91
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr91
			}
		default:
			goto tr91
		}
		goto tr23
tr91:
//line lex.rl:283
 ch = (ch << 4) | unhex(data[p]) 
	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line lex.go:17332
		if data[p] == 39 {
			goto tr92
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr93
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr93
			}
		default:
			goto tr93
		}
		goto tr23
tr19:
//line lex.rl:245
 buf = new(bytes.Buffer) 
	goto st41
tr110:
//line lex.rl:242
 buf.WriteByte(ch) 
	goto st41
tr121:
//line lex.rl:236
 buf.WriteRune(rn) 
	goto st41
tr126:
//line lex.rl:225
 buf.WriteByte(ch) 
	goto st41
tr129:
//line lex.rl:213
 buf.WriteByte('\a') 
	goto st41
tr132:
//line lex.rl:214
 buf.WriteByte('\b') 
	goto st41
tr135:
//line lex.rl:215
 buf.WriteByte('\f') 
	goto st41
tr138:
//line lex.rl:216
 buf.WriteByte('\n') 
	goto st41
tr141:
//line lex.rl:217
 buf.WriteByte('\r') 
	goto st41
tr144:
//line lex.rl:218
 buf.WriteByte('\t') 
	goto st41
tr151:
//line lex.rl:219
 buf.WriteByte('\v') 
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line lex.go:17398
		switch data[p] {
		case 85:
			goto tr96
		case 88:
			goto st54
		case 97:
			goto st57
		case 98:
			goto st58
		case 102:
			goto st59
		case 110:
			goto st60
		case 114:
			goto st61
		case 116:
			goto st62
		case 117:
			goto tr104
		case 118:
			goto st67
		case 120:
			goto st54
		}
		if 48 <= data[p] && data[p] <= 55 {
			goto tr95
		}
		goto tr94
tr95:
//line lex.rl:240
 ch = 0 
//line lex.rl:241
 ch = (ch << 3) | data[p] - '0' 
	goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
//line lex.go:17438
		if 48 <= data[p] && data[p] <= 55 {
			goto tr106
		}
		goto tr44
tr106:
//line lex.rl:241
 ch = (ch << 3) | data[p] - '0' 
	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line lex.go:17452
		if 48 <= data[p] && data[p] <= 55 {
			goto tr107
		}
		goto tr44
tr107:
//line lex.rl:241
 ch = (ch << 3) | data[p] - '0' 
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line lex.go:17466
		switch data[p] {
		case 39:
			goto tr109
		case 92:
			goto tr110
		}
		goto tr108
tr96:
//line lex.rl:235
 rn = 0 
	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line lex.go:17483
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr111
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr111
			}
		default:
			goto tr111
		}
		goto tr44
tr111:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line lex.go:17506
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr112
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr112
			}
		default:
			goto tr112
		}
		goto tr44
tr112:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line lex.go:17529
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr113
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr113
			}
		default:
			goto tr113
		}
		goto tr44
tr113:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
//line lex.go:17552
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr114
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr114
			}
		default:
			goto tr114
		}
		goto tr44
tr114:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line lex.go:17575
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr115
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr115
			}
		default:
			goto tr115
		}
		goto tr44
tr115:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line lex.go:17598
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr116
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr116
			}
		default:
			goto tr116
		}
		goto tr44
tr116:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line lex.go:17621
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr117
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr117
			}
		default:
			goto tr117
		}
		goto tr44
tr117:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line lex.go:17644
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr118
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr118
			}
		default:
			goto tr118
		}
		goto tr44
tr118:
//line lex.rl:233
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st53
tr148:
//line lex.rl:230
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line lex.go:17671
		switch data[p] {
		case 39:
			goto tr120
		case 92:
			goto tr121
		}
		goto tr119
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr122
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr122
			}
		default:
			goto tr122
		}
		goto tr44
tr122:
//line lex.rl:223
 ch = 0 
//line lex.rl:224
 ch = (ch << 4) | unhex(data[p]) 
	goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line lex.go:17708
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr123
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr123
			}
		default:
			goto tr123
		}
		goto tr44
tr123:
//line lex.rl:224
 ch = (ch << 4) | unhex(data[p]) 
	goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
//line lex.go:17731
		switch data[p] {
		case 39:
			goto tr125
		case 92:
			goto tr126
		}
		goto tr124
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch data[p] {
		case 39:
			goto tr128
		case 92:
			goto tr129
		}
		goto tr127
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 39:
			goto tr131
		case 92:
			goto tr132
		}
		goto tr130
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 39:
			goto tr134
		case 92:
			goto tr135
		}
		goto tr133
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 39:
			goto tr137
		case 92:
			goto tr138
		}
		goto tr136
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 39:
			goto tr140
		case 92:
			goto tr141
		}
		goto tr139
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 39:
			goto tr143
		case 92:
			goto tr144
		}
		goto tr142
tr104:
//line lex.rl:235
 rn = 0 
	goto st63
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
//line lex.go:17820
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr145
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr145
			}
		default:
			goto tr145
		}
		goto tr44
tr145:
//line lex.rl:230
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st64
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
//line lex.go:17843
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr146
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr146
			}
		default:
			goto tr146
		}
		goto tr44
tr146:
//line lex.rl:230
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st65
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
//line lex.go:17866
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr147
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr147
			}
		default:
			goto tr147
		}
		goto tr44
tr147:
//line lex.rl:230
 rn = (rn << 4) | rune(unhex(data[p])) 
	goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line lex.go:17889
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr148
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr148
			}
		default:
			goto tr148
		}
		goto tr44
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch data[p] {
		case 39:
			goto tr150
		case 92:
			goto tr151
		}
		goto tr149
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 32:
			goto tr812
		case 33:
			goto tr813
		case 34:
			goto tr814
		case 35:
			goto tr815
		case 36:
			goto tr816
		case 38:
			goto tr818
		case 39:
			goto tr819
		case 43:
			goto tr820
		case 45:
			goto tr821
		case 46:
			goto tr822
		case 47:
			goto tr823
		case 48:
			goto tr824
		case 58:
			goto tr826
		case 60:
			goto tr827
		case 62:
			goto tr828
		case 63:
			goto tr829
		case 64:
			goto tr830
		case 66:
			goto tr832
		case 69:
			goto tr833
		case 88:
			goto tr834
		case 95:
			goto tr835
		case 98:
			goto tr837
		case 101:
			goto tr838
		case 120:
			goto tr839
		case 124:
			goto tr840
		case 126:
			goto tr841
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr817
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr812
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr831
					}
				case data[p] >= 59:
					goto tr817
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr817
					}
				case data[p] >= 97:
					goto tr836
				}
			default:
				goto tr817
			}
		default:
			goto tr825
		}
		goto tr842
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		switch data[p] {
		case 32:
			goto tr843
		case 33:
			goto tr844
		case 34:
			goto tr845
		case 35:
			goto tr846
		case 36:
			goto tr847
		case 38:
			goto tr849
		case 39:
			goto tr850
		case 43:
			goto tr851
		case 45:
			goto tr852
		case 46:
			goto tr853
		case 47:
			goto tr854
		case 48:
			goto tr855
		case 58:
			goto tr857
		case 60:
			goto tr858
		case 62:
			goto tr859
		case 63:
			goto tr860
		case 64:
			goto tr861
		case 66:
			goto tr863
		case 69:
			goto tr864
		case 88:
			goto tr865
		case 95:
			goto tr866
		case 98:
			goto tr868
		case 101:
			goto tr869
		case 120:
			goto tr870
		case 124:
			goto tr871
		case 126:
			goto tr872
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr848
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr843
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr862
					}
				case data[p] >= 59:
					goto tr848
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr848
					}
				case data[p] >= 97:
					goto tr867
				}
			default:
				goto tr848
			}
		default:
			goto tr856
		}
		goto tr873
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch data[p] {
		case 32:
			goto tr874
		case 33:
			goto tr875
		case 34:
			goto tr876
		case 35:
			goto tr877
		case 36:
			goto tr878
		case 38:
			goto tr880
		case 39:
			goto tr881
		case 43:
			goto tr882
		case 45:
			goto tr883
		case 46:
			goto tr884
		case 47:
			goto tr885
		case 48:
			goto tr886
		case 58:
			goto tr888
		case 60:
			goto tr889
		case 62:
			goto tr890
		case 63:
			goto tr891
		case 64:
			goto tr892
		case 66:
			goto tr894
		case 69:
			goto tr895
		case 88:
			goto tr896
		case 95:
			goto tr897
		case 98:
			goto tr899
		case 101:
			goto tr900
		case 120:
			goto tr901
		case 124:
			goto tr902
		case 126:
			goto tr903
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr879
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr874
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr893
					}
				case data[p] >= 59:
					goto tr879
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr879
					}
				case data[p] >= 97:
					goto tr898
				}
			default:
				goto tr879
			}
		default:
			goto tr887
		}
		goto tr904
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		switch data[p] {
		case 32:
			goto tr905
		case 33:
			goto tr906
		case 34:
			goto tr907
		case 35:
			goto tr908
		case 36:
			goto tr909
		case 38:
			goto tr911
		case 39:
			goto tr912
		case 43:
			goto tr913
		case 45:
			goto tr914
		case 46:
			goto tr915
		case 47:
			goto tr916
		case 48:
			goto tr917
		case 58:
			goto tr919
		case 59:
			goto tr910
		case 60:
			goto tr920
		case 61:
			goto st121
		case 62:
			goto tr922
		case 63:
			goto tr923
		case 64:
			goto tr924
		case 66:
			goto tr926
		case 69:
			goto tr927
		case 88:
			goto tr928
		case 95:
			goto tr929
		case 98:
			goto tr931
		case 101:
			goto tr932
		case 120:
			goto tr933
		case 124:
			goto tr934
		case 126:
			goto tr935
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr910
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr905
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr925
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr910
					}
				case data[p] >= 97:
					goto tr930
				}
			default:
				goto tr910
			}
		default:
			goto tr918
		}
		goto tr936
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		switch data[p] {
		case 32:
			goto tr937
		case 33:
			goto tr938
		case 34:
			goto tr939
		case 35:
			goto tr940
		case 36:
			goto tr941
		case 38:
			goto tr943
		case 39:
			goto tr944
		case 43:
			goto tr945
		case 45:
			goto tr946
		case 46:
			goto tr947
		case 47:
			goto tr948
		case 48:
			goto tr949
		case 58:
			goto tr951
		case 60:
			goto tr952
		case 62:
			goto tr953
		case 63:
			goto tr954
		case 64:
			goto tr955
		case 66:
			goto tr957
		case 69:
			goto tr958
		case 88:
			goto tr959
		case 95:
			goto tr960
		case 98:
			goto tr962
		case 101:
			goto tr963
		case 120:
			goto tr964
		case 124:
			goto tr965
		case 126:
			goto tr966
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr942
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr937
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr956
					}
				case data[p] >= 59:
					goto tr942
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr942
					}
				case data[p] >= 97:
					goto tr961
				}
			default:
				goto tr942
			}
		default:
			goto tr950
		}
		goto tr967
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 32:
			goto tr968
		case 33:
			goto tr969
		case 34:
			goto tr970
		case 35:
			goto tr971
		case 36:
			goto tr972
		case 38:
			goto tr974
		case 39:
			goto tr975
		case 43:
			goto tr976
		case 45:
			goto tr977
		case 46:
			goto tr978
		case 47:
			goto tr979
		case 48:
			goto tr980
		case 58:
			goto tr982
		case 60:
			goto tr983
		case 62:
			goto tr984
		case 63:
			goto tr985
		case 64:
			goto tr986
		case 66:
			goto tr988
		case 69:
			goto tr989
		case 88:
			goto tr990
		case 95:
			goto tr991
		case 98:
			goto tr993
		case 101:
			goto tr994
		case 120:
			goto tr995
		case 124:
			goto tr996
		case 126:
			goto tr997
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr973
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr968
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr987
					}
				case data[p] >= 59:
					goto tr973
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr973
					}
				case data[p] >= 97:
					goto tr992
				}
			default:
				goto tr973
			}
		default:
			goto tr981
		}
		goto tr998
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 32:
			goto tr999
		case 33:
			goto tr1000
		case 34:
			goto tr1001
		case 35:
			goto tr1002
		case 36:
			goto tr1003
		case 38:
			goto tr1005
		case 39:
			goto tr1006
		case 43:
			goto tr1007
		case 45:
			goto tr1008
		case 46:
			goto tr1009
		case 47:
			goto tr1010
		case 48:
			goto tr1011
		case 58:
			goto tr1013
		case 60:
			goto tr1014
		case 62:
			goto tr1015
		case 63:
			goto tr1016
		case 64:
			goto tr1017
		case 66:
			goto tr1019
		case 69:
			goto tr1020
		case 88:
			goto tr1021
		case 95:
			goto tr1022
		case 98:
			goto tr1024
		case 101:
			goto tr1025
		case 120:
			goto tr1026
		case 124:
			goto tr1027
		case 126:
			goto tr1028
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1004
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr999
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1018
					}
				case data[p] >= 59:
					goto tr1004
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1004
					}
				case data[p] >= 97:
					goto tr1023
				}
			default:
				goto tr1004
			}
		default:
			goto tr1012
		}
		goto tr1029
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch data[p] {
		case 32:
			goto tr1030
		case 33:
			goto tr1031
		case 34:
			goto tr1032
		case 35:
			goto tr1033
		case 36:
			goto tr1034
		case 38:
			goto tr1036
		case 39:
			goto tr1037
		case 43:
			goto tr1038
		case 45:
			goto tr1039
		case 46:
			goto tr1040
		case 47:
			goto tr1041
		case 48:
			goto tr1042
		case 58:
			goto tr1044
		case 60:
			goto tr1045
		case 62:
			goto tr1046
		case 63:
			goto tr1047
		case 64:
			goto tr1048
		case 66:
			goto tr1050
		case 69:
			goto tr1051
		case 88:
			goto tr1052
		case 95:
			goto tr1053
		case 98:
			goto tr1055
		case 101:
			goto tr1056
		case 120:
			goto tr1057
		case 124:
			goto tr1058
		case 126:
			goto tr1059
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1035
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr1030
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1049
					}
				case data[p] >= 59:
					goto tr1035
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1035
					}
				case data[p] >= 97:
					goto tr1054
				}
			default:
				goto tr1035
			}
		default:
			goto tr1043
		}
		goto tr1060
tr341:
//line lex.rl:91

			isFconst = true
		
	goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
//line lex.go:18781
		switch data[p] {
		case 43:
			goto st69
		case 45:
			goto st69
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st125
		}
		goto tr12
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		if 48 <= data[p] && data[p] <= 57 {
			goto st125
		}
		goto tr12
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 32:
			goto tr322
		case 33:
			goto tr323
		case 34:
			goto tr324
		case 35:
			goto tr325
		case 36:
			goto tr326
		case 38:
			goto tr328
		case 39:
			goto tr329
		case 43:
			goto tr330
		case 45:
			goto tr331
		case 46:
			goto tr332
		case 47:
			goto tr333
		case 58:
			goto tr334
		case 60:
			goto tr335
		case 62:
			goto tr336
		case 63:
			goto tr337
		case 64:
			goto tr338
		case 66:
			goto tr340
		case 69:
			goto tr1061
		case 88:
			goto tr342
		case 95:
			goto tr343
		case 98:
			goto tr345
		case 101:
			goto tr1062
		case 120:
			goto tr346
		case 124:
			goto tr347
		case 126:
			goto tr348
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr327
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr322
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr339
					}
				case data[p] >= 59:
					goto tr327
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr327
					}
				case data[p] >= 97:
					goto tr344
				}
			default:
				goto tr327
			}
		default:
			goto st125
		}
		goto tr349
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st126
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st126
			}
		default:
			goto st126
		}
		goto tr154
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		switch data[p] {
		case 32:
			goto tr1063
		case 33:
			goto tr1064
		case 34:
			goto tr1065
		case 35:
			goto tr1066
		case 36:
			goto tr1067
		case 38:
			goto tr1069
		case 39:
			goto tr1070
		case 43:
			goto tr1071
		case 45:
			goto tr1072
		case 46:
			goto tr1073
		case 47:
			goto tr1074
		case 58:
			goto tr1075
		case 59:
			goto tr1068
		case 60:
			goto tr1076
		case 61:
			goto tr1068
		case 62:
			goto tr1077
		case 63:
			goto tr1078
		case 64:
			goto tr1079
		case 88:
			goto tr1081
		case 95:
			goto tr1082
		case 120:
			goto tr1084
		case 124:
			goto tr1085
		case 126:
			goto tr1086
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1068
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr1063
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto tr1068
					}
				case data[p] >= 71:
					goto tr1080
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1068
					}
				case data[p] >= 103:
					goto tr1083
				}
			default:
				goto st126
			}
		default:
			goto st126
		}
		goto tr1087
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		switch data[p] {
		case 32:
			goto tr1088
		case 33:
			goto tr1089
		case 34:
			goto tr1090
		case 35:
			goto tr1091
		case 36:
			goto tr1092
		case 38:
			goto tr1094
		case 39:
			goto tr1095
		case 43:
			goto tr1096
		case 45:
			goto tr1097
		case 46:
			goto tr1098
		case 47:
			goto tr1099
		case 48:
			goto tr1100
		case 58:
			goto tr1102
		case 60:
			goto tr1103
		case 62:
			goto st128
		case 63:
			goto tr1105
		case 64:
			goto tr1106
		case 66:
			goto tr1108
		case 69:
			goto tr1109
		case 88:
			goto tr1110
		case 95:
			goto tr1111
		case 98:
			goto tr1113
		case 101:
			goto tr1114
		case 120:
			goto tr1115
		case 124:
			goto tr1116
		case 126:
			goto tr1117
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1093
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr1088
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1107
					}
				case data[p] >= 59:
					goto tr1093
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1093
					}
				case data[p] >= 97:
					goto tr1112
				}
			default:
				goto tr1093
			}
		default:
			goto tr1101
		}
		goto tr1118
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 32:
			goto tr1119
		case 33:
			goto tr1120
		case 34:
			goto tr1121
		case 35:
			goto tr1122
		case 36:
			goto tr1123
		case 38:
			goto tr1125
		case 39:
			goto tr1126
		case 43:
			goto tr1127
		case 45:
			goto tr1128
		case 46:
			goto tr1129
		case 47:
			goto tr1130
		case 48:
			goto tr1131
		case 58:
			goto tr1133
		case 60:
			goto tr1134
		case 62:
			goto tr1135
		case 63:
			goto tr1136
		case 64:
			goto tr1137
		case 66:
			goto tr1139
		case 69:
			goto tr1140
		case 88:
			goto tr1141
		case 95:
			goto tr1142
		case 98:
			goto tr1144
		case 101:
			goto tr1145
		case 120:
			goto tr1146
		case 124:
			goto tr1147
		case 126:
			goto tr1148
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1124
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr1119
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1138
					}
				case data[p] >= 59:
					goto tr1124
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1124
					}
				case data[p] >= 97:
					goto tr1143
				}
			default:
				goto tr1124
			}
		default:
			goto tr1132
		}
		goto tr1149
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 32:
			goto tr1150
		case 33:
			goto tr1151
		case 34:
			goto tr1152
		case 35:
			goto tr1153
		case 36:
			goto tr1154
		case 38:
			goto tr1156
		case 39:
			goto tr1157
		case 43:
			goto tr1158
		case 45:
			goto tr1159
		case 46:
			goto tr1160
		case 47:
			goto tr1161
		case 58:
			goto tr1162
		case 60:
			goto tr1163
		case 62:
			goto tr1164
		case 63:
			goto tr1165
		case 64:
			goto tr1166
		case 66:
			goto tr1168
		case 69:
			goto tr1169
		case 88:
			goto tr1170
		case 95:
			goto tr1171
		case 98:
			goto tr1173
		case 101:
			goto tr1174
		case 120:
			goto tr1175
		case 124:
			goto tr1176
		case 126:
			goto tr1177
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1155
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr1150
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1167
					}
				case data[p] >= 59:
					goto tr1155
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1155
					}
				case data[p] >= 97:
					goto tr1172
				}
			default:
				goto tr1155
			}
		default:
			goto st129
		}
		goto tr1178
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 32:
			goto tr1179
		case 33:
			goto tr1180
		case 34:
			goto tr1181
		case 35:
			goto tr1182
		case 36:
			goto tr1183
		case 38:
			goto tr1185
		case 39:
			goto tr1186
		case 43:
			goto tr1187
		case 45:
			goto tr1188
		case 46:
			goto tr1189
		case 47:
			goto tr1190
		case 48:
			goto tr1191
		case 58:
			goto tr1193
		case 60:
			goto tr1194
		case 62:
			goto tr1195
		case 63:
			goto tr1196
		case 64:
			goto tr1197
		case 66:
			goto tr1199
		case 69:
			goto tr1200
		case 88:
			goto tr1201
		case 95:
			goto tr1202
		case 98:
			goto tr1204
		case 101:
			goto tr1205
		case 120:
			goto tr1206
		case 124:
			goto tr1207
		case 126:
			goto tr1208
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1184
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr1179
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1198
					}
				case data[p] >= 59:
					goto tr1184
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1184
					}
				case data[p] >= 97:
					goto tr1203
				}
			default:
				goto tr1184
			}
		default:
			goto tr1192
		}
		goto tr1209
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 32:
			goto tr1210
		case 33:
			goto tr1211
		case 34:
			goto tr1212
		case 35:
			goto tr1213
		case 36:
			goto tr1214
		case 38:
			goto tr1216
		case 39:
			goto tr1217
		case 43:
			goto tr1218
		case 45:
			goto tr1219
		case 46:
			goto tr1220
		case 47:
			goto tr1221
		case 48:
			goto tr1222
		case 58:
			goto tr1224
		case 60:
			goto tr1225
		case 62:
			goto st132
		case 63:
			goto tr1227
		case 64:
			goto tr1228
		case 66:
			goto tr1230
		case 69:
			goto tr1231
		case 88:
			goto tr1232
		case 95:
			goto tr1233
		case 98:
			goto tr1235
		case 101:
			goto tr1236
		case 120:
			goto tr1237
		case 124:
			goto tr1238
		case 126:
			goto tr1239
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1215
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr1210
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1229
					}
				case data[p] >= 59:
					goto tr1215
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1215
					}
				case data[p] >= 97:
					goto tr1234
				}
			default:
				goto tr1215
			}
		default:
			goto tr1223
		}
		goto tr1240
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		switch data[p] {
		case 32:
			goto tr1241
		case 33:
			goto tr1242
		case 34:
			goto tr1243
		case 35:
			goto tr1244
		case 36:
			goto tr1245
		case 38:
			goto tr1247
		case 39:
			goto tr1248
		case 43:
			goto tr1249
		case 45:
			goto tr1250
		case 46:
			goto tr1251
		case 47:
			goto tr1252
		case 48:
			goto tr1253
		case 58:
			goto tr1255
		case 60:
			goto tr1256
		case 62:
			goto tr1257
		case 63:
			goto tr1258
		case 64:
			goto tr1259
		case 66:
			goto tr1261
		case 69:
			goto tr1262
		case 88:
			goto tr1263
		case 95:
			goto tr1264
		case 98:
			goto tr1266
		case 101:
			goto tr1267
		case 120:
			goto tr1268
		case 124:
			goto tr1269
		case 126:
			goto tr1270
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1246
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr1241
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1260
					}
				case data[p] >= 59:
					goto tr1246
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1246
					}
				case data[p] >= 97:
					goto tr1265
				}
			default:
				goto tr1246
			}
		default:
			goto tr1254
		}
		goto tr1271
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		switch data[p] {
		case 32:
			goto tr1272
		case 33:
			goto tr1273
		case 34:
			goto tr1274
		case 35:
			goto tr1275
		case 36:
			goto tr1276
		case 38:
			goto tr1278
		case 39:
			goto tr1279
		case 43:
			goto tr1280
		case 45:
			goto tr1281
		case 46:
			goto tr1282
		case 47:
			goto tr1283
		case 48:
			goto tr1284
		case 58:
			goto tr1286
		case 60:
			goto tr1287
		case 62:
			goto tr1288
		case 63:
			goto tr1289
		case 64:
			goto tr1290
		case 66:
			goto tr1292
		case 69:
			goto tr1293
		case 88:
			goto tr1294
		case 95:
			goto tr1295
		case 98:
			goto tr1297
		case 101:
			goto tr1298
		case 120:
			goto tr1299
		case 124:
			goto tr1300
		case 126:
			goto tr1301
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1277
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr1272
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1291
					}
				case data[p] >= 59:
					goto tr1277
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1277
					}
				case data[p] >= 97:
					goto tr1296
				}
			default:
				goto tr1277
			}
		default:
			goto tr1285
		}
		goto tr1302
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		switch data[p] {
		case 32:
			goto tr1303
		case 33:
			goto tr1304
		case 34:
			goto tr1305
		case 35:
			goto tr1306
		case 36:
			goto tr1307
		case 38:
			goto tr1309
		case 39:
			goto tr1310
		case 42:
			goto st135
		case 43:
			goto tr1312
		case 45:
			goto tr1313
		case 46:
			goto tr1314
		case 47:
			goto tr1315
		case 48:
			goto tr1316
		case 58:
			goto tr1318
		case 60:
			goto tr1319
		case 62:
			goto tr1320
		case 63:
			goto tr1321
		case 64:
			goto tr1322
		case 66:
			goto tr1324
		case 69:
			goto tr1325
		case 88:
			goto tr1326
		case 95:
			goto tr1327
		case 98:
			goto tr1329
		case 101:
			goto tr1330
		case 120:
			goto tr1331
		case 124:
			goto tr1332
		case 126:
			goto tr1333
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1308
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr1303
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1323
					}
				case data[p] >= 59:
					goto tr1308
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1308
					}
				case data[p] >= 97:
					goto tr1328
				}
			default:
				goto tr1308
			}
		default:
			goto tr1317
		}
		goto tr1334
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		switch data[p] {
		case 32:
			goto tr1335
		case 33:
			goto tr1336
		case 34:
			goto tr1337
		case 35:
			goto tr1338
		case 36:
			goto tr1339
		case 38:
			goto tr1341
		case 39:
			goto tr1342
		case 43:
			goto tr1343
		case 45:
			goto tr1344
		case 46:
			goto tr1345
		case 47:
			goto tr1346
		case 48:
			goto tr1347
		case 58:
			goto tr1349
		case 60:
			goto tr1350
		case 62:
			goto tr1351
		case 63:
			goto tr1352
		case 64:
			goto tr1353
		case 66:
			goto tr1355
		case 69:
			goto tr1356
		case 88:
			goto tr1357
		case 95:
			goto tr1358
		case 98:
			goto tr1360
		case 101:
			goto tr1361
		case 120:
			goto tr1362
		case 124:
			goto tr1363
		case 126:
			goto tr1364
		case 127:
			goto tr156
		}
		switch {
		case data[p] < 49:
			switch {
			case data[p] < 9:
				if data[p] <= 8 {
					goto tr156
				}
			case data[p] > 13:
				switch {
				case data[p] > 31:
					if 37 <= data[p] && data[p] <= 44 {
						goto tr1340
					}
				case data[p] >= 14:
					goto tr156
				}
			default:
				goto tr1335
			}
		case data[p] > 57:
			switch {
			case data[p] < 91:
				switch {
				case data[p] > 61:
					if 65 <= data[p] && data[p] <= 90 {
						goto tr1354
					}
				case data[p] >= 59:
					goto tr1340
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 125 {
						goto tr1340
					}
				case data[p] >= 97:
					goto tr1359
				}
			default:
				goto tr1340
			}
		default:
			goto tr1348
		}
		goto tr1365
	st_out:
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof100: cs = 100; goto _test_eof
	_test_eof101: cs = 101; goto _test_eof
	_test_eof102: cs = 102; goto _test_eof
	_test_eof103: cs = 103; goto _test_eof
	_test_eof104: cs = 104; goto _test_eof
	_test_eof105: cs = 105; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof121: cs = 121; goto _test_eof
	_test_eof122: cs = 122; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof126: cs = 126; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
	_test_eof131: cs = 131; goto _test_eof
	_test_eof132: cs = 132; goto _test_eof
	_test_eof133: cs = 133; goto _test_eof
	_test_eof134: cs = 134; goto _test_eof
	_test_eof135: cs = 135; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 1, 2, 3, 4, 7, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67:
//line lex.rl:366

				emitError("invalid syntax")
				return
			
		case 71, 84:
//line lex.rl:370
emit(0, "EOF"); return
		case 129:
//line lex.rl:59

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
		
//line lex.rl:370
emit(0, "EOF"); return
		case 85, 90, 91, 125:
//line lex.rl:73

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
		
//line lex.rl:370
emit(0, "EOF"); return
		case 126:
//line lex.rl:94

			emit(lex.ICONST, data[mark-1:p]); return
		
//line lex.rl:370
emit(0, "EOF"); return
		case 104, 106, 108, 110, 115:
//line lex.rl:97

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
		
//line lex.rl:370
emit(0, "EOF"); return
		case 74:
//line lex.rl:151

			if numQuote != 0 {
				b = make([]byte, p-mark-2-numQuote)
				// Now use numQuote as an index into b.
				numQuote = 0
				for i := mark+1; i < p-1; i++ {
					b[numQuote] = data[i]
					numQuote++
					if data[i] == '"' {
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
			emit(lex.IDENT, string(b)); return
		
//line lex.rl:370
emit(0, "EOF"); return
		case 81:
//line lex.rl:186

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
		
//line lex.rl:370
emit(0, "EOF"); return
		case 111, 116:
//line lex.rl:266

			emit(lex.BCONST, buf.String()); return
		
//line lex.rl:370
emit(0, "EOF"); return
		case 109:
//line lex.rl:270

			if !utf8.Valid(buf.Bytes()) {
				emitError(errInvalidUTF8)
				return
			}
			emit(lex.SCONST, buf.String()); return
		
//line lex.rl:370
emit(0, "EOF"); return
		case 107:
//line lex.rl:293

			emit(lex.BITCONST, string(data[mark+2:p-1])); return
		
//line lex.rl:370
emit(0, "EOF"); return
		case 70:
//line lex.rl:299

				emitError(errInvalidHexNumeric)
				return
			
//line lex.rl:366

				emitError("invalid syntax")
				return
			
		case 5, 68, 69:
//line lex.rl:303

				emitError("invalid floating point literal")
				return
			
//line lex.rl:366

				emitError("invalid syntax")
				return
			
		case 9, 39, 40:
//line lex.rl:313

				emitError("invalid hexadecimal bytes literal")
				return
			
//line lex.rl:366

				emitError("invalid syntax")
				return
			
		case 6:
//line lex.rl:317

				emitError(fmt.Sprintf(`"%c" is not a valid binary digit`, data[p]))
				return
			
//line lex.rl:366

				emitError("invalid syntax")
				return
			
		case 72, 73, 75, 76, 77, 78, 80, 82, 83, 86, 88, 92, 95, 98, 100, 102, 105, 113:
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 87:
//line lex.rl:324
 emitToken(lex.DOT_DOT); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 133:
//line lex.rl:326
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 135:
//line lex.rl:327
 emitToken(lex.NOT_REGIMATCH); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 134:
//line lex.rl:328
 emitToken(lex.NOT_REGMATCH); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 118:
//line lex.rl:330
 emitToken(lex.HELPTOKEN); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 119:
//line lex.rl:331
 emitToken(lex.JSON_SOME_EXISTS); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 101:
//line lex.rl:332
 emitToken(lex.JSON_ALL_EXISTS); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 97:
//line lex.rl:334
 emitToken(lex.INET_CONTAINED_BY_OR_EQUALS); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 96:
//line lex.rl:335
 emitToken(lex.LSHIFT); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 123:
//line lex.rl:336
 emitToken(lex.NOT_EQUALS); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 122:
//line lex.rl:337
 emitToken(lex.LESS_EQUALS); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 124:
//line lex.rl:338
 emitToken(lex.CONTAINED_BY); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 121:
//line lex.rl:340
 emitToken(lex.INET_CONTAINS_OR_EQUALS); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 120:
//line lex.rl:341
 emitToken(lex.RSHIFT); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 99:
//line lex.rl:342
 emitToken(lex.GREATER_EQUALS); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 94:
//line lex.rl:344
 emitToken(lex.TYPEANNOTATE); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 93:
//line lex.rl:345
 emitToken(lex.TYPECAST); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 117:
//line lex.rl:347
 emitToken(lex.CONCAT); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 89:
//line lex.rl:349
 emitToken(lex.FLOORDIV); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 114:
//line lex.rl:351
 emitToken(lex.REGIMATCH); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 103:
//line lex.rl:353
 emitToken(lex.CONTAINS); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 79:
//line lex.rl:355
 emitToken(lex.INET_CONTAINS_OR_CONTAINED_BY); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 128:
//line lex.rl:357
 emitToken(lex.FETCHTEXT); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 127:
//line lex.rl:358
 emitToken(lex.FETCHVAL); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 132:
//line lex.rl:360
 emitToken(lex.FETCHTEXT_PATH); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 131:
//line lex.rl:361
 emitToken(lex.FETCHVAL_PATH); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 130:
//line lex.rl:362
 emitToken(lex.REMOVE_PATH); return 
//line lex.rl:370
emit(0, "EOF"); return
		case 112:
//line lex.rl:97

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
		
//line lex.rl:322
 emitToken(int(data[p-1])); return 
//line lex.rl:370
emit(0, "EOF"); return
//line lex.go:20508
		}
	}

	_out: {}
	}

//line lex.rl:375

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
