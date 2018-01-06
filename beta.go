package main

import (
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	http.ListenAndServe("127.0.0.1:3000", mux)

	mux.HandleFunc("/", index)
}

func index(res http.ResponseWriter, req *http.Request) {
	files := []string{
		"templates/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(res, "index", nil)
}
