package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//go:embed build/static
var static embed.FS

//go:embed html
var html embed.FS

var startDate = time.Now()

type PageData struct {
	Profile string
	Test    string
}

type StatusData struct {
	StartDate time.Time
	Status    string
}

func main() {
	addr, found := os.LookupEnv("SERVER_ADDRESS")
	if !found {
		slog.Info("Env var SERVER_ADDRESS not set, use default :8080")
		addr = ":8080"
	}

	profile, found := os.LookupEnv("PROFILE")
	if !found {
		slog.Info("Env var PROFILE not set, use default 'PROD'")
		profile = "PROD"
	}

	templates, err := template.New("").Funcs(template.FuncMap{
		// Templates func map
	}).ParseFS(html, "html/pages/*.html", "html/components/*.html")
	if err != nil {
		panic(err)
	}

	router := chi.NewMux()
	router.Use(middleware.CleanPath, middleware.Recoverer, middleware.Timeout(5*time.Second))
	// Logged routes
	router.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		staticDir, _ := fs.Sub(static, "build")
		r.Handle("/static/*", http.FileServerFS(staticDir))
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			err := templates.ExecuteTemplate(w, "index.html", PageData{
				Profile: profile,
				Test:    "bla",
			})
			if err != nil {
				slog.Error(err.Error())
			}
		})
	})

	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(StatusData{
			StartDate: startDate,
			Status:    "ok",
		})
	})

	slog.Info("Server listening on", "address", addr)
	http.ListenAndServe(addr, router)
}
