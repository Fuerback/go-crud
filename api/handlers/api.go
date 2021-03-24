package api

import (
	"encoding/json"
	"net/http"

	"github.com/Fuerback/go-crud/services"
	"github.com/julienschmidt/httprouter"
)

type Api struct{}

func (api *Api) Handlers(router *httprouter.Router) {
	router.GET("/api/useraccount/:id", api.getUser)
}

func (api *Api) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	if !validateParams(ps) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := services.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)

	w.WriteHeader(http.StatusOK)
	return
}

func validateParams(params httprouter.Params) bool {
	for _, param := range params {
		if param.Value == "" {
			return false
		}
	}
	return true
}
