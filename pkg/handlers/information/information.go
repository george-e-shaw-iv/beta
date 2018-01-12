package information

import (
	"net/http"
	"html/template"

	"github.com/george-e-shaw-iv/beta/pkg/handlers"
)

func Index(res http.ResponseWriter, _ *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.public.html",
		"templates/information/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(res, "layout", handlers.PageData{Title:"About the Fraternity"})
}

func Join(res http.ResponseWriter, _ *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.public.html",
		"templates/information/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(res, "layout", handlers.PageData{Title:"How to Join"})
}

func FAQ(res http.ResponseWriter, _ *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.public.html",
		"templates/information/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(res, "layout", handlers.PageData{Title:"FAQ"})
}

func KaiReport(res http.ResponseWriter, req *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.public.html",
		"templates/information/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(res, "layout", handlers.PageData{Title:"Elecontric Kai Report"})
}