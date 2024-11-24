package parser

import (
	"monkey/ast"
	"monkey/token"
	"testing"
)

func TestParseLetStatement(t *testing.T) {

	parser := New("let x = 2;")

	program := parser.parse()

	expectedLetStatement := ast.LetStatement{Value: &ast.NumberExpression{Token: token.Token{Type: token.INT, Literal: "2"}, Value: &ast.Indentifer{Value: "2"}}, Name: &ast.Indentifer{Value: "x"}, Token: token.Token{Type: token.LET, Literal: "let"}}

	actualLetStatement := program.Statements[0].(*ast.LetStatement)

	if !equalsTwoLetStatements(expectedLetStatement, *actualLetStatement) {
		t.Fatalf("tests failed")
	}

}

func equalsTwoLetStatements(expected ast.LetStatement, actual ast.LetStatement) bool {
	return expected.Name.Value == actual.Name.Value && expected.Value.(*ast.NumberExpression).Value.Value == actual.Value.(*ast.NumberExpression).Value.Value
}
