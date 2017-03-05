package calc

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// tokenType represents the type of token
type tokenType int

// The types of tokens that can be lexed
const (
	ttLParen tokenType = iota
	ttRParen
	ttPlus
	ttMinus
	ttStar
	ttStarStar
	ttSlash
	ttPercent
	ttComma
	ttEq
	ttLog
	ttSqrt
	ttSin
	ttCos
	ttTan
	ttAsin
	ttAcos
	ttAtan
	ttNumber
	ttIdent
	ttLet
	ttEOF
)

// token represents a lexed token
type token struct {
	Type   tokenType
	Lexeme string
}

// lex walks the specified string and returns an array of lexed Tokens
// or an non-nil error if the input could not be lexed.
func lex(src string) ([]token, error) {
	tokens := []token{}
	isNotSpace := func(r rune) bool { return unicode.IsSpace(r) == false }

	for index := 0; index < len(src); {
		slice := src[index:]
		spaceCount := strings.IndexFunc(slice, isNotSpace)
		if spaceCount < 0 {
			break
		}

		index += spaceCount
		slice = src[index:]

		if token, lexemeLen := match(slice); lexemeLen >= 0 {
			tokens = append(tokens, token)
			index += lexemeLen
		} else {
			return tokens, fmt.Errorf("Error lexing '%s'", slice)
		}
	}

	return tokens, nil
}

func match(src string) (token, int) {
	for _, m := range matchers {
		if loc := m.regexp.FindStringIndex(src); loc != nil {
			tok := token{
				Type:   m.tokType,
				Lexeme: src[loc[0]:loc[1]],
			}
			return tok, loc[1]
		}
	}
	return token{Type: ttEOF, Lexeme: ""}, -1
}

type matcher struct {
	regexp  *regexp.Regexp
	tokType tokenType
}

func makeMatcher(pattern string, tokType tokenType) matcher {
	return matcher{
		regexp:  regexp.MustCompile(pattern),
		tokType: tokType,
	}
}

var matchers = []matcher{
	makeMatcher("^\\(", ttLParen),
	makeMatcher("^\\)", ttRParen),
	makeMatcher("^\\+", ttPlus),
	makeMatcher("^\\-", ttMinus),
	makeMatcher("^\\*\\*", ttStarStar),
	makeMatcher("^\\*", ttStar),
	makeMatcher("^/", ttSlash),
	makeMatcher("^%", ttPercent),
	makeMatcher("^,", ttComma),
	makeMatcher("^=", ttEq),
	makeMatcher("^log\\b", ttLog),
	makeMatcher("^sqrt\\b", ttSqrt),
	makeMatcher("^sin\\b", ttSin),
	makeMatcher("^cos\\b", ttCos),
	makeMatcher("^tan\\b", ttTan),
	makeMatcher("^asin\\b", ttAsin),
	makeMatcher("^acos\\b", ttAcos),
	makeMatcher("^atan\\b", ttAtan),
	makeMatcher("^let\\b", ttLet),
	makeMatcher("^[0-9]+(\\.[0-9]+)?\\b", ttNumber),
	makeMatcher("^[a-zA-Z_][a-zA-Z0-9_]*\\b", ttIdent),
}
