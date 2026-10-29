package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/udeshyadhungana/interprerer/app/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Expression interface {
	Node
	expressionNode()
}

type Statement interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

/* Thoos statement */
type ThoosStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (tms *ThoosStatement) statementNode()       {}
func (tms *ThoosStatement) TokenLiteral() string { return tms.Token.Literal }

func (tms *ThoosStatement) String() string {
	var out bytes.Buffer

	out.WriteString(tms.TokenLiteral() + " ")
	out.WriteString(tms.Name.String())
	out.WriteString(" = ")

	if tms.Value != nil {
		out.WriteString(tms.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

/* Identifier */

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

/* Integer literal */
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

/* Float literal */
type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (il *FloatLiteral) expressionNode()      {}
func (il *FloatLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *FloatLiteral) String() string       { return il.Token.Literal }

/* Patha statement */
type PathaStatement struct {
	Token token.Token
	Value Expression
}

func (pms *PathaStatement) statementNode()       {}
func (pms *PathaStatement) TokenLiteral() string { return pms.Token.Literal }

func (pms *PathaStatement) String() string {
	var out bytes.Buffer

	out.WriteString(pms.TokenLiteral() + " ")
	if pms.Value != nil {
		out.WriteString(pms.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

/* Expression statement */
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	var out bytes.Buffer
	if es.Expression != nil {
		out.WriteString(es.Expression.String())
	}
	out.WriteString(";")
	return out.String()
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

// block
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (y *BlockStatement) statementNode()       {}
func (y *BlockStatement) TokenLiteral() string { return y.Token.Literal }
func (y *BlockStatement) String() string {
	var out bytes.Buffer

	out.WriteString("{\n")
	for _, s := range y.Statements {
		out.WriteString("\t")
		out.WriteString(s.String())
	}
	out.WriteString("\n}")

	return out.String()
}

// YEDI expression
type YediExpression struct {
	Token        token.Token
	Condition    Expression
	Consequent   *BlockStatement
	Alternatives []*NabhaeExpression
	Fallback     *BlockStatement
}

func (y *YediExpression) expressionNode()      {}
func (y *YediExpression) TokenLiteral() string { return y.Token.Literal }
func (y *YediExpression) String() string {
	var out bytes.Buffer

	out.WriteString("yedi (")
	out.WriteString(y.Condition.String())
	out.WriteString(") ")
	out.WriteString(y.Consequent.String())

	if y.Alternatives != nil {
		for _, a := range y.Alternatives {
			out.WriteString(a.String())
		}
	}

	if y.Fallback != nil {
		out.WriteString("natra ")
		out.WriteString(y.Fallback.String())
	}

	return out.String()
}

type NabhaeExpression struct {
	Token      token.Token
	Condition  Expression
	Consequent *BlockStatement
}

func (y *NabhaeExpression) expressionNode()      {}
func (y *NabhaeExpression) TokenLiteral() string { return y.Token.Literal }
func (y *NabhaeExpression) String() string {
	var out bytes.Buffer
	out.WriteString("nabhae (")
	out.WriteString(y.Condition.String())
	out.WriteString(")")
	out.WriteString(y.Consequent.String())
	return out.String()
}

// kaam_gar
type KaamGarExpression struct {
	Token     token.Token
	Arguments []*Identifier
	Body      *BlockStatement
}

func (f *KaamGarExpression) expressionNode()      {}
func (f *KaamGarExpression) TokenLiteral() string { return f.Token.Literal }
func (f *KaamGarExpression) String() string {
	var out bytes.Buffer

	out.WriteString("kaam_gar(")

	if f.Arguments != nil {
		for i, arg := range f.Arguments {
			out.WriteString((*arg).String())
			if i != len(f.Arguments)-1 {
				out.WriteString(",")
				out.WriteString(" ")
			}
		}
	}
	out.WriteString(") ")
	out.WriteString(f.Body.String())

	return out.String()
}

// call expression
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (f *CallExpression) expressionNode()      {}
func (f *CallExpression) TokenLiteral() string { return f.Token.Literal }
func (f *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range f.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(f.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

type StringExpression struct {
	Token token.Token
	Value string
}

func (s *StringExpression) expressionNode()      {}
func (s *StringExpression) TokenLiteral() string { return s.Token.Literal }
func (s *StringExpression) String() string {
	return fmt.Sprintf("\"%s\"", s.Value)
}

type ArrayExpression struct {
	Token    token.Token
	Elements []Expression
}

func (a *ArrayExpression) expressionNode()      {}
func (a *ArrayExpression) TokenLiteral() string { return a.Token.Literal }
func (a *ArrayExpression) String() string {
	var out bytes.Buffer

	out.WriteString("[")
	var eachString []string
	for _, e := range a.Elements {
		eachString = append(eachString, e.String())
	}
	out.WriteString(strings.Join(eachString, ", "))
	out.WriteString("]")
	return out.String()
}

type JabasammaExpression struct {
	Token      token.Token
	Condition  Expression
	Consequent *BlockStatement
}

func (c *JabasammaExpression) expressionNode() {}
func (c *JabasammaExpression) TokenLiteral() string {
	return c.Token.Literal
}
func (c *JabasammaExpression) String() string {
	var out bytes.Buffer
	out.WriteString("jaba_samma (")
	out.WriteString(c.Condition.String())
	out.WriteString(")")
	out.WriteString(c.Consequent.String())
	return out.String()
}

type GhumaExpression struct {
	Token          token.Token
	Initialization Statement
	Condition      Statement
	Update         Expression
	Body           *BlockStatement
}

func (g *GhumaExpression) expressionNode() {}
func (g *GhumaExpression) TokenLiteral() string {
	return g.Token.Literal
}
func (g *GhumaExpression) String() string {
	var out bytes.Buffer
	out.WriteString("ghuma (")
	out.WriteString(g.Initialization.String())
	out.WriteString(g.Condition.String())
	out.WriteString(g.Update.String())
	out.WriteString(")")
	out.WriteString(g.Body.String())
	return out.String()
}

/* Hash */
type HashExpression struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func (h *HashExpression) expressionNode() {}
func (h *HashExpression) TokenLiteral() string {
	return h.Token.Literal
}

func (h *HashExpression) String() string {
	var out bytes.Buffer
	out.WriteString("{")
	var inside []string
	for key, val := range h.Pairs {
		inside = append(inside, fmt.Sprintf("%s%s%s", key.String(), ": ", val.String()))
	}
	out.WriteString(strings.Join(inside, ", "))
	out.WriteString("}")
	return out.String()
}

// when you index an array or hashmap
type IndexExpression struct {
	Token   token.Token
	Operand Expression
	Index   Expression
}

func (a *IndexExpression) expressionNode()      {}
func (a *IndexExpression) TokenLiteral() string { return a.Token.Literal }
func (a *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString(a.Operand.String())
	out.WriteString("[")
	out.WriteString(fmt.Sprintf("%d", a.Index))
	out.WriteString("]")
	return out.String()
}
