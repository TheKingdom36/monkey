package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(Input string) Parser {
	p := Parser{l: lexer.New(Input)}

	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parse() *ast.Program {
	program := ast.Program{Statements: []ast.Statement{}}

	if p.curToken.Type == token.LET {
		letStatement := p.parseLet()
		program.Statements = append(program.Statements, letStatement)
	}

	return &program
}

func (p *Parser) parseLet() *ast.LetStatement {
	letStatement := ast.LetStatement{Token: p.curToken}

	if p.peekToken.Type != token.IDENT {
		return nil
	}

	p.NextToken()

	letStatement.Name = &ast.Indentifer{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekToken.Type != token.ASSIGN {
		return nil
	}

	p.NextToken()
	p.NextToken()
	letStatement.Value = p.parseExpression()

	return &letStatement
}

func (p *Parser) parseExpression() ast.Expression {

	if p.curToken.Type == token.INT {
		if p.peekToken.Type == token.PLUS {
			return 
		}else{
		return &ast.NumberExpression{Token: p.curToken, Value: &ast.Indentifer{Token: p.curToken, Value: p.curToken.Literal}}
		}
		} else {
		return nil
	}

	// keep processing until we hit a semi colon
	// process the left
	//process the right
}
