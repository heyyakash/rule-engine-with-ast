package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/heyyakash/rule-engine-with-ast/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type data struct {
	Id   string `json:"_id"`
	Rule string `json:"rule"`
}

func GetAllRules(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var res []data
	cur, err := configs.ASTCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Print(err)
		return
	}
	for cur.Next(context.TODO()) {
		var rule bson.M
		if err := cur.Decode(&rule); err != nil {
			log.Print(err)
			http.Error(w, "Error decoding rule", http.StatusInternalServerError)
			return
		}

		id := rule["_id"].(primitive.ObjectID).Hex()
		rulestring := rule["rule"].(string)
		res = append(res, data{
			Id:   id,
			Rule: rulestring,
		})
	}
	response, err := json.Marshal(&res)
	if err != nil {
		log.Print(err)
		http.Error(w, "Some internal error occured", http.StatusInternalServerError)
	}
	w.Write(response)

}
