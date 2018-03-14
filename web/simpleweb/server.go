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
		log.Printf("template parsing errr: %v\n", err)
	}

	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Printf("template existing error: %v\n", err)
	}
}

func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
