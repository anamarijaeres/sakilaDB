package actors

import "github.com/go-chi/chi/v5"

func Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", ListActors)
	router.Post("/", CreateActor)
	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", ListActorsId)
		r.Put("/", UpdateActor)
		r.Delete("/", DeleteActor)
	})
	router.Get("/filmsById/{id}", ListFilmsByActorId)
	router.Get("/search", SearchActorsByName) // ?s=
	return router
}
