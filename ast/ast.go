package ast

import "github.com/v2/golang-intrepeter/token"

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

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l *LetStatement) statementNode() {}

type Identifier struct {
	Token token.Token
	Value string
}

func (l *Identifier) TokenLiteral() string {
	return l.Token.Literal
}

func (l *Identifier) expressionNode() {}
