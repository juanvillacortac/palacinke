package ast

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/juandroid007/palacinke/pkg/token"
)

type Node interface {
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

func Json(node Node) ([]byte, error) {
	var out bytes.Buffer
	str, err := json.Marshal(node)
	if err != nil {
		return nil, err
	}
	err = json.Indent(&out, str, "", "\t")
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

/***** Root statement ********************************************************/

type Program struct {
	Statements []Statement `json:"statements"`
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

/***** Identifier ************************************************************/

type Identifier struct {
	Token token.Token `json:"token"` // the token.IDENT
	Value string      `json:"value"`
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

/***** Let statement *********************************************************/

type LetStatement struct {
	Token token.Token `json:"token"` // the token.LET token
	Name  *Identifier `json:"name"`
	Value Expression  `json:"value"`
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" " + token.ASSIGN + " ") // "="

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(token.SEMICOLON) // ";"

	return out.String()
}

/***** Return statement ******************************************************/

type ReturnStatement struct {
	Token       token.Token `json:"token"` // the 'return' token
	ReturnValue Expression  `json:"return_value"`
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(token.SEMICOLON) // ";"

	return out.String()
}

/***** Expression statement **************************************************/

type ExpressionStatement struct {
	Token      token.Token `json:"token"`
	Expression Expression  `json:"expression"`
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	var out bytes.Buffer
	if es.Expression != nil {
		out.WriteString(es.Expression.String())
	}
	return out.String()
}

/***** Integer literal *******************************************************/

type IntegerLiteral struct {
	Token token.Token `json:"token"`
	Value int64       `json:"value"`
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

/***** String literal ********************************************************/

type StringLiteral struct {
	Token token.Token `json:"token"`
	Value string      `json:"value"`
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

/***** Prefix expression *****************************************************/

type PrefixExpression struct {
	Token    token.Token `json:"token"` // The prefix token, like ! or -
	Operator string      `json:"operator"`
	Right    Expression  `json:"right"`
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

/***** Infix expression ******************************************************/

type InfixExpression struct {
	Token    token.Token `json:"token"` // The operator token, like +, -, * or /
	Left     Expression  `json:"left"`
	Operator string      `json:"operator"`
	Right    Expression  `json:"right"`
}

func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

/***** Boolean ***************************************************************/

type Boolean struct {
	Token token.Token `json:"token"`
	Value bool        `json:"value"`
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

/***** Block statement *******************************************************/

type BlockStatement struct {
	Token      token.Token `json:"token"` // the { token
	Statements []Statement `json:"statements"`
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for i, s := range bs.Statements {
		if i == len(bs.Statements)-1 {
			out.WriteString(s.String())
		} else {
			out.WriteString(s.String() + " ")
		}
	}

	return out.String()
}

/***** If expression *********************************************************/

type IfExpression struct {
	Token       token.Token     `json:"token"` // The 'if' token
	Condition   Expression      `json:"condition"`
	Consequence *BlockStatement `json:"consequence"`
	Alternative *BlockStatement `json:"alternative"`
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if ")
	out.WriteString(ie.Condition.String() + " ")
	out.WriteString("{ ")
	out.WriteString(ie.Consequence.String())
	out.WriteString("}")

	if ie.Alternative != nil {
		out.WriteString(" else { ")
		out.WriteString(ie.Alternative.String())
		out.WriteString(" }")
	}

	return out.String()
}

/***** Function literal ******************************************************/

type FunctionLiteral struct {
	Token      token.Token     `json:"token"` // The 'fn' token
	Parameters []*Identifier   `json:"parameters"`
	Body       *BlockStatement `json:"body"`
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

/***** Call expression *******************************************************/

type CallExpression struct {
	Token     token.Token  `json:"token"`    // The '(' token
	Function  Expression   `json:"function"` // Identifier/FunctionLiteral
	Arguments []Expression `json:"arguments"`
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, arg := range ce.Arguments {
		args = append(args, arg.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
