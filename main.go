package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var (
	tpl *template.Template
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	requestHandler()
}

func requestHandler() {
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", indexPage)
	fmt.Printf("Running server on port 8080...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func indexPage(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln("Template didn't execute: ", err)
	}
}
