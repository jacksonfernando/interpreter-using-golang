package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifier , Literal
	IDENT = "IDENT"
	INT   = "INT"

	// Operator
	ASSIGN = "="
	PLUS   = "+"

	// Delimiter
	COMMA     = ","
	SEMICOLON = ";"

	lPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
