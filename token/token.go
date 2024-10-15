package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEAGAL TokenType = "ILLEAGAL"
	EOF      TokenType = "EOF"

	IDENT TokenType = "IDENT" // add, foobar, x, y, ...
	INT   TokenType = "INT"   // 1343456
	// Operators
	ASSIGN TokenType = "="
	PLUS   TokenType = "+"
	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"
	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)
