package lexer

import (
	"testing"

	"github.com/udeshyadhungana/interprerer/app/token"
)

func TestNextToken(t *testing.T) {
	input := `thoos 界 = 5;
	thoos ten = 10;
	
	thoos add = kaam_gar(x, y) {
		x + y;
	};
	
	thoos result = add(five, ten);
	!-/*5;
	5 < 10 > 66;
	
	yedi (5 < 10) {
		patha sacho;
	} nabhae (3 < 10) {
	 	patha jhut;
	} natra {
	 	patha sacho;
	}
	
	10 == 10;
	10 != 9;

	"newstring"
	[newarray, "newArray"]

	jaba_samma(sacho) {
		jhut;
	}

	ghuma(thoos i = 0; i < 100; i = i + 1) {
		jhut;
	}

	{"foo": "bar"}
	69.69
	$sacho$
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.THOOS, "thoos"},
		{token.IDFIER, "界"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.THOOS, "thoos"},
		{token.IDFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.THOOS, "thoos"},
		{token.IDFIER, "add"},
		{token.ASSIGN, "="},
		{token.KAAM_GAR, "kaam_gar"},
		{token.LPAREN, "("},
		{token.IDFIER, "x"},
		{token.COMMA, ","},
		{token.IDFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDFIER, "x"},
		{token.PLUS, "+"},
		{token.IDFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.THOOS, "thoos"},
		{token.IDFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDFIER, "add"},
		{token.LPAREN, "("},
		{token.IDFIER, "five"},
		{token.COMMA, ","},
		{token.IDFIER, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "66"},
		{token.SEMICOLON, ";"},
		{token.YEDI, "yedi"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.PATHA, "patha"},
		{token.SACHO, "sacho"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.NABHAE, "nabhae"},
		{token.LPAREN, "("},
		{token.INT, "3"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.PATHA, "patha"},
		{token.JHUT, "jhut"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.NATRA, "natra"},
		{token.LBRACE, "{"},
		{token.PATHA, "patha"},
		{token.SACHO, "sacho"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.STRING, "newstring"},
		{token.LBRACKET, "["},
		{token.IDFIER, "newarray"},
		{token.COMMA, ","},
		{token.STRING, "newArray"},
		{token.RBRACKET, "]"},
		{token.JABA_SAMMA, "jaba_samma"},
		{token.LPAREN, "("},
		{token.SACHO, "sacho"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.JHUT, "jhut"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.GHUMA, "ghuma"},
		{token.LPAREN, "("},
		{token.THOOS, "thoos"},
		{token.IDFIER, "i"},
		{token.ASSIGN, "="},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.IDFIER, "i"},
		{token.LT, "<"},
		{token.INT, "100"},
		{token.SEMICOLON, ";"},
		{token.IDFIER, "i"},
		{token.ASSIGN, "="},
		{token.IDFIER, "i"},
		{token.PLUS, "+"},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.JHUT, "jhut"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.FLOAT, "69.69"},
		{token.EOF, ""},
	}
	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] failed - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] failed - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
