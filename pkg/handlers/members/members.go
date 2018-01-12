package members

import (
	"net/http"
	"github.com/george-e-shaw-iv/beta/pkg/handlers"
	"html/template"
)

func Dashboard(res http.ResponseWriter, req *http.Request) {
	var tmp *template.Template

	pub := []string{
		"templates/layout.html",
		"templates/navbar.public.html",
		"templates/members/active/index.public.html",
	}

	priv := []string{
		"templates/layout.html",
		"templates/navbar.private.html",
		"templates/members/active/index.private.html",
	}

	_, err := req.Cookie("btp_active")
	if err != nil {
		tmp = template.Must(template.ParseFiles(pub...))
		tmp.ExecuteTemplate(res, "layout", handlers.PageData{Title:"Login"})
		return
	}

	tmp = template.Must(template.ParseFiles(priv...))
	tmp.ExecuteTemplate(res, "layout", handlers.PageData{Title:"Dashboard"})
}
