package repository

import (
	"github.com/inc-brick/Medici/server/main/entity"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type ArtistEventRepository struct {
	Db *sqlx.DB
}

func (r ArtistEventRepository) Insert(e entity.ArtistEvent) {
	tx, err := r.Db.Begin()
	if err != nil {
		log.Errorf("Failed to begin transaction. %s", err.Error())
	}
	log.Info("transaction begin: OK")
	_, err = tx.Exec("INSERT INTO `artist_event_mast` VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		e.EventId,
		e.ArtistId,
		e.StartDate,
		e.EndDate,
		e.EventNameJA,
		e.EventNameEN,
		e.EventDescriptionJA,
		e.EventDescriptionEN,
		e.EventImagePath,
		e.EventStartAt,
		e.EventEndAt,
		e.EventPhone,
		e.EventMailAddress,
		e.EventUrl)
	if err != nil {
		log.Errorf("Failed to execute query. %s", err.Error())
	}
	err = tx.Commit()
	if err != nil {
		log.Errorf("Failed to commit. %s", err.Error())
	}
}

func (r ArtistEventRepository) SelectByEventId(eventId int) (entity.ArtistEvent, error) {
	artistEvent := entity.ArtistEvent{}
	err := r.Db.Get(&artistEvent, "SELECT * FROM `artist_event_mast` WHERE `event_id` = ?", eventId)
	if err != nil {
		log.Errorf("Failed to execute query. %s", err.Error())
	}
	return artistEvent, err
}

func (r ArtistEventRepository) SelectMaxEventId() (int, error) {
	var maxEventId int
	err := r.Db.Get(&maxEventId, "SELECT max(`event_id`) FROM `artist_event_mast`")
	if err != nil {
		log.Errorf("Failed to execute query. %s", err.Error())
	}
	return maxEventId, err
}