package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"routing/models"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Get all criminal reports
func GetCriminalReports(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, hero_id, villain_id, description, time_of_incident FROM criminal_reports")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var reports []models.CriminalReport
	for rows.Next() {
		var report models.CriminalReport
		err = rows.Scan(&report.ID, &report.HeroID, &report.VillainID, &report.Description, &report.TimeOfIncident)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reports = append(reports, report)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)

	log.Printf("Status Code: %d, Success: Showed all criminal reports\n", http.StatusOK)
}

// Get a criminal report by ID
func GetCriminalReportByID(db *sql.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT id, hero_id, villain_id, description, time_of_incident FROM criminal_reports WHERE id = $1", id)

	var report models.CriminalReport
	err = row.Scan(&report.ID, &report.HeroID, &report.VillainID, &report.Description, &report.TimeOfIncident)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No criminal report found with the given ID.", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)

	log.Printf("Status Code: %d, Success: Showed criminal reports by ID : %d\n", http.StatusOK, id)
}

// Create a criminal report
func CreateCriminalReport(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var report models.CriminalReport
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO criminal_reports (hero_id, villain_id, description, time_of_incident) VALUES ($1, $2, $3, $4) RETURNING id`
	err = db.QueryRow(query, report.HeroID, report.VillainID, report.Description, report.TimeOfIncident).Scan(&report.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)

	log.Printf("Status Code: %d, Success: Created Criminal Report with ID: %d\n", http.StatusOK, report.ID)
}

// Update a criminal report
func UpdateCriminalReport(db *sql.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var report models.CriminalReport
	err = json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE criminal_reports SET hero_id = $1, villain_id = $2, description = $3, time_of_incident = $4 WHERE id = $5 RETURNING id`
	err = db.QueryRow(query, report.HeroID, report.VillainID, report.Description, report.TimeOfIncident, id).Scan(&report.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)

	log.Printf("Status Code: %d, Success: Updated Criminal Report with ID: %d\n", http.StatusOK, report.ID)
}

// Delete a criminal report
func DeleteCriminalReport(db *sql.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `DELETE FROM criminal_reports WHERE id = $1`
	_, err = db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

	log.Printf("Status Code: %d, Success: Deleted Criminal Report with ID: %d\n", http.StatusOK, id)
}

