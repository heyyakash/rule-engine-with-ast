package helpers_test

import (
	"testing"

	"github.com/heyyakash/rule-engine-with-ast/helpers"
)

type combineTestcase struct {
	Rules  []string
	Expect map[string]interface{}
}

var OutputCombinedMap = map[string]interface{}{
	"Type":     "LogicalExpression",
	"Operator": "OR",
	"Left": map[string]interface{}{
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
	},
	"Right": map[string]interface{}{
		"Type":     "LogicalExpression",
		"Operator": "AND",
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
					"Value": "'Marketing'",
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
					"Value": "20000",
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
	},
}

func TestCombination(t *testing.T) {
	testcases := []combineTestcase{
		{Rules: []string{"((age > 30 AND department = 'Sales') OR (age < 25 AND department = 'Marketing')) AND (salary > 50000 OR experience > 5)", "((age > 30 AND department = 'Marketing')) AND (salary > 20000 OR experience > 5)"}, Expect: OutputCombinedMap},
		{Rules: []string{"((age > 30 AND department = 'Sales') OR (age < 25 AND department = 'Marketing')) AND (salary > 50000 OR experience > 5)", "((age > 30 AND department = 'Marketing')) AND (salary > 20000 OR experience > 5)", "((age > 30 AND department = 'Marketing')) AND (salary > 20000 OR experience > 5)"}, Expect: OutputCombinedMap},
	}

	for i, tc := range testcases {
		combinedMap := helpers.AstToMap(helpers.CombineAsT(tc.Rules))
		if !CompareASTMap(combinedMap, tc.Expect) {
			t.Error("Combined rules not yielding correct result for ", i+1, " testcase")
		}
	}
}
