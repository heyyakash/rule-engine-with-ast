package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/heyyakash/rule-engine-with-ast/configs"
	"github.com/heyyakash/rule-engine-with-ast/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateRequest struct {
	Rule string `json:"rule"`
}

type Document struct {
	Rule string                 `json:"rule"`
	Tree map[string]interface{} `json:"tree"`
}

type response struct {
}

func CreateASTHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var req CreateRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Couldn't parse data", http.StatusBadRequest)
		return
	}

	token := helpers.Tokenize(req.Rule)
	parser := helpers.NewParser(token)
	ast := parser.Parse()
	astMap := helpers.ASTToMAp(ast)
	res, err := configs.ASTCollection.InsertOne(context.TODO(), Document{Rule: req.Rule, Tree: astMap})
	if err != nil {
		http.Error(w, "Internal Error Occured", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(res.InsertedID.(primitive.ObjectID).Hex()))

}
