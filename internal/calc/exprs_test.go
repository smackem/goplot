package calc

import (
	"math"
	"testing"
)

func Test_Expr_Eval(t *testing.T) {
	tests := []struct {
		name string
		expr Expr
		ctx  EvalContext
		want Number
	}{
		{
			name: "identExpr",
			expr: identExpr{ident: "x"},
			ctx:  contextWith("x", 100),
			want: 100,
		},
		{
			name: "numberExpr",
			expr: numberExpr{number: 100},
			ctx:  contextWith("_", 0),
			want: 100,
		},
		{
			name: "addExpr",
			expr: addExpr{numberExpr{number: 100}, numberExpr{number: 150}},
			ctx:  contextWith("_", 0),
			want: 250,
		},
		{
			name: "minusExpr",
			expr: minusExpr{numberExpr{number: 100}, identExpr{ident: "x"}},
			ctx:  contextWith("x", 25),
			want: 75,
		},
		{
			name: "timesExpr",
			expr: timesExpr{
				addExpr{left: numberExpr{number: 25}, right: identExpr{ident: "y"}},
				divideExpr{left: numberExpr{number: 25}, right: numberExpr{number: 5}},
			},
			ctx:  contextWith("y", 25),
			want: 250,
		},
		{
			name: "powerExpr",
			expr: powerExpr{numberExpr{number: 2}, numberExpr{number: 10}},
			ctx:  contextWith("_", 0),
			want: 1024,
		},
		{
			name: "logExpr",
			expr: logExpr{numberExpr{number: 4}, numberExpr{number: 16}},
			ctx:  contextWith("_", 0),
			want: 2,
		},
		{
			name: "lognExpr",
			expr: logExpr{numberExpr{number: math.E}, numberExpr{number: Number(math.Pow(math.E, 10))}},
			ctx:  contextWith("_", 0),
			want: 10,
		},
		{
			name: "sqrtExpr",
			expr: sqrtExpr{inner: identExpr{ident: "x"}},
			ctx:  contextWith("x", 25),
			want: 5,
		},
		{
			name: "sinExpr",
			expr: sinExpr{inner: identExpr{ident: "x"}},
			ctx:  contextWith("x", math.Pi/2),
			want: 1,
		},
		{
			name: "cosExpr",
			expr: cosExpr{inner: identExpr{ident: "x"}},
			ctx:  contextWith("x", math.Pi),
			want: -1,
		},
		{
			name: "tanExpr",
			expr: tanExpr{inner: identExpr{ident: "x"}},
			ctx:  contextWith("x", 0),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.expr.Eval(tt.ctx); got != tt.want {
				t.Errorf("expr.Eval() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

// ============================================================================

func contextWith(ident string, value Number) EvalContext {
	ctx := map[string]Number{
		ident: value,
	}
	return context(ctx)
}
