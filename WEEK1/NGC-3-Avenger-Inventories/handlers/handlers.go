package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"routing/models"

	"github.com/julienschmidt/httprouter"
)



func GetItems(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, item_code, stock, description, category_id, status, brand_id, location_id FROM items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item

		err = rows.Scan(&item.ID, &item.Name, &item.ItemCode, &item.Stock, &item.Description, &item.CategoryID, &item.Status, &item.BrandID, &item.LocationID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func GetItemByID(db *sql.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT id, name, item_code, stock, description, category_id, status, brand_id, location_id FROM items WHERE id = $1", id)

	var item models.Item
	err = row.Scan(&item.ID, &item.Name, &item.ItemCode, &item.Stock, &item.Description, &item.CategoryID, &item.Status, &item.BrandID, &item.LocationID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No item found with the given ID.", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}


func CreateItem(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var item models.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO items (name, item_code, stock, description, category_id, status, brand_id, location_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err = db.QueryRow(query, item.Name, item.ItemCode, item.Stock, item.Description, item.CategoryID, item.Status, item.BrandID, item.LocationID).Scan(&item.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}


func UpdateItem(db *sql.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var item models.Item
	err = json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE items SET name = $1, item_code = $2, stock = $3, description = $4, category_id = $5, status = $6, brand_id = $7, location_id = $8 WHERE id = $9 RETURNING id`
	err = db.QueryRow(query, item.Name, item.ItemCode, item.Stock, item.Description, item.CategoryID, item.Status, item.BrandID, item.LocationID, id).Scan(&item.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}


func DeleteItem(db *sql.DB, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `DELETE FROM items WHERE id = $1`
	_, err = db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
