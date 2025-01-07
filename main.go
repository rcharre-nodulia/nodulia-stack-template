package main

import (
	"embed"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//go:embed static
var static embed.FS

//go:embed html
var html embed.FS

type PageData struct {
	Locale string
}

func main() {
	addr, found := os.LookupEnv("FE_ADDRESS")
	if !found {
		slog.Info("Env var FE_ADDRESS not set, use default :8080")
		addr = ":8080"
	}

	templates, err := template.New("").Funcs(template.FuncMap{
		// Templates func map
	}).ParseFS(html, "html/pages/*.html", "html/components/*.html")
	if err != nil {
		panic(err)
	}

	router := chi.NewMux()
	router.Use(middleware.Logger)
	router.Use(middleware.CleanPath)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(5 * time.Second))

	router.Handle("/static/*", http.FileServerFS(static))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "index.html", PageData{})
		if err != nil {
			slog.Error(err.Error())
		}
	})

	slog.Info("Server listening on", "address", addr)
	http.ListenAndServe(addr, router)
}
