package tokens

import "fmt"

type TokenType = string

const (
	W TokenType = "W"
	I TokenType = "I"
	O TokenType = "O"
	R TokenType = "R"
	N TokenType = "N"
	C TokenType = "C"
)

type Token struct {
	Type   TokenType
	Code   int
	Lexeme string
	Line   int
	Col    int
}

func (t Token) String() string {
	return fmt.Sprintf("%s%d", t.Type, t.Code)
}