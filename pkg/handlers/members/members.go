package members

import (
	"html/template"
	"net/http"

	"github.com/george-e-shaw-iv/beta/pkg/handlers"
	"github.com/george-e-shaw-iv/beta/pkg/database/models/user"
	"strconv"
	"strings"
	"time"
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

		roll, err := strconv.Atoi(req.Form["roll"][0])
		if err != nil {
			tmp = template.Must(template.ParseFiles(pub...))
			tmp.ExecuteTemplate(res, "layout", handlers.PageData{
				Title: "Login",
				Message: "Error parsing roll number. Try again.",
			})
			return
		}

		u, err := user.Fetch(roll)
		if err != nil {
			tmp = template.Must(template.ParseFiles(pub...))
			tmp.ExecuteTemplate(res, "layout", handlers.PageData{
				Title: "Login",
				Message: "Invalid roll number and/or password.",
			})
			return
		}

		err = u.Authenticate(req.Form["password"][0])
		if err != nil {
			tmp = template.Must(template.ParseFiles(pub...))
			tmp.ExecuteTemplate(res, "layout", handlers.PageData{
				Title: "Login",
				Message: "Invalid roll number and/or password.",
			})
			return
		}

		http.SetCookie(res, &http.Cookie{
			Name: "btp_active",
			Value: strconv.Itoa(u.Roll)+":"+u.Secret,
			Path: "/",
		})

		tmp = template.Must(template.ParseFiles(priv...))
		tmp.ExecuteTemplate(res, "layout", handlers.PageData{
			Title: "Dashboard",
			Message: "You've logged in successfully, welcome back " + u.FirstName + ".",
			ExternalData: u,
		})
		return
	}

	cookie, err := req.Cookie("btp_active")
	if err != nil {
		tmp = template.Must(template.ParseFiles(pub...))
		tmp.ExecuteTemplate(res, "layout", handlers.PageData{Title: "Login"})
		return
	}

	val := strings.Split(cookie.Value, ":")
	if len(val) != 2 {
		http.SetCookie(res, &http.Cookie{Name: "btp_active", Expires: time.Unix(0, 0)})
		tmp = template.Must(template.ParseFiles(pub...))
		tmp.ExecuteTemplate(res, "layout", handlers.PageData{Title: "Login"})
		return
	}

	roll, err := strconv.Atoi(val[0])
	if err != nil {
		http.SetCookie(res, &http.Cookie{Name: "btp_active", Expires: time.Unix(0, 0)})
		tmp = template.Must(template.ParseFiles(pub...))
		tmp.ExecuteTemplate(res, "layout", handlers.PageData{Title: "Login"})
		return
	}

	u, err := user.Fetch(roll)
	if err != nil {
		http.SetCookie(res, &http.Cookie{Name: "btp_active", Expires: time.Unix(0, 0)})
		tmp = template.Must(template.ParseFiles(pub...))
		tmp.ExecuteTemplate(res, "layout", handlers.PageData{Title: "Login"})
		return
	}

	if u.Secret != val[1] {
		http.SetCookie(res, &http.Cookie{Name: "btp_active", Expires: time.Unix(0, 0)})
		tmp = template.Must(template.ParseFiles(pub...))
		tmp.ExecuteTemplate(res, "layout", handlers.PageData{Title: "Login"})
		return
	}

	tmp = template.Must(template.ParseFiles(priv...))
	tmp.ExecuteTemplate(res, "layout", handlers.PageData{Title: "Dashboard", ExternalData: u})
}
