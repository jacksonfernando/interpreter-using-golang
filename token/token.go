package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifier , literal
	IDENT = "IDENT"
	INT   = "INT"

	// operator
	ASSIGN = "="
	PLUS   = "+"

	// Delimiter
	COMMA     = ","
	SEMICOLON = ";"

	lPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
)
