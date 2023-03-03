package models

type FilmActor struct {
	ActorId int `gorm:"type:smallint;primaryKey;index:,unique"`
	FilmId  int `gorm:"type:smallint;primaryKey;index:,unique"`
}

func (FilmActor) TableName() string {
	return "film_actor"
}

type FilmActorRequest struct {
	*FilmActor
}
