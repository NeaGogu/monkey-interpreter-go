package parser

import (
	"NeaGogu/monkey-interpreter/ast"
	"NeaGogu/monkey-interpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x  =5;
	let y = 7;
	let foobar = 6969;
`

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()
	checkParseErrors(t, p)

	if prog == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	if len(prog.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. Got=%d", len(prog.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := prog.Statements[i]
		t.Log("cur statement", stmt)

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

	}
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 10;
	return foo(10);
	return "skibidi";
`

	l := lexer.New(input)
	p := New(l)

	prog := p.ParseProgram()
	checkParseErrors(t, p)

	if prog == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	if len(prog.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. Got=%d", len(prog.Statements))
	}

	for _, st := range prog.Statements {
		rst, ok := st.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("Statement is not ast.ReturnStatement; got=%q", st.TokenLiteral())
		}

		t.Log("cur statement", rst)

		if rst.TokenLiteral() != "return" {
			t.Errorf("Return statement literal is not 'return'; got=%q", rst.TokenLiteral())
		}
	}
}

func checkParseErrors(t *testing.T, p *Parser) {
	t.Helper()
	errors := p.errors

	if len(errors) == 0 {
		return
	}

	t.Logf("got %d parsing errors", len(errors))

	for _, e := range errors {
		t.Logf("parse error: %q\n", e)
	}

	t.FailNow()
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got = %q", stmt.TokenLiteral())
		return false
	}

	// check if its an ast.LetStatement (type cast)
	//
	// check if letStmt.Token.Type is let
	// check if letStmt.Name.Value is expectedId
	letStmt, ok := stmt.(*ast.LetStatement)

	if !ok {
		t.Errorf("stmt is not ast.LetStatement. got=%T", letStmt)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt identifier name is not %q. got=%q", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not %q. got=%s", name, letStmt.Name)
		return false
	}

	return true
}
