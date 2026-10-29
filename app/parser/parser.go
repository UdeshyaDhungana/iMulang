package parser

import (
	"fmt"
	"strconv"

	"github.com/udeshyadhungana/interprerer/app/ast"
	"github.com/udeshyadhungana/interprerer/app/lexer"
	"github.com/udeshyadhungana/interprerer/app/token"
)

/* for pratt's */
type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string

	/* For pratt's parser */
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

const (
	_ int = iota
	LOWEST
	ASSIGN
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[token.TokenType]int{
	token.ASSIGN:   ASSIGN,
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.LT_EQ:    LESSGREATER,
	token.GT:       LESSGREATER,
	token.GT_EQ:    LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.MOD:      PRODUCT,
	token.LPAREN:   CALL,
	token.LBRACKET: CALL,
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDFIER, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.JHUT, p.parseBoolean)
	p.registerPrefix(token.SACHO, p.parseBoolean)
	p.registerPrefix(token.KAAM_GAR, p.parseKaamGar)
	p.registerPrefix(token.LPAREN, p.parseLeftParenthesis)
	p.registerPrefix(token.YEDI, p.parseYediExpression)
	p.registerPrefix(token.STRING, p.parseStringExpression)
	p.registerPrefix(token.LBRACKET, p.parseArrayExpression)
	p.registerPrefix(token.JABA_SAMMA, p.parseJabasammaExpression)
	p.registerPrefix(token.GHUMA, p.parseGhumaExpression)
	p.registerPrefix(token.LBRACE, p.parseHashExpression)

	// infix functions for operators
	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.MOD, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.GT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.LT_EQ, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.ASSIGN, p.parseInfixExpression)
	p.registerInfix(token.LPAREN, p.parseCallExpression)
	p.registerInfix(token.LBRACKET, p.parseIndexExpression)

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) CheckAndReportErrors() bool {
	if len(p.l.Errors()) > 0 {
		p.l.ReportErrors()
		return true
	}
	if len(p.errors) > 0 {
		p.reportErrors()
		return true
	}
	return false
}

func (p *Parser) reportErrors() {
	for _, s := range p.errors {
		fmt.Println(s)
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if len(p.errors) > 0 {
			return nil
		}
		program.Statements = append(program.Statements, stmt)
		p.nextToken()
	}
	return program
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	bs := &ast.BlockStatement{Token: p.curToken}
	s := []ast.Statement{}
	if !p.curTokenIs(token.LBRACE) {
		p.errors = append(p.errors, "expected { at the beginning of block statement")
		return nil
	}
	p.nextToken()
	for p.curToken.Type != token.RBRACE && p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		s = append(s, stmt)
		p.nextToken()
	}
	bs.Statements = s
	return bs
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.THOOS:
		return p.parseThoosStatement()
	case token.PATHA:
		return p.parsePathaStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseThoosStatement() *ast.ThoosStatement {
	stmt := &ast.ThoosStatement{Token: p.curToken}

	if !p.expectPeek(token.IDFIER) {
		p.errors = append(p.errors, "expected identifier after thoos")
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		p.errors = append(p.errors, "expected = after identifier")
		return nil
	}
	p.nextToken()

	stmt.Value = p.parseExpressionUsingPratt(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	if !p.curTokenIs(token.SEMICOLON) {
		p.errors = append(p.errors, "expected semicolon at the end of statement")
		return nil
	}

	return stmt
}

func (p *Parser) parsePathaStatement() *ast.PathaStatement {
	stmt := &ast.PathaStatement{Token: p.curToken}

	p.nextToken()
	stmt.Value = p.parseExpressionUsingPratt(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	if !p.curTokenIs(token.SEMICOLON) {
		p.errors = append(p.errors, "expected semicolon at the end of statement")
	}

	return stmt
}

func (p *Parser) parseYediExpression() ast.Expression {
	stmt := &ast.YediExpression{Token: p.curToken}
	if !p.curTokenIs(token.YEDI) {
		p.errors = append(p.errors, "expected yedi at the start of yedi expression")
		return nil
	}
	p.nextToken()
	if !p.curTokenIs(token.LPAREN) {
		p.errors = append(p.errors, "expected ( after yedi")
		return nil
	}
	// evaluate expression
	p.nextToken()
	stmt.Condition = p.parseExpressionUsingPratt(LOWEST)
	p.nextToken()
	if !p.curTokenIs(token.RPAREN) && !p.peekTokenIs(token.LBRACE) {
		p.errors = append(p.errors, "expected ) after condition and block statement after that")
		return nil
	}
	p.nextToken()

	stmt.Consequent = p.parseBlockStatement()

	// if statement with no consequences
	if !p.peekTokenIs(token.NATRA) && !p.peekTokenIs(token.NABHAE) {
		return stmt
	}

	// with alternatives
	for p.peekTokenIs(token.NABHAE) {
		p.nextToken()
		exp := p.parseNabhaeExpression()
		stmt.Alternatives = append(stmt.Alternatives, exp)
	}

	// with fallback
	if !p.expectPeek(token.NATRA) {
		return nil
	}
	p.nextToken()
	stmt.Fallback = p.parseBlockStatement()

	return stmt
}

func (p *Parser) parseNabhaeExpression() *ast.NabhaeExpression {
	res := &ast.NabhaeExpression{Token: p.curToken}
	if !p.curTokenIs(token.NABHAE) {
		p.errors = append(p.errors, "yedi alternative expression must begin with `nabhae`")
	}
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	p.nextToken()
	condition := p.parseExpressionUsingPratt(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	res.Condition = condition

	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	consequent := p.parseBlockStatement()
	res.Consequent = consequent
	return res
}

func (p *Parser) parseJabasammaExpression() ast.Expression {
	stmt := &ast.JabasammaExpression{Token: p.curToken}
	if !p.curTokenIs(token.JABA_SAMMA) {
		p.errors = append(p.errors, "jaba_samma expected")
		return nil
	}
	p.nextToken()
	if !p.curTokenIs(token.LPAREN) {
		p.errors = append(p.errors, "expected ( after jaba_samma")
		return nil
	}
	// parse condition
	p.nextToken()
	stmt.Condition = p.parseExpressionUsingPratt(LOWEST)
	p.nextToken()
	if !p.curTokenIs(token.RPAREN) && !p.peekTokenIs(token.LBRACE) {
		p.errors = append(p.errors, "expected sequence ) {")
		return nil
	}
	p.nextToken()
	stmt.Consequent = p.parseBlockStatement()
	return stmt
}

func (p *Parser) parseGhumaExpression() ast.Expression {
	stmt := &ast.GhumaExpression{Token: p.curToken}
	if !p.curTokenIs(token.GHUMA) {
		p.errors = append(p.errors, "expected: ghuma")
		return nil
	}
	p.nextToken()
	if !p.curTokenIs(token.LPAREN) {
		p.errors = append(p.errors, "expected '(' after ghuma")
		return nil
	}
	p.nextToken()
	stmt.Initialization = p.parseStatement()
	if !p.curTokenIs(token.SEMICOLON) {
		p.errors = append(p.errors, "expected semicolon after initialization")
		return nil
	}
	p.nextToken()
	stmt.Condition = p.parseStatement()
	if !p.curTokenIs(token.SEMICOLON) {
		p.errors = append(p.errors, "expected semicolon after condition")
		return nil
	}
	p.nextToken()
	stmt.Update = p.parseExpressionUsingPratt(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		p.errors = append(p.errors, "expected ')' after update in ghuma")
		return nil
	}
	if !p.expectPeek(token.LBRACE) {
		p.errors = append(p.errors, "expected { for ghuma body")
		return nil
	}
	stmt.Body = p.parseBlockStatement()
	return stmt
}

/* expressions */
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpressionUsingPratt(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// the most beautiful function of this program
func (p *Parser) parseExpressionUsingPratt(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseBoolean() ast.Expression {
	switch p.curToken.Type {
	case token.SACHO:
		return &ast.Boolean{Token: p.curToken, Value: true}
	case token.JHUT:
		return &ast.Boolean{Token: p.curToken, Value: false}
	default:
		panic(fmt.Sprintf("not a boolean: %s\n", p.curToken.Literal))
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	lit := &ast.FloatLiteral{Token: p.curToken}
	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for (%s) found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpressionUsingPratt(PREFIX)
	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Left:     left,
		Operator: p.curToken.Literal,
	}
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpressionUsingPratt(precedence)
	return expression
}

func (p *Parser) parseLeftParenthesis() ast.Expression {
	p.nextToken()
	exp := p.parseExpressionUsingPratt(LOWEST)
	if !p.expectPeek(token.RPAREN) {
		p.errors = append(p.errors, "mismatched parenthesis")
		return nil
	}
	return exp
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	result := ast.CallExpression{Token: p.curToken, Function: function}
	result.Arguments = p.parseArguments()
	return &result
}

func (p *Parser) parseArguments() []ast.Expression {
	args := []ast.Expression{}
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return nil
	}
	p.nextToken()
	args = append(args, p.parseExpressionUsingPratt(LOWEST))
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpressionUsingPratt(LOWEST))
	}
	if !p.expectPeek(token.RPAREN) {
		p.errors = append(p.errors, "expected ) after args list")
		return nil
	}
	return args
}

func (p *Parser) parseKaamGar() ast.Expression {
	result := ast.KaamGarExpression{Token: p.curToken}
	if !p.curTokenIs(token.KAAM_GAR) {
		p.errors = append(p.errors, "expected: kaam_gar")
		return nil
	}
	p.nextToken()
	if !p.curTokenIs(token.LPAREN) {
		p.errors = append(p.errors, "expected ( after kaam_gar")
		return nil
	}
	p.nextToken()
	if p.curTokenIs(token.RPAREN) {
		result.Arguments = nil
	} else {
		for {
			currentArg := p.parseIdentifier()
			r, ok := currentArg.(*ast.Identifier)
			if !ok {
				p.errors = append(p.errors, "parameters must be identifiers")
			}
			result.Arguments = append(result.Arguments, r)
			p.nextToken()
			if p.curTokenIs(token.RPAREN) {
				break
			}
			if !p.curTokenIs(token.COMMA) {
				p.errors = append(p.errors, "expected comma after argument")
				return nil
			}
			p.nextToken()
		}
	}
	p.nextToken()
	result.Body = p.parseBlockStatement()
	return &result
}

func (p *Parser) parseStringExpression() ast.Expression {
	return &ast.StringExpression{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseArrayExpression() ast.Expression {
	result := &ast.ArrayExpression{Token: p.curToken}
	var current ast.Expression

	if p.peekTokenIs(token.RBRACKET) {
		return result
	}
	p.nextToken()
	for {
		current = p.parseExpressionUsingPratt(LOWEST)
		result.Elements = append(result.Elements, current)
		p.nextToken()
		if p.curTokenIs(token.COMMA) {
			p.nextToken()
		} else if p.curTokenIs(token.RBRACKET) {
			break
		} else {
			p.errors = append(p.errors, "malformed array expression")
			return nil
		}
	}
	return result
}

func (p *Parser) parseIndexExpression(expr ast.Expression) ast.Expression {
	result := &ast.IndexExpression{Token: p.curToken, Operand: expr}
	p.nextToken()
	indexExpr := p.parseExpressionUsingPratt(LOWEST)
	p.nextToken()
	result.Index = indexExpr
	return result
}

func (p *Parser) parseHashExpression() ast.Expression {
	result := &ast.HashExpression{Token: p.curToken, Pairs: make(map[ast.Expression]ast.Expression)}
	for !p.curTokenIs(token.RBRACE) {
		p.nextToken()
		k := p.parseExpressionUsingPratt(LOWEST)
		if !p.expectPeek(token.COLON) {
			return nil
		}
		p.nextToken()
		v := p.parseExpressionUsingPratt(LOWEST)
		result.Pairs[k] = v
		p.nextToken()
	}
	return result
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}
