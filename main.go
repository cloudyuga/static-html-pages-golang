package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/blue", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/blue.html")
		tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/green", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/green.html")
		tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	fmt.Println("Server listening on port : 8080")
	http.ListenAndServe(":8080", nil)

}
