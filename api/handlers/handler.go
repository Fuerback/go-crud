package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Fuerback/go-crud/services"
	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	user, err := services.Get(params["id"])
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	json.NewEncoder(w).Encode(user)
}
