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

func ASTToMAp(node *Node) map[string]interface{} {
	if node == nil {
		return nil
	}
	astMap := map[string]interface{}{
		"Type": node.Type,
	}

	if node.Name != "" {
		astMap["Name"] = node.Name
	}
	if node.Value != "" {
		astMap["Value"] = node.Value
	}
	if node.Operator != "" {
		astMap["Operator"] = node.Operator
	}
	if node.Left != nil {
		astMap["Left"] = ASTToMAp(node.Left)
	}
	if node.Right != nil {
		astMap["Right"] = ASTToMAp(node.Right)
	}

	return astMap

}

func MapToAST(astMap map[string]interface{}) *Node {
	node := &Node{
		Type: astMap["Type"].(string),
	}

	if operator, ok := astMap["Operator"].(string); ok {
		node.Operator = operator
	}
	if name, ok := astMap["Name"].(string); ok {
		node.Name = name
	}
	if value, ok := astMap["Value"].(string); ok {
		node.Value = value
	}

	if left, ok := astMap["Left"].(map[string]interface{}); ok {
		node.Left = MapToAST(left)
	}

	if right, ok := astMap["Right"].(map[string]interface{}); ok {
		node.Right = MapToAST(right)
	}

	return node
}

func CombineAsT(rules []string) *Node {
	var nodes []*Node
	for _, rule := range rules {
		tokens := Tokenize(rule)
		parser := NewParser(tokens)
		n := parser.Parse()
		nodes = append(nodes, n)
	}

	combinedAST := MergeALLAST(nodes)
	return combinedAST
}

func MergeALLAST(nodes []*Node) *Node {
	if len(nodes) == 0 {
		return nil
	}

	if len(nodes) == 1 {
		return nodes[0]
	}

	combined := nodes[0]

	for i := 1; i < len(nodes); i++ {
		combined = &Node{
			Type:     "LogicalExpression",
			Left:     combined,
			Right:    nodes[i],
			Operator: "OR",
		}
	}

	return combined
}
