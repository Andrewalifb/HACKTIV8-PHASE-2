package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"webserver/models"
)

func GetHeroes(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT heroid, name, universe, skill, imageurl FROM heroes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var heroes []models.Heroes
	for rows.Next() {
		var hero models.Heroes

		err = rows.Scan(&hero.ID, &hero.Name, &hero.Universe, &hero.Skill, &hero.Imageurl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		heroes = append(heroes, hero)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}

func GetVillain(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT villainid, name, universe, imageurl FROM villain")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var villains []models.Villain
	for rows.Next() {
		var villain models.Villain

		err = rows.Scan(&villain.ID, &villain.Name, &villain.Universe, &villain.Imageurl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		villains = append(villains, villain)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(villains)
}