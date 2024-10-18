package helpers

type Node struct {
	Type  TokenType
	Left  *Node
	Right *Node
	Value string
}

func Precedence(op string) int {
	switch op {
	case "AND":
		return 2
	case "OR":
		return 1
	}
	return 0
}

func ProcessOperator(operators *[]string, operands *[]*Node) {

	op := (*operators)[len(*operators)-1]
	*operators = (*operators)[:len(*operators)-1]

	right := (*operands)[len(*operands)-1]
	*operands = (*operands)[:len(*operands)-1]

	left := (*operands)[len(*operands)-1]
	*operands = (*operands)[:len(*operands)-1]

	newNode := &Node{
		Type:  Operator,
		Value: op,
		Left:  left,
		Right: right,
	}

	*operands = append(*operands, newNode)

}

func CreateAST(tokens []Token) *Node {
	operators := []string{}
	operands := []*Node{}

	for _, v := range tokens {
		if v.Type == Operand {
			operands = append(operands, &Node{Type: Operand, Value: v.Value})
		} else if v.Type == Parenthesis {
			if v.Value == "(" {
				operators = append(operators, v.Value)
			} else if v.Value == ")" {
				for len(operators) > 0 && operators[len(operators)-1] != "(" {
					ProcessOperator(&operators, &operands)
				}
				operators = operators[:len(operators)-1]
			}
		} else {
			for len(operators) > 0 && Precedence(operators[len(operators)-1]) >= Precedence(v.Value) {
				ProcessOperator(&operators, &operands)
			}
			operators = append(operators, v.Value)
		}
	}

	for len(operators) > 0 {
		ProcessOperator(&operators, &operands)
	}

	return operands[0]
}
