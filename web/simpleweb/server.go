package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db  *sql.DB
	err error
)

// PageVariables helper struct
type PageVariables struct {
	Date string
	Time string
}

// HomePage display home.html template
func HomePage(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	HomePageVars := PageVariables{
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("tpl/home.html")
	if err != nil {
		log.Printf("template parsing error: %v\n", err)
	}

	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Printf("template existing error: %v\n", err)
	}
}

// curl http://localhost:1234?pin=1111
func askPin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("pin") != "1111" {
			http.Error(w, "wrong pin", http.StatusForbidden)
			return
		}
		h(w, r)
	}
}

// same as askPin but no hard coded pin value check
func askDatabaseForPin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pin string
		if err := db.QueryRow("SELECT pin FROM pins").Scan(&pin); err != nil {
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}
		if r.FormValue("pin") != pin {
			http.Error(w, "wrong pin", http.StatusForbidden)
			return
		}
		h(w, r)
	}
}

func main() {
	if db, err = sql.Open("sqlite3", "./pins.db"); err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/", askDatabaseForPin(HomePage))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
