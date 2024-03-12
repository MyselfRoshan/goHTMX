package router

import (
	"net/http"

	"github.com/MyselfRoshan/goHTMX/internals/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes() http.Handler {
	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// cors handeling
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	// Serve file inside /static folder
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static", fs))

	router.Get("/", handlers.HandleGetIndex)
	router.Post("/", handlers.HandlePostIndex)

	return router
}
