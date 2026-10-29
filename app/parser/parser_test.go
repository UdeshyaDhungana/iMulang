package parser

import (
	"fmt"
	"testing"

	"github.com/udeshyadhungana/interprerer/app/ast"
	"github.com/udeshyadhungana/interprerer/app/lexer"
	"github.com/udeshyadhungana/interprerer/app/token"
)

func checkParserErrors(t *testing.T, p *Parser) {
	/* Check lexer errors */
	if len(p.l.Errors()) != 0 {
		t.Errorf("encountered %d errors from lexer\n", len(p.l.Errors()))
		p.l.ReportErrors()
		t.FailNow()
	}
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("encountered %d errors while parsing\n", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q\n", msg)
	}
	t.FailNow()
}

/* THOOS */

func TestThoosStatements(t *testing.T) {
	input := `
	thoos x = 4;
	thoos y = x;
	thoos foobar = 234543;
	`

	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testThoosStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testThoosStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "thoos" {
		t.Errorf("s.TokenLiteral not 'thoos'. got=%q", s.TokenLiteral())
		return false
	}

	thoosStmt, ok := s.(*ast.ThoosStatement)
	if !ok {
		t.Errorf("s not *ast.ThoosStatement. got=%T", s)
		return false
	}

	if thoosStmt.Name.Value != name {
		t.Errorf("thoosStmt.Name.Value not '%s'. got=%s", name, thoosStmt.Name.Value)
		return false
	}

	if thoosStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, thoosStmt.Name)
		return false
	}

	return true
}

/* PATHA */
func TestPathaStatements(t *testing.T) {
	input := `
	patha 2;
	patha a;
	patha x + y;
	patha "mero string";`

	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 4 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		ptmjStmt, ok := stmt.(*ast.PathaStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnState")
			continue
		}
		if ptmjStmt.TokenLiteral() != "patha" {
			t.Errorf("ptmjStmt.TokenLiteral not 'patha', got %q", program.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral is not %s. got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "555"

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != 555 {
		t.Errorf("literal.Value is not %d. got=%d", 555, literal.Value)
	}
	if literal.TokenLiteral() != "555" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "555", literal.TokenLiteral())
	}
}

func TestFloatLiteralExpression(t *testing.T) {
	input := "69.69"

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.FloatLiteral)
	if !ok {
		t.Fatalf("exp not *ast.FloatLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != 69.69 {
		t.Errorf("literal.Value is not %f. got=%f", 69.69, literal.Value)
	}
	if literal.TokenLiteral() != "69.69" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "69.69", literal.TokenLiteral())
	}
}

func TestBooleanLiteralExpression(t *testing.T) {
	input := "jhut;"

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.Boolean)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != false {
		t.Errorf("literal.Value is not %t. got=%t", false, literal.Value)
	}
	if literal.TokenLiteral() != "jhut" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "555", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"!5;", "!", int64(5)},
		{"-15;", "-", int64(15)},
		{"!jhut;", "!", false},
		{"!sacho;", "!", true},
	}

	for _, tt := range prefixTests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. Got=%d\n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not an ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not %s. got=%s", tt.operator, exp.Operator)
		}

		if !testLiteral(t, exp.Right, tt.value) {
			return
		}
	}
}

func testLiteral(t *testing.T, il ast.Expression, value any) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if ok {
		v, ok := value.(int64)
		if !ok {
			t.Errorf("value is not int64. got=%T", v)
			return false
		}
		if integ.Value != v {
			t.Errorf("integ.Value not (%d). got=(%d)", v, integ.Value)
			return false
		}

		if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
			t.Errorf("integ.TokenLiteral not %s. got=%s", value, integ.TokenLiteral())
			return false
		}
		return true
	}
	b, ok := il.(*ast.Boolean)
	if ok {
		v, ok := value.(bool)
		if !ok {
			t.Errorf("value is not bool. got=%T", b)
			return false
		}
		if b.Value != v {
			t.Errorf("boolean.Value not %t, got=%t", v, b.Value)
			return false
		}

		if b.Token.Type != token.JHUT && b.Token.Type != token.SACHO {
			t.Errorf("boolean.TokenLiteral not %s or %s. got=%s", token.SACHO, token.JHUT, b.TokenLiteral())
			return false
		}
		return true
	}

	id, ok := il.(*ast.Identifier)
	if ok {
		v, ok := value.(string)
		if !ok {
			t.Errorf("value is not string. got=%T", b)
			return false
		}
		if id.Value != v {
			t.Errorf("id.Value not %s, got=%s", v, id.Value)
		}

		if id.Token.Type != token.IDFIER {
			t.Errorf("identifier.TokenLiteral not %s, got=%s\n", token.IDFIER, id.Token.Type)
			return false
		}
		return true
	}

	t.Errorf("il is not *ast.IntegerLiteral or *ast.Boolean. got=%T", il)
	return false
}

func TestParsingInixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d elements. got=%d\n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is not ast.InfixExpression. got=%T", stmt.Expression)
		}

		if !testLiteral(t, exp.Left, tt.leftValue) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}

		if !testLiteral(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1 + (2 + 3) + 4;", "((1 + (2 + 3)) + 4);"},
		{"-a * b;",
			"((-a) * b);"},
		{"!-a;",
			"(!(-a));"},
		{
			"a + b + c;",
			"((a + b) + c);",
		},
		{
			"a + b - c;",
			"((a + b) - c);",
		},
		{
			"a * b * c;",
			"((a * b) * c);",
		},
		{
			"a * b / c;",
			"((a * b) / c);",
		},
		{
			"a + b / c;",
			"(a + (b / c));",
		},
		{
			"a + b * c + d / e - f;",
			"(((a + (b * c)) + (d / e)) - f);",
		},
		{
			"3 + 4; -5 * 5;",
			"(3 + 4);((-5) * 5);",
		},
		{
			"5 > 4 == 3 < 4;",
			"((5 > 4) == (3 < 4));",
		},
		{
			"5 < 4 != 3 > 4;",
			"((5 < 4) != (3 > 4));",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5;",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)));",
		},
		{
			"sacho;",
			"sacho;",
		},
		{
			"jhut;",
			"jhut;",
		},
		{
			"3 < 5 == jhut;",
			"((3 < 5) == jhut);",
		},
		{
			"3 > 5 == sacho;",
			"((3 > 5) == sacho);",
		},
		{
			"a + add(b * c) + d;",
			"((a + add((b * c))) + d);",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8));",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)));",
		},
		{
			"add(a + b + c * d / f + g);",
			"add((((a + b) + ((c * d) / f)) + g));",
		},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()

		if actual != tt.expected {
			t.Errorf("expected=[%q], got=[%q]", tt.expected, actual)
		}
	}
}

func TestYediStatementParsing(t *testing.T) {
	tests := []struct {
		program  string
		expected string
	}{
		{
			`yedi (2 + 3 == 5) {
		1;
		};`,
			"",
		},
		{
			`yedi (sacho) {
			23;
			} natra {
			43;
			};`,
			"",
		},
		{
			`yedi(jhut) {
				69;
			} nabhae (3 < 10) {
			 	420;
			} natra {
				1;
			}`,
			"",
		},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.program)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
		}
		_, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}
	}
}

func TestKaamGarStatement(t *testing.T) {
	tests := []struct {
		literal        string
		expectedParams []string
	}{
		{
			"kaam_gar(x, y) { patha x + y; };",
			[]string{"x", "y"},
		},
		{
			"kaam_gar() { patha 3; };",
			[]string{},
		},
		{
			"kaam_gar(a, b, c) { patha a + b + c; };",
			[]string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.literal)
		p := NewParser(l)
		program := p.ParseProgram()

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok || stmt == nil {
			t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement")
		}
		f, ok := stmt.Expression.(*ast.KaamGarExpression)
		if !ok || f == nil {
			t.Fatalf("stmt.Expression is not *ast.KaamGarExpression")
		}

		if len(f.Arguments) != len(tt.expectedParams) {
			t.Errorf("length parameters wrong. want %d, got=%d\n", len(tt.expectedParams), len(f.Arguments))
		}

		for i, ident := range tt.expectedParams {
			testLiteral(t, f.Arguments[i], ident)
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		statement string
		expected  string
	}{
		{
			`"foobar"`,
			"foobar",
		},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.statement)
		p := NewParser(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("expected 1 statement. got %d", len(program.Statements))
		}
	}
}

func TestArrayLength(t *testing.T) {
	tests := []struct {
		statement string
		expected  int
	}{
		{"[1,2,3]",
			3},
		{
			"[\"one\", \"two\", \"three\"]",
			3,
		},
		{
			"[\"udeshya\", \"dhungana\", kaam_gar(x) { x * x }]",
			3,
		},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.statement)
		p := NewParser(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("expected 1 statement. got %d", len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("statement not an expression statement")
		}
		arr, ok := stmt.Expression.(*ast.ArrayExpression)
		if !ok {
			t.Fatalf("expression not an array expression")
		}
		if len(arr.Elements) != tt.expected {
			t.Fatalf("array length does not match")
		}
	}
}

func TestHashIndex(t *testing.T) {
	tests := []struct {
		statement string
		index     int
		expected  any
	}{
		{"thoos arr = [1,2,3]; arr[0]",
			0,
			1,
		},
		{
			"thoos m = [\"one\", jhut, \"three\"]; m[1]",
			1,
			false,
		},
		{
			// 2 kept as a dummy
			"2; [\"udeshya\", \"dhungana\", 4][2]",
			2,
			4,
		},
		{
			`
			thoos x = {"foo": 49, "baz": 3};
			x[0]
			`,
			0,
			49,
		},
		{
			`
			2;
			{"udeshya": 49, "baz": 3}[1]
			`,
			1,
			3,
		},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.statement)
		p := NewParser(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 2 {
			t.Fatalf("expected 2 statements. got %d", len(program.Statements))
		}
		stmt, ok := program.Statements[1].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("statement not an expression statement")
		}
		_, ok = stmt.Expression.(*ast.IndexExpression)
		if !ok {
			t.Fatalf("expression not an array expression")
		}
	}
}

func TestJabasammaExpressionParsing(t *testing.T) {
	tests := []struct {
		program  string
		expected string
	}{
		{
			`jaba_samma (sacho) {
				1;
			}
			`,
			"",
		},
		{
			`jaba_samma (jhut) {
				23;
			}`,
			"",
		},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.program)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
		}
		es, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}
		_, ok = es.Expression.(*ast.JabasammaExpression)
		if !ok {
			t.Fatalf("expressiion not an ast.JabasammaExpression. got=%T", es.Expression)
		}
	}
}

func TestHashExpressionParsing(t *testing.T) {
	tests := []struct {
		program    string
		statements int
		pairs      int
	}{
		{
			`
			{"foo": "bar", "baz": 23};
			`,
			1,
			2,
		},
		{
			`
			{"foo": "bar", "baz": 23, "foo": "baz"};
			`,
			1,
			3,
		},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.program)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != tt.statements {
			t.Fatalf("expected %d statements, got=%d", tt.statements, len(program.Statements))
		}
		s, ok := program.Statements[tt.statements-1].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[%d] is not an expression statement", tt.statements-1)
		}
		h, ok := s.Expression.(*ast.HashExpression)
		if !ok {
			t.Fatalf("expression is not hash expression")
		}
		if len(h.Pairs) != tt.pairs {
			t.Fatalf("not enough pairs parsed. got=%d, expected=%d", len(h.Pairs), tt.pairs)
		}
	}
}

func TestComments(t *testing.T) {
	tests := []struct {
		statement string
		expected  string
	}{
		{
			`
			thoos x = 34;
			$ bich ko line $
			x
			`,
			"x",
		},
		{
			`
			thoos x = 34;
			$ bich ko line $
			x
			$ last ko line $
			`,
			"x",
		},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.statement)
		p := NewParser(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 2 {
			t.Fatalf("program.Statements not enough. want=%d, got=%d", 2, len(program.Statements))
		}
		s, ok := program.Statements[1].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[1] not expressionstatement, got=%T", s)
		}
		if s.Expression.String() != tt.expected {
			t.Fatalf("Expected `x` got=%s", s.Expression.String())
		}
	}
}

func TestCustom(t *testing.T) {
	tests := []struct {
		statement string
	}{
		{
			`
			$ recursion wala program $
			thoos recursion = kaam_gar(x) {
    			yedi (x == 0) {
        			patha 1;
    			} natra {
        			patha x * recursion(x - 1);
    			}
			};

			$ aba yo comment hale
			recursion(4);

			$ yo incomplete comment
			bhan("Hello world")
			`,
		},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.statement)
		p := NewParser(l)

		_ = p.ParseProgram()
		checkParserErrors(t, p)
	}
}
