package calc

import "fmt"

func parse(input []token) (*Function, error) {
	parser := parser{input: input, index: 0}
	return parser.parseFunction()
}

type parser struct {
	input []token
	index int
}

func (p parser) current() token {
	if p.index >= len(p.input) {
		return token{Type: ttEOF, Lexeme: ""}
	}
	return p.input[p.index]
}

func (p *parser) next() token {
	tok := p.current()
	p.index++
	return tok
}

func (p *parser) assert(tt tokenType) (token, error) {
	tok := p.current()
	if tok.Type != tt {
		return emptyToken, fmt.Errorf("Expected %v, found %v", tt, tok.Type)
	}
	return tok, nil
}

func (p *parser) expect(tt tokenType) (token, error) {
	tok := p.next()
	if tok.Type != tt {
		return emptyToken, fmt.Errorf("Expected %v, found %v", tt, tok.Type)
	}
	return tok, nil
}

func (p *parser) parseFunction() (*Function, error) {
	ident, err := p.expect(ttIdent)
	if err != nil {
		return nil, err
	}

	_, err = p.expect(ttLBracket)
	if err != nil {
		return nil, err
	}

	lower, err := p.expect(ttNumber)
	if err != nil {
		return nil, err
	}

	_, err = p.expect(ttDotDot)
	if err != nil {
		return nil, err
	}

	upper, err := p.expect(ttNumber)
	if err != nil {
		return nil, err
	}

	_, err = p.expect(ttRBracket)
	if err != nil {
		return nil, err
	}

	_, err = p.expect(ttArrow)
	if err != nil {
		return nil, err
	}

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	return &Function{
		param: ident.Lexeme,
		lower: lower.parseNumber(),
		upper: upper.parseNumber(),
		expr:  expr,
	}, nil
}

func (p *parser) parseExpr() (Expr, error) {
	left, err := p.parseProduct()
	if err != nil {
		return nil, err
	}
	for {
		switch p.current().Type {
		case ttPlus:
			p.next()
			right, err := p.parseProduct()
			if err != nil {
				return nil, err
			}
			left = &addExpr{left: left, right: right}
		case ttMinus:
			p.next()
			right, err := p.parseProduct()
			if err != nil {
				return nil, err
			}
			left = &minusExpr{left: left, right: right}
		default:
			return left, nil
		}
	}
}

func (p *parser) parseProduct() (Expr, error) {
	return nil, nil
}
