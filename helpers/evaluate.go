package helpers

import (
	"strconv"
	"strings"
)

func ConvertToFloat64(num interface{}) float64 {
	if n, ok := num.(int); ok {
		return float64(n)
	}
	return num.(float64)
}

func CompareNumbers(left, right interface{}) int {
	leftNum := ConvertToFloat64(left)
	rightNum, rightOk := right.(float64)
	if !rightOk {
		return 0
	}

	if leftNum > rightNum {
		return 1
	} else if leftNum < rightNum {
		return -1
	}
	return 0
}

func ResolveValues(node *Node, data map[string]interface{}) interface{} {
	if node.Type == "Identifier" {
		return data[node.Name]
	} else if node.Type == "Literal" {
		if num, err := strconv.ParseFloat(node.Value, 64); err == nil {
			return num
		}
		return strings.Trim(node.Value, "'")
	}
	return nil
}

func Evaluate(node *Node, data map[string]interface{}) bool {
	switch node.Type {
	case "LogicalExpression":
		left := Evaluate(node.Left, data)
		right := Evaluate(node.Right, data)

		if node.Operator == "AND" {
			return left && right
		} else if node.Operator == "OR" {
			return left || right
		}

	case "Comparison":
		leftVal := ResolveValues(node.Left, data)
		rightVal := ResolveValues(node.Right, data)

		switch node.Operator {
		case ">":
			return CompareNumbers(leftVal, rightVal) > 0
		case "<":
			return CompareNumbers(leftVal, rightVal) < 0
		case "=":
			return leftVal == rightVal
		}
	}
	return false
}
