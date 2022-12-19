package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "acceuil.html", nil)
}

func soonHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "cc.html", nil)
}

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", soonHandler)
	fmt.Println("server started on port 9999 => http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
