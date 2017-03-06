package calc

type number float64

type function struct {
	param string
	lower number
	upper number
	expr  *expr
}

type expr interface{}

type binaryExpr struct {
	left  *expr
	right *expr
}

type unaryExpr struct {
	expr *expr
}

type addExpr binaryExpr
type minusExpr binaryExpr
type timesExpr binaryExpr
type divideExpr binaryExpr
type moduloExpr binaryExpr
type powerExpr binaryExpr

type negateExpr unaryExpr
type sqrtExpr unaryExpr
type sinExpr unaryExpr
type cosExpr unaryExpr
type tanExpr unaryExpr
type asinExpr unaryExpr
type acosExpr unaryExpr
type atanExpr unaryExpr
