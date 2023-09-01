package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func StartServer() {
	r := chi.NewRouter() // Basic CORS setup
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allowed origins, can be more specific like: []string{"http://localhost:3000"}
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	// Apply CORS middleware to the router
	r.Use(corsMiddleware.Handler)
	r.Use(middleware.Logger)

	r.Get("/", test)
	r.Get("/strats", getStrats)
	r.Post("/transactions", getTransactions)

	http.ListenAndServe(":8080", r)
}
