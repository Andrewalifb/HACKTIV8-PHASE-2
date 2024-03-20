package main

import (
	"fmt"
	"log"
	"net/http"
	"webserver/config"
	"webserver/handlers"
)

func main() {
	load := config.LoadDBConfig()

	db, err := config.ConnectDB(load)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/get/heroes", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetHeroes(db, w, r)
	})

	mux.HandleFunc("/get/villain", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetVillain(db,w, r)
	})

	fmt.Println("Running server on port :8080")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server :", err.Error())
	}
}