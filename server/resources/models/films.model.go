package models

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/render"
)

type Film struct {
	FilmId             int       `gorm:"type:smallint;primaryKey"`
	Title              string    `gorm:"type:varchar(45)"`
	Description        string    `gorm:"type:text;default:null"`
	ReleaseYear        int       `gorm:"type:year;default:2000"`
	OriginalLanguageId int       `gorm:"type:tinyint;default:null"`
	LanguageId         int       `gorm:"type:tinyint;unsigned;notNull;default:1"`
	Length             int       `gorm:"type:smallint;default:null"`
	Rating             string    `gorm:"type:enum('G','PG','PG-13','R','NC-17');default:'G'"`
	RentalRate         float64   `gorm:"type:decimal(4,2);default:1.99"`
	LastUpdate         time.Time `gorm:"autoCreateTime"`
}

func (Film) TableName() string {
	return "film"
}

type FilmRequest struct {
	*Film
}

func (a *FilmRequest) Bind(r *http.Request) error {
	if a.Film == nil {
		return errors.New("missing required Film fields")
	}

	a.Film.Title = strings.ToUpper(a.Film.Title)
	a.Film.Description = strings.ToUpper(a.Film.Description)

	return nil
}

type FilmResponse struct {
	*Film
}

func NewFilmResponse(film *Film) *FilmResponse {
	return &FilmResponse{film}
}

func NewFilmListResponse(films []*Film) []render.Renderer {
	list := []render.Renderer{}
	for _, film := range films {
		list = append(list, NewFilmResponse(film))
	}
	return list
}

func (a *FilmResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
