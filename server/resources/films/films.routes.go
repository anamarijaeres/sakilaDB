package films

import "github.com/go-chi/chi/v5"

func Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", ListFilms)
	router.Post("/", CreateFilm) //w
	router.Get("/{id}", GetFilmById)

	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", GetFilmById)       //w
		r.Put("/", UpdateFilmById)    //w
		r.Delete("/", DeleteFilmById) //w
	})

	//router.Put("/pupdate/{id4}", PartialUpdate) //broken
	// router.Get("categories/{id5}", GetFilmCategoriesByFilmId)
	router.Get("/", SearchFilms)             // ?search=P
	router.Get("/rating", ListFilmsByRating) // ?rating=PG, G, R
	router.Get("/rent", ListAffordableFilms) //?rent=4

	router.Get("/store/{id1}", GetStoreByFilmId)
	router.Get("/store/address/{id3}", GetAddressByStoreId)

	router.Get("/category", ListFilmsByCategory)
	router.Get("/category/{id4}", GetFilmCategoriesByFilmId)
	return router
}
