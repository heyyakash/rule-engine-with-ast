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
	ast := helpers.CombineAsT(req.Rules)
	astMap := helpers.ASTToMAp(ast)
	ruleString := strings.Join(req.Rules, " OR ")
	if _, err := configs.ASTCollection.InsertOne(context.TODO(), Document{Rule: ruleString, Tree: astMap}); err != nil {
		w.Write([]byte(fmt.Sprintf("Some error occured : %s", err)))
		return
	}

	w.Write([]byte("Rule Added Sucessfully"))

}
