package ast

import "NeaGogu/monkey-interpreter/token"

// let y = 15;

// let y = 15;
// let <identifier> = <expression>

// every node in the AST has to implement this interface
type Node interface {
	// will be used mainly for debugging and testing
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

// the Program node is the root of the AST
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

// let a = 123
type LetStatement struct {
	Value Expression  // 123
	Name  *Identifier // 'a'
	Token token.Token // the token.LET token
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string      // the name of the ident: e.g let a = 123 value is "a"
}

// identifier is an expression (it does not produce  a value) for convenience
// It’s to keep things simple. Identifiers in other parts
// of a Monkey program do produce values, e.g.: let x = valueProducingIdentifier;. And to
// keep the number of different node types small, we’ll use Identifier here to represent the name
// in a variable binding and later reuse it, to represent an identifier as part of or as a complete
// expression.
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// --- return statements

type ReturnStatement struct {
	ReturnValue Expression
	Token       token.Token // the token.Return token
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
