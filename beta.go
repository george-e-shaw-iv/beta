package main

import (
	"html/template"
	"net/http"
	"github.com/george-e-shaw-iv/beta/pkg/handlers/information"
	"log"
	"github.com/george-e-shaw-iv/beta/pkg/handlers/members"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))

	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	mux.HandleFunc("/information", information.Index)
	mux.HandleFunc("/information/join", information.Join)
	mux.HandleFunc("/information/faq", information.FAQ)
	mux.HandleFunc("/information/chapter-resources/kai_report", information.KaiReport)

	mux.HandleFunc("/members/active", members.Dashboard)

	log.Println("Server listening at localhost:3000 - CTRL+C to exit")
	http.ListenAndServe("127.0.0.1:3000", mux)
}

func index(res http.ResponseWriter, req *http.Request) {
	if len(req.URL.Path[1:]) > 0 {
		http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	files := []string{
		"templates/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(res, "index", nil)
}
