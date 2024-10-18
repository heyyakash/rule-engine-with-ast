package helpers

import (
	"fmt"
	"strings"
)

type Node struct {
	Type     string
	Value    string
	Left     *Node
	Right    *Node
	Operator string
	Name     string
}

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) Peek() *Token {
	if p.pos >= len(p.tokens) {
		return nil
	}
	return &p.tokens[p.pos]
}

func (p *Parser) Consume() *Token {
	if p.pos >= len(p.tokens) {
		return nil
	}
	val := &p.tokens[p.pos]
	p.pos++
	return val
}

func (p *Parser) Parse() *Node {
	return p.parseExpression()
}

func (p *Parser) parseExpression() *Node {
	return p.parseAndOr()
}

func (p *Parser) parseAndOr() *Node {
	left := p.parseCondition()

	for {
		token := p.Peek()
		if token != nil && (token.Value == "AND" || token.Value == "OR") {
			operator := p.Consume().Value
			right := p.parseCondition()
			left = &Node{
				Type:     "LogicalExpression",
				Operator: operator,
				Left:     left,
				Right:    right,
			}
		} else {
			break
		}
	}

	return left
}

func (p *Parser) parseCondition() *Node {
	token := p.Peek()
	if token != nil && token.Value == "(" {
		p.Consume()
		expr := p.parseExpression()
		p.Consume()
		return expr
	} else {
		return p.parseComparison()
	}
}

func (p *Parser) parseComparison() *Node {
	left := p.parseOperand()
	operator := p.Consume().Value
	right := p.parseOperand()

	return &Node{
		Type:     "Comparison",
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}

func (p *Parser) parseOperand() *Node {
	token := p.Consume()

	if token.Type == "Literal" {
		return &Node{Type: "Literal", Value: strings.Trim(token.Value, "")}
	} else if token.Type == "Identifier" {
		return &Node{Type: "Identifier", Name: token.Value}
	}
	return nil
}

func PrintAST(node *Node, indent string) {
	if node == nil {
		return
	}
	fmt.Println(indent+"Type:", node.Type)
	if node.Operator != "" {
		fmt.Println(indent+"Operator:", node.Operator)
	}
	if node.Name != "" {
		fmt.Println(indent+"Operator:", node.Operator)
	}
	if node.Value != "" {
		fmt.Println(indent+"Value:", node.Value)
	}
	if node.Name != "" {
		fmt.Println(indent+"Name:", node.Name)
	}
	fmt.Println(indent + "Left:")
	PrintAST(node.Left, indent+"  ")
	fmt.Println(indent + "Right:")
	PrintAST(node.Right, indent+"  ")
}
