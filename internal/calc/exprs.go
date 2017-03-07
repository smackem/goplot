package calc

import "math"

type Number float64

type Function struct {
	param string
	lower Number
	upper Number
	expr  Expr
}

type Expr interface {
	Eval() Number
}

type binaryExpr struct {
	left  Expr
	right Expr
}

type unaryExpr struct {
	inner Expr
}

type addExpr binaryExpr

func (x addExpr) Eval() Number {
	return x.left.Eval() + x.right.Eval()
}

type minusExpr binaryExpr

func (x minusExpr) Eval() Number {
	return x.left.Eval() - x.right.Eval()
}

type timesExpr binaryExpr

func (x timesExpr) Eval() Number {
	return x.left.Eval() * x.right.Eval()
}

type divideExpr binaryExpr

func (x divideExpr) Eval() Number {
	return x.left.Eval() / x.right.Eval()
}

type moduloExpr binaryExpr

func (x moduloExpr) Eval() Number {
	return Number(int(x.left.Eval()) % int(x.right.Eval()))
}

type powerExpr binaryExpr

func (x powerExpr) Eval() Number {
	left := float64(x.left.Eval())
	right := float64(x.right.Eval())
	return Number(math.Pow(left, right))
}

type negateExpr unaryExpr
type sqrtExpr unaryExpr
type sinExpr unaryExpr
type cosExpr unaryExpr
type tanExpr unaryExpr
type asinExpr unaryExpr
type acosExpr unaryExpr
type atanExpr unaryExpr
