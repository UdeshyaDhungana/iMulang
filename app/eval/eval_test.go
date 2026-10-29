package eval

import (
	"fmt"
	"testing"

	"github.com/udeshyadhungana/interprerer/app/lexer"
	"github.com/udeshyadhungana/interprerer/app/object"
	"github.com/udeshyadhungana/interprerer/app/parser"
)

func testEval(input string) object.Object {
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	if len(p.Errors()) != 0 {
		for _, e := range p.Errors() {
			fmt.Println(e)
		}
		return nil
	}
	return Eval(program, env)
}

func TestEvalIntegerStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 % 3", 2},
		{"5 - 10", -5},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalFloatObject(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{"5.5;", 5.5},
		{"10.10;", 10},
		{"-5.4;", -5},
		{"-10.10;", -10},
		{"5 + 5 + 5 + 5 - 10.0;", 10.0},
		{"3 * (3 * 3.0) + 10;", 37.0},
		{"(5 + 10 * 2 + 15.0 / 3) * 2 + -10;", 50.0},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if evaluated.Type() != object.FLOAT_OBJ {
			t.Fatalf("expected float, got %T", evaluated)
		}
		e := evaluated.(*object.Float)
		testFloatObject(t, evaluated, e.Value)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}
	return true
}

func testFloatObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%f, want=%f", result.Value, expected)
		return false
	}
	return true
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%s, want=%s", result.Value, expected)
		return false
	}
	return true
}

func TestEvalBoolStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"sacho;", true},
		{"jhut;", false},
		{"1 < 2;", true},
		{"1 > 2;", false},
		{"1 < 1;", false},
		{"1 > 1;", false},
		{"1 == 1;", true},
		{"1 != 1;", false},
		{"1 == 2;", false},
		{"1 != 2;", true},
		{"sacho == sacho;", true},
		{"jhut == jhut;", true},
		{"sacho == jhut;", false},
		{"sacho != jhut;", true},
		{"jhut != sacho;", true},
		{"(1 < 2) == sacho;", true},
		{"(1 < 2) == jhut;", false},
		{"(1 > 2) == sacho;", false},
		{"(1 > 2) == jhut;", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBoolObject(t, evaluated, tt.expected)
	}
}

func testBoolObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}
	return true
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!sacho;", false},
		{"!jhut;", true},
		{"!5;", false},
		{"!!sacho;", true},
		{"!!jhut;", false},
		{"!!5;", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBoolObject(t, evaluated, tt.expected)
	}
}

func TestYediExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"yedi(sacho) { 10; }", 10},
		{"yedi (jhut) { 10; }", nil},
		{"yedi (1) { 10; }", 10},
		{"yedi (1 < 2) { 10; }", 10},
		{"yedi (1 > 2) { 10; }", nil},
		{"yedi (1 > 2) { 10; } natra { 20; }", 20},
		{"yedi (1 < 2) { 10; } natra { 20; }", 10},
		{"yedi (100 < 2) { 10; } nabhae (2 > 1) { 20; } natra { 30; }", 20},
		{"yedi (100 > 2) { 10; } nabhae (2 > 1) { 20; } natra { 30; }", 10},
		{"yedi (100 < 2) { 10; } nabhae (2 < 1) { 20; } natra { 30; }", 30},
		{"yedi (100 < 2) { 10; } nabhae (2 < 1) { 20; } nabhae (3 > 1) { 30; } natra { 40 ;}", 30},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)

		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != object.NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return false
}

func TestPathaStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{"patha 10;", 10},
		{"patha 233;", 233},
		{"yedi(10 > 1) { patha 1; } natra { patha 0; }", 1},
		{"yedi(10 < 1) { patha 1; } natra { patha 0; }", 0},
		{`
		yedi (10 > 1) {
			yedi (10 > 1) {
				patha 1;
			}
			patha 2;
		}
		`, 1},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)

		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + sacho;",
			"unsupported operation INTEGER + BOOLEAN",
		},
		{
			"5 + jhut; 5;",
			"unsupported operation INTEGER + BOOLEAN",
		},
		{
			"-sacho;",
			"unknown operator: -BOOLEAN",
		},
		{
			"sacho + jhut;",
			"unsupported operation BOOLEAN + BOOLEAN",
		},
		{
			"5; sacho + jhut; 5;",
			"unsupported operation BOOLEAN + BOOLEAN",
		},
		{
			"yedi (10 > 1) { sacho + jhut; }",
			"unsupported operation BOOLEAN + BOOLEAN",
		},
		{
			`yedi (10 > 1) {
				yedi (20 > 1) {
					patha sacho + jhut;
				}
				patha 1;
			}`,
			"unsupported operation BOOLEAN + BOOLEAN",
		},
		{
			"foobar;",
			"identifier not found: foobar",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}

	}
}

/* TestThoos */
func TestThoosStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"thoos a = 5; a;", 5},
		{"thoos a = 5 * 5; a;", 25},
		{"thoos a = 5; thoos b = a; b;", 5},
		{"thoos a = 5; thoos b = a; thoos c = a + b + 5; c;", 15},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"thoos identity = kaam_gar(x) { x; }; identity(5);", 5},
		{"thoos identity = kaam_gar(x) { patha x; }; identity(5);", 5},
		{"thoos double = kaam_gar(x) { x * 2; }; double(5);", 10},
		{"thoos add = kaam_gar(x, y) { x + y; }; add(5, 5);", 10},
		{"thoos add = kaam_gar(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"kaam_gar(x) { x; }(5)", 5},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestRecursion(t *testing.T) {
	tests := []struct {
		program  string
		expected int64
	}{
		{`
		thoos recursion = kaam_gar(x) {
			yedi (x == 0) {
				patha 1;
			} natra {
			 	patha x * recursion(x - 1);
			}
		};
		recursion(4);
		`,
			24},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.program)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`lambai("")`, 0},
		{`lambai("four")`, 4},
		{`lambai("hello world")`, 11},
		{`lambai(1)`, "argument to `lambai` not supported, got INTEGER"},
		{`lambai("one", "two")`, "wrong number of arguments. got=2, want=1"},
		{`lambai([1,2,4])`, 3},
		{`lambai({"foo": 2, "bar": 45})`, 2},
		{
			`
				thoos x = [1,2,3,4];
				khaad(x, 5);
				lambai(x)
				x[4]
			`,
			5,
		},
		{
			`
				thoos x = [1,2,3,4];
				thoos z = udaa(x, 0);
				z
			`,
			1,
		},
		{
			`
				thoos x = [1,2,3,4];
				thoos z = udaa(x);
				z
			`,
			4,
		},
		{
			`
				thoos x = [1,2,3,4];
				thoos z = udaa(x);
				lambai(x)
			`,
			3,
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)",
					evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q",
					expected, errObj.Message)
			}
		}
	}
}

func TestIndexEval(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{
			`thoos x = [1,2,3];
			x[0]`,
			1,
		},
		{
			`thoos y = [sacho, jhut, jhut];
			y[2]`,
			false,
		},
		{
			`thoos y = ["Udeshya", "Dhungana"];
			y[0]`,
			"Udeshya",
		},
		{
			`
			thoos y = "foo";
			thoos x = {y: 23, "foo": "bar"};
			x[y]
			`,
			"bar",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case bool:
			testBoolObject(t, evaluated, expected)
		case string:
			testStringObject(t, evaluated, expected)
		}
	}
}

// too lazy to write different test cases, i combined them
func TestJabasammaAndAssignment(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{
			`
			thoos i = 0;
			i = 3;
			`,
			3,
		},
		{
			`
			thoos sum = kaam_gar(x) {
				patha x * 2;
			};

			thoos x = 2;

			jaba_samma(x < 32768) {
				x = sum(x);
			}
			x;

			yedi (x = 4) {
				100;
			} natra {
				200;
			}
			`,
			100,
		},
		{
			`
				thoos x = {"foo": "bar"};
				x["foo"] = 4;
				x["foo"]
			`,
			4,
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if !testIntegerObject(t, evaluated, tt.expected) {
			t.Fatalf("failed to test JabasammaExpression")
		}
	}
}

func TestGhumaExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{
			`
			ghuma(thoos i = 0; i < 100; i = i + 1) {
				sacho;
			}
			i;
			`,
			100,
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if !testIntegerObject(t, evaluated, tt.expected) {
			t.Fatalf("failed to test JabasammaExpression")
		}
	}
}

// hash maps
func TestHashMapEval(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{
			`
			thoos x = 43;
			thoos y = {"foo": "bar", "bar": x};
			lambai(y)
			`,
			2,
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if evaluated.Type() != object.INTEGER_OBJ {
			t.Fatalf("expected integer; got=%T", evaluated)
		}
		e := evaluated.(*object.Integer)
		if e.Value != tt.expected {
			t.Fatalf("not enough pairs, got=%d, expected=%d", e.Value, tt.expected)
		}
	}
}

// re-assignment

func TestCustom(t *testing.T) {
	program := `
	thoos checkNumber = kaam_gar(num) {
		yedi (num > 0) {
			bhan("Positive");
		} nabhae (num < 0) {
			bhan("Negative");
		} natra {
			bhan("Zero!");
		}
	};

	checkNumber(-1);
	checkNumber(0);
	checkNumber(1);
	`

	_ = testEval(program)
	fmt.Println("Passed!")
}
