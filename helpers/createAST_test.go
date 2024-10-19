package helpers_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/heyyakash/rule-engine-with-ast/helpers"
)

var TestOutputMap = map[string]interface{}{
	"Type":     "LogicalExpression",
	"Operator": "AND",
	"Left": map[string]interface{}{
		"Type":     "LogicalExpression",
		"Operator": "OR",
		"Left": map[string]interface{}{
			"Type":     "LogicalExpression",
			"Operator": "AND",
			"Left": map[string]interface{}{
				"Type":     "Comparison",
				"Operator": ">",
				"Left": map[string]interface{}{
					"Type": "Identifier",
					"Name": "age",
				},
				"Right": map[string]interface{}{
					"Type":  "Literal",
					"Value": "30",
				},
			},
			"Right": map[string]interface{}{
				"Type":     "Comparison",
				"Operator": "=",
				"Left": map[string]interface{}{
					"Type": "Identifier",
					"Name": "department",
				},
				"Right": map[string]interface{}{
					"Type":  "Literal",
					"Value": "'Sales'",
				},
			},
		},
		"Right": map[string]interface{}{
			"Type":     "LogicalExpression",
			"Operator": "AND",
			"Left": map[string]interface{}{
				"Type":     "Comparison",
				"Operator": "<",
				"Left": map[string]interface{}{
					"Type": "Identifier",
					"Name": "age",
				},
				"Right": map[string]interface{}{
					"Type":  "Literal",
					"Value": "25",
				},
			},
			"Right": map[string]interface{}{
				"Type":     "Comparison",
				"Operator": "=",
				"Left": map[string]interface{}{
					"Type": "Identifier",
					"Name": "department",
				},
				"Right": map[string]interface{}{
					"Type":  "Literal",
					"Value": "'Marketing'",
				},
			},
		},
	},
	"Right": map[string]interface{}{
		"Type":     "LogicalExpression",
		"Operator": "OR",
		"Left": map[string]interface{}{
			"Type":     "Comparison",
			"Operator": ">",
			"Left": map[string]interface{}{
				"Type": "Identifier",
				"Name": "salary",
			},
			"Right": map[string]interface{}{
				"Type":  "Literal",
				"Value": "50000",
			},
		},
		"Right": map[string]interface{}{
			"Type":     "Comparison",
			"Operator": ">",
			"Left": map[string]interface{}{
				"Type": "Identifier",
				"Name": "experience",
			},
			"Right": map[string]interface{}{
				"Type":  "Literal",
				"Value": "5",
			},
		},
	},
}

func TestCreateAST(t *testing.T) {
	testRule := "((age > 30 AND department = 'Sales') OR (age < 25 AND department = 'Marketing')) AND (salary > 50000 OR experience > 5)"
	tokens := helpers.Tokenize(testRule)
	parser := helpers.NewParser(tokens)
	astMap := helpers.AstToMap(parser.Parse())
	if !CompareASTMap(astMap, TestOutputMap) {
		t.Error("AST creation did not yield expected result")
	}

}

func CompareASTMap(input, target map[string]interface{}) bool {

	if len(input) != len(target) {
		return false
	}

	for key, value := range input {
		value2, ok := target[key]
		if !ok {
			return false
		}

		switch v1 := value.(type) {
		case map[string]interface{}:
			v2, ok := value2.(map[string]interface{})
			if !ok || !CompareASTMap(v1, v2) {
				fmt.Printf("map mismatch\n")
				return false
			}

		case string:
			v2, ok := value2.(string)
			if !ok || v1 != v2 {
				fmt.Printf("string mismatch input (%v) != output (%v)\n", v1, v2)
				return false
			}

		case int:
			v2, ok := value2.(int)
			if !ok || v1 != v2 {
				fmt.Printf("int mismatch input (%v) != output (%v)\n", v1, v2)
				return false
			}

		case float64:
			v2, ok := value2.(float64)
			if !ok || v1 != v2 {
				fmt.Printf("float mismatch input (%v) != output (%v)\n", v1, v2)
				return false
			}

		default:
			if !reflect.DeepEqual(value, value2) {
				return false
			}
		}
	}

	return true

}
