package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	router "github.com/Fuerback/go-crud/api/routers"
	"github.com/Fuerback/go-crud/services"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db/mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service := services.NewService(db)

	r := router.Router()

	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))

}
