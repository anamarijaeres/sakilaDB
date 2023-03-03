package films

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	db "tsi.co/go-api2/database"
	e "tsi.co/go-api2/error"

	m "tsi.co/go-api2/resources/models"
)

func ListFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var films []*m.Film
	db.DB.Model(&m.Film{}).Preload("Actors").Find(&films)

	render.RenderList(w, r, m.NewFilmListResponse(films))

}

func CreateFilm(w http.ResponseWriter, r *http.Request) {
	var data m.FilmRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	film := data.Film
	log.Print(film)

	e := db.DB.Create(film)
	if e.Error == nil {
		render.Status(r, http.StatusCreated)
	} else {
		render.Status(r, http.StatusForbidden)
	}
	log.Print(e.Error)

	render.Render(w, r, m.NewFilmResponse(film))
}

func ListFilmsByRating(w http.ResponseWriter, r *http.Request) {

	//log.Print(ratingIdParam)
	var films []*m.Film

	ratingQuery := r.URL.Query().Get("rating")

	log.Print(ratingQuery)

	db.DB.Where("rating = ?", ratingQuery).Find(&films)

	render.RenderList(w, r, m.NewFilmListResponse(films))
}

func ListAffordableFilms(w http.ResponseWriter, r *http.Request) {
	//log.Print(ratingIdParam)

	var films []*m.Film

	rentQuery := r.URL.Query().Get("rent")

	log.Print(rentQuery)

	db.DB.Where("rental_rate <= ?", rentQuery).Find(&films)

	render.RenderList(w, r, m.NewFilmListResponse(films))
}

func GetFilmById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	var film *m.Film

	db.DB.Where("film_id = ?", idParam).Find(&film)
	render.Render(w, r, m.NewFilmResponse(film))
}

func GetStoreByFilmId(w http.ResponseWriter, r *http.Request) {
	filmIdParam := chi.URLParam(r, "id1")

	var stores []*m.Store
	var inventory []*m.Inventory

	db.DB.Where("film_id = ?", filmIdParam).Find(&inventory)
	log.Print(inventory)

	for _, inv := range inventory {
		var store *m.Store
		store_id := inv.StoreId
		db.DB.Where("store_id = ?", store_id).Find(&store)
		stores = append(stores, store)

	}

	render.RenderList(w, r, m.NewStoreListResponse(stores))

}

func GetAddressByStoreId(w http.ResponseWriter, r *http.Request) {
	storeId := chi.URLParam(r, "id3")
	log.Print(storeId)
	var store *m.Store
	db.DB.Where("store_id =?", storeId).Find(&store)
	var address *m.Address
	db.DB.Where("address_id=?", store.AddressId).Find(&address)

	render.Render(w, r, m.NewAddressResponse(address))
}

func SearchFilms(w http.ResponseWriter, r *http.Request) {
	var films []*m.Film
	q := r.URL.Query().Get("search")

	q = q + "%"
	log.Print(q)

	db.DB.Where("title LIKE ?", q).Find(&films)

	render.RenderList(w, r, m.NewFilmListResponse(films))
}

func UpdateFilmById(w http.ResponseWriter, r *http.Request) {
	var data m.FilmRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	idParam := chi.URLParam(r, "id")
	film := data.Film
	film.LastUpdate = time.Now()

	if db.DB.Model(&film).Where("film_id = ?", idParam).Updates(&film).RowsAffected == 0 {
		db.DB.Create(&film)
	}

	render.Status(r, http.StatusAccepted)
	render.Render(w, r, m.NewFilmResponse(film))
}
func PartialUpdate(w http.ResponseWriter, r *http.Request) {
	// Get the actor ID from the URL parameter.
	filmID := chi.URLParam(r, "id4")
	var film m.Film

	db.DB.First(&film, filmID)
	json.NewDecoder(r.Body).Decode(&film)

	log.Print(film)

	db.DB.Save(&film)
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(film)

}

func DeleteFilmById(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	var film *m.Film

	e := db.DB.Delete(&film, idParam)

	if e.Error == nil {
		render.Status(r, http.StatusAccepted)
	} else {
		render.Status(r, http.StatusForbidden)
	}
	log.Print(e.Error)
	render.Render(w, r, m.NewFilmResponse(film))
}

func ListFilmsByCategory(w http.ResponseWriter, r *http.Request) {

	//log.Print(ratingIdParam)
	var films []*m.Film
	var category *m.Category
	var filmcategory []*m.FilmCategory

	cQuery := r.URL.Query().Get("c")

	log.Print(cQuery)

	db.DB.Where("name = ?", cQuery).Find(&category)

	db.DB.Where("category_id = ?", category.CategoryId).Find(&filmcategory)

	for _, c := range filmcategory {
		var film *m.Film

		db.DB.Where("film_id = ?", c.FilmId).Find(&film)

		films = append(films, film)

	}
	render.RenderList(w, r, m.NewFilmListResponse(films))
}

// func GetFilmCategoriesByFilmId(w http.ResponseWriter, r *http.Request) {
// 	var categories []*m.Category

// 	filmID := chi.URLParam(r, "id4")

// 	log.Print(filmID)

// 	db.DB.Where("film_id = ?", filmID).Find(&category)

// 	db.DB.Where("category_id = ?", category.CategoryId).Find(&filmcategory)

// }
