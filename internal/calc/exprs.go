package calc

import "math"

type binaryExpr struct {
	left  Expr
	right Expr
}

type addExpr binaryExpr

func (x addExpr) Eval(ctx EvalContext) Number {
	return x.left.Eval(ctx) + x.right.Eval(ctx)
}

type minusExpr binaryExpr

func (x minusExpr) Eval(ctx EvalContext) Number {
	return x.left.Eval(ctx) - x.right.Eval(ctx)
}

type timesExpr binaryExpr

func (x timesExpr) Eval(ctx EvalContext) Number {
	return x.left.Eval(ctx) * x.right.Eval(ctx)
}

type divideExpr binaryExpr

func (x divideExpr) Eval(ctx EvalContext) Number {
	return x.left.Eval(ctx) / x.right.Eval(ctx)
}

type moduloExpr binaryExpr

func (x moduloExpr) Eval(ctx EvalContext) Number {
	return Number(int(x.left.Eval(ctx)) % int(x.right.Eval(ctx)))
}

type powerExpr binaryExpr

func (x powerExpr) Eval(ctx EvalContext) Number {
	left := float64(x.left.Eval(ctx))
	right := float64(x.right.Eval(ctx))
	return Number(math.Pow(left, right))
}

type logExpr binaryExpr

func (x logExpr) Eval(ctx EvalContext) Number {
	panic("unimplemented")
}

type unaryExpr struct {
	inner Expr
}

type negateExpr unaryExpr

func (x negateExpr) Eval(ctx EvalContext) Number {
	inner := float64(x.inner.Eval(ctx))
	return Number(-inner)
}

type sqrtExpr unaryExpr

func (x sqrtExpr) Eval(ctx EvalContext) Number {
	inner := float64(x.inner.Eval(ctx))
	return Number(math.Sqrt(inner))
}

type sinExpr unaryExpr

func (x sinExpr) Eval(ctx EvalContext) Number {
	inner := float64(x.inner.Eval(ctx))
	return Number(math.Sin(inner))
}

type cosExpr unaryExpr

func (x cosExpr) Eval(ctx EvalContext) Number {
	inner := float64(x.inner.Eval(ctx))
	return Number(math.Cos(inner))
}

type tanExpr unaryExpr

func (x tanExpr) Eval(ctx EvalContext) Number {
	inner := float64(x.inner.Eval(ctx))
	return Number(math.Tan(inner))
}

type asinExpr unaryExpr

func (x asinExpr) Eval(ctx EvalContext) Number {
	inner := float64(x.inner.Eval(ctx))
	return Number(math.Asin(inner))
}

type acosExpr unaryExpr

func (x acosExpr) Eval(ctx EvalContext) Number {
	inner := float64(x.inner.Eval(ctx))
	return Number(math.Acos(inner))
}

type atanExpr unaryExpr

func (x atanExpr) Eval(ctx EvalContext) Number {
	inner := float64(x.inner.Eval(ctx))
	return Number(math.Atan(inner))
}

type identExpr struct {
	ident string
}

func (x identExpr) Eval(ctx EvalContext) Number {
	return ctx.Get(x.ident)
}

type numberExpr struct {
	number Number
}

func (x numberExpr) Eval(ctx EvalContext) Number {
	return x.number
}
