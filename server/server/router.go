package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"tsi.co/go-api2/resources/actors"
	"tsi.co/go-api2/resources/films"
)

func Router() chi.Router {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{AllowedOrigins: []string{"https://*", "http://*"}, AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, ExposedHeaders: []string{"Link"}, AllowCredentials: false, MaxAge: 300}))
	router.Mount("/actors", actors.Routes())
	router.Mount("/films", films.Routes())

	return router
}
