package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/heyyakash/rule-engine-with-ast/configs"
	"github.com/heyyakash/rule-engine-with-ast/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EvaluateRequest struct {
	RuleId string                 `json:"rule_id"`
	Data   map[string]interface{} `json:"data"`
}

type EvaluationResponse struct {
	Result bool `json:"result"`
}

func EvaluateASTHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var req EvaluateRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error occuered while parsing", http.StatusBadRequest)
		return
	}
	objectID, err := primitive.ObjectIDFromHex(req.RuleId)
	if err != nil {
		http.Error(w, "Error occuered while parsing", http.StatusBadRequest)
		return
	}
	var rule Document
	filter := bson.D{{Key: "_id", Value: objectID}}
	if err := configs.ASTCollection.FindOne(context.TODO(), filter).Decode(&rule); err != nil {
		log.Print(err)
		http.Error(w, "Rule not found", http.StatusNotFound)
		return
	}
	ast := helpers.MapToAST(rule.Tree)
	result := helpers.Evaluate(ast, req.Data)
	res, err := json.Marshal(&EvaluationResponse{Result: result})
	if err != nil {
		http.Error(w, "Some internal error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(res))
}
