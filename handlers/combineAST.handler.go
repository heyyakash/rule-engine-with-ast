package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/heyyakash/rule-engine-with-ast/configs"
	"github.com/heyyakash/rule-engine-with-ast/helpers"
)

type CombineRequest struct {
	Rules []string `json:"rules"`
}

func CombineASTHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var req CombineRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Write([]byte("Couldn't parse data"))
		return
	}
	ruleSet := helpers.GenerateSet(req.Rules)
	for _, v := range ruleSet {
		if err := helpers.ValidateRule(v); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	ast := helpers.CombineAsT(ruleSet)
	astMap := helpers.AstToMap(ast)
	ruleString := strings.Join(ruleSet, " OR ")
	if _, err := configs.ASTCollection.InsertOne(context.TODO(), Document{Rule: ruleString, Tree: astMap}); err != nil {
		w.Write([]byte(fmt.Sprintf("Some error occured : %s", err)))
		return
	}

	w.Write([]byte("Rule Added Sucessfully"))

}
