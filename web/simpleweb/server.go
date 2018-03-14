package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
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

func main() {
	http.HandleFunc("/", askPin(HomePage))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
