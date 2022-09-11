// Package token defines constants representing the lexical tokens of the Ruedalang
// programming language and basic operations on them.
package token

type TokenType string

// Pre-defined list of tokens.
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// Reserved keywords.
var keywords = map[string]TokenType{
	"fun":    FUNCTION,
	"var":    VAR,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// Token struct represents the lexer token.
type Token struct {
	Type    TokenType
	Literal string
}

// NewToken returns a new Token struct value.
func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

// IsKeyword reports whether name is keyword, such as "fn" or "var".
func IsKeyword(name string) bool {
	_, ok := keywords[name]
	return ok
}

// Lookup maps an identifier to its keyword token or IDENT.
func Lookup(ident string) TokenType {
	if tok, is_keyword := keywords[ident]; is_keyword {
		return tok
	}
	return IDENT
}
