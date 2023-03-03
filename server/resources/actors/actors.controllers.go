package actors

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	db "tsi.co/go-api2/database"
	e "tsi.co/go-api2/error"
	m "tsi.co/go-api2/resources/models"
)

func ListActors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var actors []*m.Actor

	db.DB.Find(&actors)

	render.RenderList(w, r, m.NewActorListResponse(actors))

}

func ListActorsId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	idParam := chi.URLParam(r, "id")

	var actor *m.Actor

	db.DB.Where("actor_id = ?", idParam).Find(&actor)
	render.Render(w, r, m.NewActorResponse(actor))

}

func SearchActorsByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var actors []*m.Actor
	q := r.URL.Query().Get("s")

	q = q + "%"
	log.Print(q)

	db.DB.Where("first_name LIKE ?", q).Find(&actors)

	render.RenderList(w, r, m.NewActorListResponse(actors))
}

func CreateActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Request", "*")
	var data m.ActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	} else {
		render.Status(r, http.StatusAccepted)
	}

	actor := data.Actor
	log.Print(len(actor.FirstName))

	if len(actor.FirstName) == 0 || len(actor.LastName) == 0 {
		render.Status(r, http.StatusBadRequest)
	} else {

		e := db.DB.Create(actor)
		if e.Error == nil {

		} else {
			render.Status(r, http.StatusBadRequest)
		}
		log.Print("Error")
		log.Print(e.Error)
		render.Render(w, r, m.NewActorResponse(actor))
	}
}

func UpdateActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var data m.ActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	idParam := chi.URLParam(r, "id")
	actor := data.Actor
	actor.LastUpdate = time.Now()

	if db.DB.Model(&actor).Where("actor_id = ?", idParam).Updates(&actor).RowsAffected == 0 {
		db.DB.Create(&actor)
	}

	render.Status(r, http.StatusAccepted)
	render.Render(w, r, m.NewActorResponse(actor))
}

/**
 * Delete actor with the given id.
 * @param int - The id of actor.
 * @returns NewActorResponse
 */
func DeleteActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	idParam := chi.URLParam(r, "id")
	var actor *m.Actor

	e := db.DB.Delete(&actor, idParam)

	if e.Error == nil {
		render.Status(r, http.StatusAccepted)
	} else {
		render.Status(r, http.StatusForbidden)
	}
	log.Print(e.Error)

	render.Render(w, r, m.NewActorResponse(actor))
}

func ListFilmsByActorId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	actorIdParam := chi.URLParam(r, "id")

	var films []*m.Film
	var filmactor []*m.FilmActor

	db.DB.Where("actor_id = ?", actorIdParam).Find(&filmactor)
	log.Print(filmactor)

	for _, fa := range filmactor {
		var film *m.Film
		film_id := fa.FilmId
		db.DB.Where("film_id = ?", film_id).Find(&film)
		films = append(films, film)

	}
	render.RenderList(w, r, m.NewFilmListResponse(films))
}
