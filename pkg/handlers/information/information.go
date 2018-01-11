package information

import (
	"net/http"
	"html/template"

	"github.com/george-e-shaw-iv/beta/pkg/handlers"
)

func Index(res http.ResponseWriter, req *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/information/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(res, "layout", handlers.PageData{Title:"About the Fraternity"})
}

func Join(res http.ResponseWriter, req *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/information/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(res, "layout", handlers.PageData{Title:"How to Join"})
}

func FAQ(res http.ResponseWriter, req *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/information/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(res, "layout", handlers.PageData{Title:"FAQ"})
}