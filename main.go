package main

import (
	"fmt"

	"github.com/heyyakash/rule-engine-with-ast/helpers"
)

func main() {
	rule := "((age > 30 AND department = 'Sales') OR (age < 25 AND department = 'Marketing'))"
	tokens := helpers.Tokenize(rule)
	// for _, v := range tokens {
	// 	fmt.Println(v)
	// }
	parser := helpers.NewParser(tokens)
	ast := parser.Parse()
	// helpers.PrintAST(ast, "")
	// Define a test map
	test := map[string]interface{}{
		"age":        31,
		"department": "Sales",
		"salary":     5000000,
		"experience": 8,
	}

	// Evaluate the AST against the test map
	result := helpers.Evaluate(ast, test)

	fmt.Println("Evaluation Result:", result) // Output: true or false
}
