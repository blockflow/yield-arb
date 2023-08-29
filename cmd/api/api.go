package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func StartServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	http.ListenAndServe(":8080", r)
}
