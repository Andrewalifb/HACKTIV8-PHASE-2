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


	router.GET("/criminal_reports", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handlers.GetCriminalReports(db, w, r)
	})
	
	router.GET("/criminal_reports/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlers.GetCriminalReportByID(db, w, r, ps)
	})
	
	router.POST("/criminal_reports", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handlers.CreateCriminalReport(db, w, r)
	})
	
	router.PUT("/criminal_reports/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlers.UpdateCriminalReport(db, w, r, ps)
	})
	
	router.DELETE("/criminal_reports/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handlers.DeleteCriminalReport(db, w, r, ps)
	})

	fmt.Println("Running server on port :8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
