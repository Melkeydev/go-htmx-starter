package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	tmplStr = "<div id=\"counter\">{{.CounterValue}}</div>"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.renderHomePage)
	r.Post("/increase", s.increaseCounter)
	r.Post("/decrease", s.decreaseCounter)
	r.Post("/add-film", s.handleFormSubmit)

	return r
}

func (s *Server) renderHomePage(w http.ResponseWriter, r *http.Request) {
	indexHTMLPath := filepath.Join("internal", "web", "index.html")
	tmpl, err := template.ParseFiles(indexHTMLPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	films := []Film{
		{
			Title:    "honeydew",
			Director: "melkey",
		},
		{
			Title:    "monkeydew",
			Director: "smih",
		},
		{
			Title:    "monkeydew",
			Director: "smih",
		},
	}

	data := map[string]any{
		"CounterValue": s.Counter.getValue(),
		"Films":        films,
	}

	tmpl.Execute(w, data)
}

type Film struct {
	Title    string
	Director string
}

type any = interface{}

func (s *Server) increaseCounter(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("counter").Parse(tmplStr))
	s.Counter.Increase()

	data := map[string]any{
		"CounterValue": s.Counter.getValue(),
	}

	tmpl.ExecuteTemplate(w, "counter", data)
}

func (s *Server) decreaseCounter(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("counter").Parse(tmplStr))
	s.Counter.Decrease()
	data := map[string]int{
		"CounterValue": s.Counter.getValue(),
	}

	tmpl.ExecuteTemplate(w, "counter", data)
}

func (s *Server) handleFormSubmit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmx request received")
	fmt.Println(r.Header.Get("HX-Request"))

	// extract the data
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")

	// we might want to add this data to a db
	// we need to return the form

	// we are using html-fragments
	indexHTMLPath := filepath.Join("internal", "web", "index.html")
	tmpl := template.Must(template.ParseFiles(indexHTMLPath))

	tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})

}
