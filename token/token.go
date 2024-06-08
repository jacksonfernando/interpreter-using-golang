package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	SLASH    = "/"
	ASTERISK = "*"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACKET = "{"
	RBRACKET = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"

	LT = "<"
	GT = ">"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keyword = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookUpIdent(tokenLiteral string) TokenType {
	if tok, ok := keyword[tokenLiteral]; ok {
		return tok
	}
	return IDENT
}
