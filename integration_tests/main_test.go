package integration_tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	handler "github.com/Fuerback/go-crud/api/handlers"
	"github.com/Fuerback/go-crud/business/domain"
	"github.com/Fuerback/go-crud/repositories"
	"github.com/Fuerback/go-crud/services"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var r *mux.Router
var savedID string

type Object struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestMain(m *testing.M) {
	db, err := sql.Open("sqlite3", "../db/mydb_test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = clearDB(db)
	if err != nil {
		log.Fatal(err)
	}

	respository := repositories.NewService(db)
	service := services.NewService(respository)

	r = mux.NewRouter()

	h := handler.New(service)
	h.Handlers(r)
	os.Exit(m.Run())
}

func clearDB(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from user_account")
	tx.Commit()
	return err
}

func TestGetAllUsersWhenEmptyTable(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/useraccount?limit=10&offset=0", nil)
	response := executeRequest(req)

	if err != nil {
		t.Fatalf("Error getting users: %s", err.Error())
	}

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]\n" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetUnknownUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/useraccount/khb40487", nil)
	response := executeRequest(req)

	if err != nil {
		t.Fatalf("Error getting user: %s", err.Error())
	}

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	if body := response.Body.String(); body != "{\"message\":\"sql: no rows in result set\"}" {
		t.Errorf("Expected an warning message. Got %s", body)
	}
}

func TestSaveUser(t *testing.T) {
	userDTO := getUserDTO()
	body, _ := json.Marshal(userDTO)
	requestReader := bytes.NewReader(body)

	req, err := http.NewRequest("POST", "/api/useraccount", requestReader)
	response := executeRequest(req)

	if err != nil {
		t.Fatalf("Error saving user: %s", err.Error())
	}

	checkResponseCode(t, http.StatusCreated, response.Code)

	if body := response.Body.String(); body != "" {
		t.Errorf("Expected an empty string. Got %s", body)
	}
}

func TestGetUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/useraccount?limit=10&offset=0", nil)
	response := executeRequest(req)

	if err != nil {
		t.Fatalf("Error getting users: %s", err.Error())
	}

	if body := response.Body.String(); body == "[]\n" {
		t.Errorf("Expected a list of objects. Got %s", body)
	}

	var object []Object
	err = json.Unmarshal(response.Body.Bytes(), &object)
	if err != nil {
		panic(err)
	}

	savedID = object[0].ID
	path := fmt.Sprintf("/api/useraccount/%s", savedID)
	req, err = http.NewRequest("GET", path, nil)
	response = executeRequest(req)

	if err != nil {
		t.Fatalf("Error getting user: %s", err.Error())
	}

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateUser(t *testing.T) {
	userDTO := getUserDTO()
	userDTO.Name = "changed name"
	body, _ := json.Marshal(userDTO)
	requestReader := bytes.NewReader(body)

	path := fmt.Sprintf("/api/useraccount/%s", savedID)
	req, err := http.NewRequest("PUT", path, requestReader)
	response := executeRequest(req)

	if err != nil {
		t.Fatalf("Error updating users: %s", err.Error())
	}

	checkResponseCode(t, http.StatusOK, response.Code)

	req, err = http.NewRequest("GET", path, nil)
	response = executeRequest(req)

	if err != nil {
		t.Fatalf("Error getting user: %s", err.Error())
	}

	checkResponseCode(t, http.StatusOK, response.Code)

	var object Object
	err = json.Unmarshal(response.Body.Bytes(), &object)
	if err != nil {
		panic(err)
	}

	if object.Name != "changed name" {
		t.Fatalf("Error on getting updated field: Expected %s, Actual %s", "changed name", object.Name)
	}
}

func TestDeletedUser(t *testing.T) {
	path := fmt.Sprintf("/api/useraccount/%s", savedID)
	req, err := http.NewRequest("DELETE", path, nil)
	response := executeRequest(req)

	if err != nil {
		t.Fatalf("Error deleting user: %s", err.Error())
	}

	checkResponseCode(t, http.StatusOK, response.Code)

	req, err = http.NewRequest("GET", path, nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	if body := response.Body.String(); body != "{\"message\":\"sql: no rows in result set\"}" {
		t.Errorf("Expected an warning message. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func getUserDTO() *domain.UserAccountRequestDto {
	return &domain.UserAccountRequestDto{
		Name:          "felipe",
		Email:         "felipe@gmail.com",
		Document:      "40843908744",
		Bank:          "Caixa",
		Agency:        43665,
		AgencyDigit:   4,
		AccountNumber: 2463456,
		AccountDigit:  4,
		AccountType:   "poupança",
		Status:        "rascunho",
	}
}
