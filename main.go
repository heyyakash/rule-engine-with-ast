package main

import (
	"fmt"

	"github.com/heyyakash/rule-engine-with-ast/helpers"
)

func main() {
	rule := "((age > 30 AND department = 'Sales') OR (age < 25 AND department = 'Marketing')) AND (salary > 50000 OR experience > 5)"
	tokens := helpers.TokenizeRule(rule)
	ast := helpers.CreateAST(tokens)
	fmt.Print(ast)
	// test := map[string]interface{}{
	// 	"age":        31,
	// 	"department": "Sales",
	// 	"salary":     5000000,
	// 	"experience": 8,
	// }

}
