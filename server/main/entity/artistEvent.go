package entity

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"time"
)

type ArtistEvent struct {
	EventId int `json:"event_id" db:"event_id"`
	ArtistId int `json:"artist_id" db:"artist_id"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate time.Time `json:"end_date" db:"end_date"`
	EventNameJA string `json:"event_name_ja" db:"event_name_ja"`
	EventNameEN string `json:"event_name_en" db:"event_name_en"`
	EventDescriptionJA string `json:"event_description_ja" db:"event_description_ja"`
	EventDescriptionEN string `json:"event_description_en" db:"event_description_en"`
	EventImagePath string `json:"event_image_path" db:"event_image_path"`
	EventStartAt time.Time `json:"event_start_at" db:"event_start_at"`
	EventEndAt time.Time `json:"event_end_at" db:"event_end_at"`
	EventPhone int `json:"event_phone" db:"event_phone"`
	EventMailAddress string `json:"event_mail_address" db:"event_mail_address"`
	EventUrl string `json:"event_url" db:"event_url"`
}

//func (e ArtistEvent) Init(evenId int,
//							artistId int,
//							startDate time.Time,
//							endDate time.Time,
//							eventNameJA string,
//							eventNameEN string,
//							eventDescriptionJA string,
//							eventDescriptionEN string,
//							eventImagePath string,
//							eventStartAt time.Time,
//							eventEndAt time.Time,
//							eventPhone int,
//							eventMailAddress string,
//							eventUrl string) {
//
//}

func (e ArtistEvent) Insert(tx *sql.Tx) {
	_, err := tx.Exec("INSERT INTO `artist_event_mast` VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
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
}

func (e ArtistEvent) SelectByEventId(db *sqlx.DB, eventId int) (ArtistEvent, error) {
	err := db.Get(&e, "SELECT * FROM `artist_event_mast` WHERE `event_id` = ?", eventId)
	if err != nil {
		log.Errorf("Failed to execute query. %s", err.Error())
	}
	return e, err
}

func (e ArtistEvent) SelectMaxEventId(db *sqlx.DB) (int, error) {
	var maxEventId int
	err := db.Get(&maxEventId, "SELECT max(`event_id`) FROM `artist_event_mast`")
	if err != nil {
		log.Errorf("Failed to execute query. %s", err.Error())
	}
	return maxEventId, err
}