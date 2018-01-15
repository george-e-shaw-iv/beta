package members

import (
	"html/template"
	"net/http"

	"github.com/george-e-shaw-iv/beta/pkg/encryption"
	"github.com/george-e-shaw-iv/beta/pkg/handlers"
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

	if req.Method == "POST" {
		req.ParseForm()

		if req.Form["username"][0] == "root" && req.Form["password"][0] == "root" {
			c := http.Cookie{
				Name:  "btp_active",
				Value: encryption.RandomString(16),
				Path:  "/",
			}

			http.SetCookie(res, &c)
			tmp = template.Must(template.ParseFiles(priv...))
			tmp.ExecuteTemplate(res, "layout", handlers.PageData{Title: "Dashboard"})
		}
	}

	_, err := req.Cookie("btp_active")
	if err != nil {
		tmp = template.Must(template.ParseFiles(pub...))
		tmp.ExecuteTemplate(res, "layout", handlers.PageData{Title: "Login"})
		return
	}

	tmp = template.Must(template.ParseFiles(priv...))
	tmp.ExecuteTemplate(res, "layout", handlers.PageData{Title: "Dashboard"})
}
