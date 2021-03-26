package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Fuerback/go-crud/business/domain"
	"github.com/Fuerback/go-crud/services"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

type Handler struct {
	service services.UserAccountService
}

func New(s services.UserAccountService) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Handlers(r *mux.Router) {
	r.HandleFunc("/api/useraccount/{id}", h.getUser).Methods("GET")
	r.HandleFunc("/api/useraccount", h.saveUser).Methods("POST")
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	user, err := h.service.Get(params["id"])
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) saveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var u domain.UserAccountRequestDto

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatJSONError(err.Error()))
		return
	}

	v := validator.New()
	err = v.Struct(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatJSONError(err.Error()))
		return
	}

	//@TODO precisamos validar os dados antes de salvar na base de dados. Pergunta: Como fazer isso?
	err = h.service.Save(&u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(formatJSONError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
