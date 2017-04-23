package calc

import (
	"math"
)

// Parse parses the given string and returns the parsed Function or an error.
func Parse(src string) (*Function, error) {
	tokens, err := lex(src)
	if err != nil {
		return nil, err
	}

	return parse(tokens)
}

// The Number type represents a numerical value as used and produced by the calc package.
type Number float64

// Function represents a parsed function evaluation rule as produced by the calc package's Parse function.
type Function struct {
	// Param is the identifier of the function parameter.
	Param string
	lower Expr
	upper Expr
	body  Expr
}

// Eval evaluates the given Function and returns the resulting x and y values.
// The returned slices are parallel, both containing steps elements.
// ys[index] is the result of evaluating the parsed function with parameter value xs[index].
func (f Function) Eval(steps int) (xs []Number, ys []Number) {
	xs = make([]Number, steps)
	ys = make([]Number, steps)

	ctx := newContext()
	lower := f.lower.Eval(ctx)
	upper := f.upper.Eval(ctx)
	stepWidth := (upper - lower) / Number(steps)
	x := lower

	for i := 0; i < steps; i++ {
		ctx[f.Param] = x
		xs[i] = x
		ys[i] = f.body.Eval(ctx)
		x += stepWidth
	}

	return
}

// The Expr interface provides functionality exposed by all parsed expressions.
type Expr interface {
	// Eval evaluates the expression with the given EvalContext and returns the result.
	Eval(ctx EvalContext) Number
}

// The EvalContext interface is used by expression evaluation functions to access
// identifier/value bindings.
type EvalContext interface {
	// Get returns the value bound to the given identifier or the zero value
	// if no such identifier found.
	Get(ident string) Number
}

///////////////////////////////////////////////////////////////////////////////

type context map[string]Number

func newContext() context {
	return context{
		"pi": math.Pi,
		"e":  math.E,
	}
}

func (c context) Get(ident string) Number {
	return c[ident]
}
