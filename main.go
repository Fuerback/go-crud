package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	handler "github.com/Fuerback/go-crud/api/handlers"
	"github.com/Fuerback/go-crud/repositories"
	"github.com/Fuerback/go-crud/services"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db/mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	respository := repositories.NewService(db)
	service := services.NewService(respository)

	r := mux.NewRouter()

	h := handler.New(service)
	h.Handlers(r)

	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))

}
