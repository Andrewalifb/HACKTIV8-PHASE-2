package main

import (
	"fmt"
	"log"
	"net/http"
	"routing/config"
	"routing/handlers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	load := config.LoadDBConfig()

	db, err := config.ConnectDB(load)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := httprouter.New()


	router.GET("/items", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handlers.GetItems(db, w, r)
	})

	router.GET("/items/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlers.GetItemByID(db, w, r, ps)
	})

	router.POST("/items", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handlers.CreateItem(db, w, r)
	})

	router.PUT("/items/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlers.UpdateItem(db, w, r, ps)
	})

	router.DELETE("/items/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlers.DeleteItem(db, w, r, ps)
	})

	fmt.Println("Running server on port :8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
