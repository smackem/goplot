package calc

import "testing"

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
