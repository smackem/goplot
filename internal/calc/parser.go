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

	if _, err = p.expect(ttLBracket); err != nil {
		return nil, err
	}

	lower, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	if _, err = p.expect(ttColon); err != nil {
		return nil, err
	}

	upper, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	if _, err = p.expect(ttRBracket); err != nil {
		return nil, err
	}

	if _, err = p.expect(ttArrow); err != nil {
		return nil, err
	}

	expr, err := p.parseExpr()
	if err != nil {
		return nil, err
	}

	p.assert(ttEOF)

	return &Function{
		Param: ident.Lexeme,
		lower: lower,
		upper: upper,
		body:  expr,
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
			left = addExpr{left: left, right: right}
		case ttMinus:
			p.next()
			right, err := p.parseProduct()
			if err != nil {
				return nil, err
			}
			left = minusExpr{left: left, right: right}
		default:
			return left, nil
		}
	}
}

func (p *parser) parseProduct() (Expr, error) {
	left, err := p.parseMolecule()
	if err != nil {
		return nil, err
	}
	for {
		switch p.current().Type {
		case ttStar:
			p.next()
			right, err := p.parseMolecule()
			if err != nil {
				return nil, err
			}
			left = timesExpr{left: left, right: right}
		case ttSlash:
			p.next()
			right, err := p.parseMolecule()
			if err != nil {
				return nil, err
			}
			left = divideExpr{left: left, right: right}
		case ttPercent:
			p.next()
			right, err := p.parseMolecule()
			if err != nil {
				return nil, err
			}
			left = moduloExpr{left: left, right: right}
		default:
			return left, nil
		}
	}
}

func (p *parser) parseMolecule() (Expr, error) {
	left, err := p.parseAtom()
	if err != nil {
		return nil, err
	}
	for {
		switch p.current().Type {
		case ttCaret:
			p.next()
			right, err := p.parseAtom()
			if err != nil {
				return nil, err
			}
			left = powerExpr{left: left, right: right}
		case ttLog:
			p.next()
			right, err := p.parseAtom()
			if err != nil {
				return nil, err
			}
			left = logExpr{left: left, right: right}
		default:
			return left, nil
		}
	}
}

func (p *parser) parseAtom() (Expr, error) {
	current := p.next()
	switch current.Type {
	case ttNumber:
		return numberExpr{number: current.parseNumber()}, nil
	case ttIdent:
		return identExpr{ident: current.Lexeme}, nil
	case ttMinus:
		inner, err := p.parseAtom()
		if err != nil {
			return nil, err
		}
		return negateExpr{inner: inner}, nil
	case ttSqrt:
		inner, err := p.parseAtom()
		if err != nil {
			return nil, err
		}
		return sqrtExpr{inner: inner}, nil
	case ttSin:
		inner, err := p.parseAtom()
		if err != nil {
			return nil, err
		}
		return sinExpr{inner: inner}, nil
	case ttCos:
		inner, err := p.parseAtom()
		if err != nil {
			return nil, err
		}
		return cosExpr{inner: inner}, nil
	case ttTan:
		inner, err := p.parseAtom()
		if err != nil {
			return nil, err
		}
		return tanExpr{inner: inner}, nil
	case ttAsin:
		inner, err := p.parseAtom()
		if err != nil {
			return nil, err
		}
		return asinExpr{inner: inner}, nil
	case ttAcos:
		inner, err := p.parseAtom()
		if err != nil {
			return nil, err
		}
		return acosExpr{inner: inner}, nil
	case ttAtan:
		inner, err := p.parseAtom()
		if err != nil {
			return nil, err
		}
		return atanExpr{inner: inner}, nil
	case ttLParen:
		inner, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		if _, err = p.expect(ttRParen); err != nil {
			return nil, err
		}
		return inner, nil
	default:
		return nil, fmt.Errorf("Expected atom, found %+v", current)
	}
}
