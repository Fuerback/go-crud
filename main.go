package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db/mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := httprouter.New()
	handlers.Handlers(router)
	http.ListenAndServe(":8080", router)
}
