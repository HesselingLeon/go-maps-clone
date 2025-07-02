package main

import (
	"log"
	"net/http"

	"github.com/hesselingleon/go-map-clone/internal/config"
	"github.com/hesselingleon/go-map-clone/internal/database"
	"github.com/hesselingleon/go-map-clone/internal/router"
)

func main() {
	config := config.NewConfig()
	db, err := database.NewDatabase(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := router.NewRouter(db)

	log.Println("http://localhost:8080/docs")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
