package server

import (
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

	return r
}

func (s *Server) renderHomePage(w http.ResponseWriter, r *http.Request) {
	indexHTMLPath := filepath.Join("internal", "web", "index.html")
	tmpl, err := template.ParseFiles(indexHTMLPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]int{
		"CounterValue": s.Counter.getValue(),
	}

	tmpl.Execute(w, data)
}

func (s *Server) increaseCounter(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("counter").Parse(tmplStr))
	s.Counter.Increase()
	data := map[string]int{
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
