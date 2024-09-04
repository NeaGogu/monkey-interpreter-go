package parser

import (
	"NeaGogu/monkey-interpreter/ast"
	"NeaGogu/monkey-interpreter/lexer"
	"NeaGogu/monkey-interpreter/token"
	"fmt"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	prog := &ast.Program{}
	prog.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()

		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}

		// NOTE: parseProgram advances the token after getting the statement, so DO NOT advance the token after finding
		// the semicolon
		p.nextToken()
	}

	return prog
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// lalalala test
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	st := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: skip expression for now, only look for semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return st
}

// lalal test 21
func (p *Parser) parseLetStatement() *ast.LetStatement {
	st := &ast.LetStatement{Token: p.curToken}

	// check if first token is identifier
	// this is where we use the p.peekToken field
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	st.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// check if next token is token.EQ
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// check if next token is expression (skip for now)
	// TODO: skip expressions until we know how to parse;
	// we'll skip until we find a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return st
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// assertion function; if the peeked token is the expected token,
// then the tokens are advanced and true is returned. Otherwise return false
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

func (p *Parser) peekError(expected token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s; got %s", expected, p.peekToken.Type)

	p.errors = append(p.errors, msg)
}
