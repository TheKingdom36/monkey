package ast

import (
	"monkey/token"
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
	} else {
		return ""
	}
}

type LetStatement struct {
	Name  *Indentifer
	Value Expression
	Token token.Token
}

func (letStatement *LetStatement) TokenLiteral() string {
	return letStatement.Token.Literal
}

func (ls *LetStatement) statementNode() {}

type Indentifer struct {
	Token token.Token
	Value string
}

func (indentifer *Indentifer) TokenLiteral() string {
	return indentifer.Token.Literal
}

func (indentifer *Indentifer) expressionNode() {
}

type NumberExpression struct {
	Value *Indentifer
	Token token.Token
}

func (numberExpression *NumberExpression) TokenLiteral() string {
	return numberExpression.Token.Literal
}

func (numberExpression *NumberExpression) expressionNode() {
}

type OperatorExpression struct {
	Left     *Expression
	Operator string
	Right    *Expression
	Token    token.Token
}

func (operatorExpression *OperatorExpression) TokenLiteral() string {
	return operatorExpression.Token.Literal
}

func (operatorExpression *OperatorExpression) expressionNode() {
}

type ReturnStatement struct {
	Expression *Expression
	Token      token.Token
}

func (returnStatement *ReturnStatement) TokenLiteral() string {
	return returnStatement.Token.Literal
}

func (returnStatement *ReturnStatement) statementNode() {
}
