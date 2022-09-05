// Package ast declares the types used to represent syntax trees.
package ast

import (
	"github.com/ferueda/ruedalang/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// ----------------------------------------------------------------------------
// Statements

// A statement is represented by a tree consisting of one
// or more of the following concrete statement nodes.

// An VarStatement node represents a variable bnding (var).
type VarStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (as *VarStatement) statementNode()       {}
func (as *VarStatement) TokenLiteral() string { return as.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
