package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	// "strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	ID           int64    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Quantity     int    `json:"quantity,omitempty"`
	Unit         string `json:"unit,omitempty"`
	BrandName    string `json:"-"`
}

var db *sql.DB

func main() {
	// Connect to SQLite database
	var err error
	db, err = sql.Open("sqlite3", "./inventory.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the products table for each brand
	brands := []string{"BrandA", "BrandB", "BrandC"}
	for _, brand := range brands {
		sqlStmt := `
			CREATE TABLE IF NOT EXISTS ` + brand + ` (
				product_id INTEGER PRIMARY KEY AUTOINCREMENT,
				product_name TEXT NOT NULL,
				quantity INTEGER NOT NULL,
				unit TEXT NOT NULL
			);
		`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Create and start the router
	r := mux.NewRouter()
	r.HandleFunc("/{brand}", getProductsHandler).Methods("GET")
	r.HandleFunc("/{brand}", createProductHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	brand := mux.Vars(r)["brand"]

	// Retrieve all products for the brand
	rows, err := db.Query("SELECT product_id, product_name, quantity, unit FROM "+brand)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Convert rows to JSON array
	products := []*Product{}
	for rows.Next() {
		product := &Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Quantity, &product.Unit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		product.BrandName = brand
		products = append(products, product)
	}
	err = rows.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	brand := mux.Vars(r)["brand"]

	// Decode the request body
	product := &Product{BrandName: brand}
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the new product
	sqlStmt := `
		INSERT INTO ` + brand + ` (product_name, quantity, unit)
		VALUES (?, ?, ?);
	`
	result, err := db.Exec(sqlStmt, product.Name, product.Quantity, product.Unit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	product.ID, err = result.LastInsertId()
	
	// Write the response
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
