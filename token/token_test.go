package token

import (
	"strings"
	"testing"
)

func TestNewToken(t *testing.T) {
	testTable := []struct {
		input           rune
		expectedType    TokenType
		expectedLiteral string
	}{
		{'=', TokenType("="), "="},
		{'+', TokenType("+"), "+"},
		{'{', TokenType("{"), "{"},
		{'}', TokenType("}"), "}"},
		{'(', TokenType("("), "("},
		{')', TokenType(")"), ")"},
		{';', TokenType(";"), ";"},
		{',', TokenType(","), ","},
	}

	for _, tt := range testTable {
		tok := NewToken(TokenType(tt.input), byte(tt.input))

		if tok.Type != tt.expectedType {
			t.Fatalf("invalid token. Want %q, got %q", tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("invalid literal. Want %q, got %q", tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestLookup(t *testing.T) {
	for key, val := range keywords {
		if Lookup(string(key)) != val {
			t.Errorf("Lookup of %s failed", key)
		}

		if Lookup(strings.ToUpper(string(key))) != IDENT {
			t.Errorf("Lookup of %s failed", key)
		}
	}
}
