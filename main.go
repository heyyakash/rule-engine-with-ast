package main

import (
	"log"
	"net/http"

	"github.com/heyyakash/rule-engine-with-ast/configs"
	"github.com/heyyakash/rule-engine-with-ast/handlers"
)

func startServer() {
	http.HandleFunc("/create", handlers.CreateASTHandler)
	http.HandleFunc("/evaluate", handlers.EvaluateASTHandler)
	http.HandleFunc("/combine", handlers.CombineASTHandler)
	log.Print("Server Started")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Couldn't start server , ", err)
	}
}

func main() {
	// rule := "((age > 30 AND department = 'Sales') OR (age < 25 AND department = 'Marketing'))"
	// tokens := helpers.Tokenize(rule)
	// parser := helpers.NewParser(tokens)
	// ast := parser.Parse()
	// test := map[string]interface{}{
	// 	"age":        31,
	// 	"department": "Sales",
	// 	"salary":     5000000,
	// 	"experience": 8,
	// }

	// result := helpers.Evaluate(ast, test)
	// fmt.Println("Evaluation Result:", result)
	// astMap := helpers.ASTToMAp(ast)
	// _ = helpers.MapToAST(astMap)

	configs.ConnectDB()
	startServer()
}
