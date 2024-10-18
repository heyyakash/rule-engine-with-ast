package helpers

import (
	"fmt"
	"strconv"
)

func Evalute(node *Node, data map[string]interface{}) bool {
	if node.Type == Operand {
		return EvaluteCondition(node.Value, data)
	}

	if node.Value == "AND" {
		return Evalute(node.Left, data) && Evalute(node.Right, data)
	} else if node.Value == "OR" {
		return Evalute(node.Left, data) || Evalute(node.Right, data)
	}

	return false
}

func EvaluteCondition(condition string, data map[string]interface{}) bool {
	var field string
	var operator string
	var value string

	fmt.Sscanf(condition, "%s %s %s", &field, &operator, &value)
	fieldValue, exists := data[field]
	if !exists {
		return false
	}
	switch fieldValue.(type) {
	case int:
		intFieldValue := fieldValue.(int)
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return CompareIntValues(intFieldValue, intValue, operator)
	case string:
		strFieldValue := fieldValue.(string)
		return compareStringValues(strFieldValue, value, operator)

	}

	return false
}

func CompareIntValues(fieldValue, comparisonValue int, operator string) bool {
	switch operator {
	case ">":
		return fieldValue > comparisonValue
	case "<":
		return fieldValue < comparisonValue
	case "=":
		return fieldValue == comparisonValue
	case "<=":
		return fieldValue <= comparisonValue
	case ">=":
		return fieldValue >= comparisonValue
	default:
		return false
	}
}

func compareStringValues(fieldValue, comparisonValue, operator string) bool {
	if operator == "=" {
		return fieldValue == comparisonValue
	}
	return false
}
