package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

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
	r.HandleFunc("/api/useraccount", h.getUsers).Methods("GET")
	r.HandleFunc("/api/useraccount/{id}", h.getUser).Methods("GET")
	r.HandleFunc("/api/useraccount", h.saveUser).Methods("POST")
	r.HandleFunc("/api/useraccount/{id}", h.updateUser).Methods("PUT")
	r.HandleFunc("/api/useraccount/{id}", h.deleteUser).Methods("DELETE")
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	limit, _ := strconv.Atoi(r.FormValue("limit"))
	offset, _ := strconv.Atoi(r.FormValue("offset"))

	p := &domain.PaginatorDTO{
		Offset: offset,
		Limit:  limit,
	}
	users, err := h.service.GetAll(*p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(formatJSONError(err.Error()))
		return

	}
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	err := h.service.Delete(params["id"])
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(formatJSONError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	ID := params["id"]

	var user domain.UserAccountRequestDto

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatJSONError(err.Error()))
		return
	}

	v := validator.New()
	err = v.Struct(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatJSONError(err.Error()))
		return
	}

	err = h.service.Update(ID, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(formatJSONError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	user, err := h.service.Get(params["id"])
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(formatJSONError(err.Error()))
		return
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

	err = h.service.Save(&u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(formatJSONError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
