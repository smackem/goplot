package calc

import (
	"fmt"
	"reflect"
	"testing"
)

func makeFunction(param string, lower Number, upper Number, exprTokens []token) []token {
	tokens := []token{
		// x[0..0] ->
		token{Type: ttIdent, Lexeme: param},
		token{Type: ttLBracket, Lexeme: "["},
		token{Type: ttNumber, Lexeme: fmt.Sprintf("%f", lower)},
		token{Type: ttDotDot, Lexeme: ".."},
		token{Type: ttNumber, Lexeme: fmt.Sprintf("%f", upper)},
		token{Type: ttRBracket, Lexeme: "]"},
		token{Type: ttArrow, Lexeme: "->"},
	}
	return append(tokens, exprTokens...)
}

func Test_parse(t *testing.T) {
	tests := []struct {
		name    string
		input   []token
		want    *Function
		wantErr bool
	}{
		{
			name: "x[0..0] -> x",
			input: makeFunction("x", 0, 0, []token{
				token{Type: ttIdent, Lexeme: "x"},
			}),
			want: &Function{
				Param: "x",
				lower: numberExpr{number: 0.0},
				upper: numberExpr{number: 0.0},
				body:  identExpr{ident: "x"},
			},
		},
		{
			name: "x[-100..5*2] -> x",
			input: []token{
				token{Type: ttIdent, Lexeme: "x"},
				token{Type: ttLBracket, Lexeme: "["},
				token{Type: ttMinus, Lexeme: "-"},
				token{Type: ttNumber, Lexeme: "100"},
				token{Type: ttDotDot, Lexeme: ".."},
				token{Type: ttNumber, Lexeme: "5"},
				token{Type: ttStar, Lexeme: "*"},
				token{Type: ttNumber, Lexeme: "2"},
				token{Type: ttRBracket, Lexeme: "]"},
				token{Type: ttArrow, Lexeme: "->"},
				token{Type: ttIdent, Lexeme: "x"},
			},
			want: &Function{
				Param: "x",
				lower: negateExpr{inner: numberExpr{number: 100}},
				upper: timesExpr{left: numberExpr{number: 5}, right: numberExpr{number: 2}},
				body:  identExpr{ident: "x"},
			},
		},
		{
			name: "y[1..100] -> y + 1",
			input: makeFunction("y", 1, 100, []token{
				token{Type: ttIdent, Lexeme: "y"},
				token{Type: ttPlus, Lexeme: "+"},
				token{Type: ttNumber, Lexeme: "1"},
			}),
			want: &Function{
				Param: "y",
				lower: numberExpr{number: 1},
				upper: numberExpr{number: 100},
				body:  addExpr{left: identExpr{ident: "y"}, right: numberExpr{number: 1.0}},
			},
		},
		{
			name: "x[1..100] -> x * (x - 1) + x^2 / sin -x",
			input: makeFunction("x", 1, 100, []token{
				token{Type: ttIdent, Lexeme: "x"},
				token{Type: ttStar, Lexeme: "*"},
				token{Type: ttLParen, Lexeme: "("},
				token{Type: ttIdent, Lexeme: "x"},
				token{Type: ttMinus, Lexeme: "-"},
				token{Type: ttNumber, Lexeme: "1"},
				token{Type: ttRParen, Lexeme: ")"},
				token{Type: ttPlus, Lexeme: "+"},
				token{Type: ttIdent, Lexeme: "x"},
				token{Type: ttCaret, Lexeme: "^"},
				token{Type: ttNumber, Lexeme: "2"},
				token{Type: ttSlash, Lexeme: "/"},
				token{Type: ttSin, Lexeme: "sin"},
				token{Type: ttMinus, Lexeme: "-"},
				token{Type: ttIdent, Lexeme: "x"},
			}),
			want: &Function{
				Param: "x",
				lower: numberExpr{number: 1},
				upper: numberExpr{number: 100},
				body: addExpr{
					left: timesExpr{
						left:  identExpr{ident: "x"},
						right: minusExpr{left: identExpr{ident: "x"}, right: numberExpr{number: 1}},
					},
					right: divideExpr{
						left:  powerExpr{left: identExpr{ident: "x"}, right: numberExpr{number: 2}},
						right: sinExpr{inner: negateExpr{inner: identExpr{ident: "x"}}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
