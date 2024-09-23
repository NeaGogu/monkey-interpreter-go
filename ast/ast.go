package ast

import (
	"NeaGogu/monkey-interpreter/token"
	"bytes"
)

// let y = 15;

// let y = 15;
// let <identifier> = <expression>

// every node in the AST has to implement this interface
type Node interface {
	// will be used mainly for debugging and testing
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// ----------- let statements

// let a = 123
type LetStatement struct {
	Value Expression  // 123
	Name  *Identifier // 'a'
	Token token.Token // the token.LET token
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral())
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	// will be removed after we know how to parse expressions
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

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

func (i *Identifier) String() string {
	return i.Value
}

// --- return statements

type ReturnStatement struct {
	ReturnValue Expression
	Token       token.Token // the token.Return token
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	// will be removed after we know how to parse expressions
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

// -------- expressions

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

// the main reason we have an "expressioon statement" is so we can add it to the
// statement slice of the Program
func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
