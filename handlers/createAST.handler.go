package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/heyyakash/rule-engine-with-ast/configs"
	"github.com/heyyakash/rule-engine-with-ast/helpers"
)

type CreateRequest struct {
	Rule string `json:"rule"`
}

type Document struct {
	Rule string                 `json:"rule"`
	Tree map[string]interface{} `json:"tree"`
}

func CreateASTHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var req CreateRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Write([]byte("Couldn't parse data"))
		return
	}

	token := helpers.Tokenize(req.Rule)
	parser := helpers.NewParser(token)
	ast := parser.Parse()
	astMap := helpers.ASTToMAp(ast)
	if _, err := configs.ASTCollection.InsertOne(context.TODO(), Document{Rule: req.Rule, Tree: astMap}); err != nil {
		w.Write([]byte(fmt.Sprintf("Some error occured : %s", err)))
		return
	}

	w.Write([]byte("Rule Added Sucessfully"))

}
