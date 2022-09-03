// Package token contains constants which are used by the lexer
package token

// Pre-defined TokenTypes
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"
	// Operators
	ASSIGN = "="
	PLUS   = "+"
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
)

type TokenType string

// Token struct represents the lexer token.
type Token struct {
	Type    TokenType
	Literal string
}

// NewToken returns a new Token struct value.
func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

// Reserved keywords.
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"var": VAR,
}

// LookupIdentifier used to determinate whether identifier is keyword nor not.
func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENT
}
