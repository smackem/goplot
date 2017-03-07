package calc

import (
	"reflect"
	"testing"
)

func Test_lex(t *testing.T) {
	tests := []struct {
		name    string
		src     string
		want    []token
		wantErr bool
	}{
		{
			name: "Basic Operators Without Whitespace",
			src:  "()+-*",
			want: []token{
				token{Type: ttLParen, Lexeme: "("},
				token{Type: ttRParen, Lexeme: ")"},
				token{Type: ttPlus, Lexeme: "+"},
				token{Type: ttMinus, Lexeme: "-"},
				token{Type: ttStar, Lexeme: "*"},
			},
			wantErr: false,
		},
		{
			name: "Basic Operators With Whitespace",
			src:  "  ( )+ -   *  ",
			want: []token{
				token{Type: ttLParen, Lexeme: "("},
				token{Type: ttRParen, Lexeme: ")"},
				token{Type: ttPlus, Lexeme: "+"},
				token{Type: ttMinus, Lexeme: "-"},
				token{Type: ttStar, Lexeme: "*"},
			},
			wantErr: false,
		},
		{
			name: "All Operators With Whitespace",
			src:  "( ) + - * ^ / % , .. -> [ ]",
			want: []token{
				token{Type: ttLParen, Lexeme: "("},
				token{Type: ttRParen, Lexeme: ")"},
				token{Type: ttPlus, Lexeme: "+"},
				token{Type: ttMinus, Lexeme: "-"},
				token{Type: ttStar, Lexeme: "*"},
				token{Type: ttCaret, Lexeme: "^"},
				token{Type: ttSlash, Lexeme: "/"},
				token{Type: ttPercent, Lexeme: "%"},
				token{Type: ttComma, Lexeme: ","},
				token{Type: ttDotDot, Lexeme: ".."},
				token{Type: ttArrow, Lexeme: "->"},
				token{Type: ttLBracket, Lexeme: "["},
				token{Type: ttRBracket, Lexeme: "]"},
			},
			wantErr: false,
		},
		{
			name: "Numbers",
			src:  "1 2 123 123.125",
			want: []token{
				token{Type: ttNumber, Lexeme: "1"},
				token{Type: ttNumber, Lexeme: "2"},
				token{Type: ttNumber, Lexeme: "123"},
				token{Type: ttNumber, Lexeme: "123.125"},
			},
			wantErr: false,
		},
		{
			name: "Keywords",
			src:  "log sqrt sin cos tan asin acos atan",
			want: []token{
				token{Type: ttLog, Lexeme: "log"},
				token{Type: ttSqrt, Lexeme: "sqrt"},
				token{Type: ttSin, Lexeme: "sin"},
				token{Type: ttCos, Lexeme: "cos"},
				token{Type: ttTan, Lexeme: "tan"},
				token{Type: ttAsin, Lexeme: "asin"},
				token{Type: ttAcos, Lexeme: "acos"},
				token{Type: ttAtan, Lexeme: "atan"},
			},
			wantErr: false,
		},
		{
			name: "Identifiers",
			src:  "letter logger sinner x _x Abc_def_123",
			want: []token{
				token{Type: ttIdent, Lexeme: "letter"},
				token{Type: ttIdent, Lexeme: "logger"},
				token{Type: ttIdent, Lexeme: "sinner"},
				token{Type: ttIdent, Lexeme: "x"},
				token{Type: ttIdent, Lexeme: "_x"},
				token{Type: ttIdent, Lexeme: "Abc_def_123"},
			},
			wantErr: false,
		},
		{
			name: "Function",
			src:  "x[0..100] -> sin x * cos((x) ^ PI/4.0)",
			want: []token{
				token{Type: ttIdent, Lexeme: "x"},
				token{Type: ttLBracket, Lexeme: "["},
				token{Type: ttNumber, Lexeme: "0"},
				token{Type: ttDotDot, Lexeme: ".."},
				token{Type: ttNumber, Lexeme: "100"},
				token{Type: ttRBracket, Lexeme: "]"},
				token{Type: ttArrow, Lexeme: "->"},
				token{Type: ttSin, Lexeme: "sin"},
				token{Type: ttIdent, Lexeme: "x"},
				token{Type: ttStar, Lexeme: "*"},
				token{Type: ttCos, Lexeme: "cos"},
				token{Type: ttLParen, Lexeme: "("},
				token{Type: ttLParen, Lexeme: "("},
				token{Type: ttIdent, Lexeme: "x"},
				token{Type: ttRParen, Lexeme: ")"},
				token{Type: ttCaret, Lexeme: "^"},
				token{Type: ttIdent, Lexeme: "PI"},
				token{Type: ttSlash, Lexeme: "/"},
				token{Type: ttNumber, Lexeme: "4.0"},
				token{Type: ttRParen, Lexeme: ")"},
			},
			wantErr: false,
		},
		{
			name:    "Empty",
			src:     "",
			want:    []token{},
			wantErr: false,
		},
		{
			name:    "Blank",
			src:     "       ",
			want:    []token{},
			wantErr: false,
		},
		{
			name:    "Unknown Symbol",
			src:     "$",
			want:    []token{},
			wantErr: true,
		},
		{
			name: "Expression with Unknown Symbol",
			src:  "1000 * 1000$",
			want: []token{
				token{Type: ttNumber, Lexeme: "1000"},
				token{Type: ttStar, Lexeme: "*"},
				token{Type: ttNumber, Lexeme: "1000"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lex(tt.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("lex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_match(t *testing.T) {
	tests := []struct {
		name  string
		src   string
		want  token
		want1 int
	}{
		{
			name:  "LParen",
			src:   "(",
			want:  token{Type: ttLParen, Lexeme: "("},
			want1: 1,
		},
		{
			name:  "Integer Number",
			src:   "123",
			want:  token{Type: ttNumber, Lexeme: "123"},
			want1: 3,
		},
		{
			name:  "Float Number",
			src:   "123.25",
			want:  token{Type: ttNumber, Lexeme: "123.25"},
			want1: 6,
		},
		{
			name:  "Unknown Symbol",
			src:   "§",
			want:  token{Type: ttEOF, Lexeme: ""},
			want1: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := match(tt.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("match() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("match() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
