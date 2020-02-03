package entity

import "time"

type ComArtist struct {
	ArtistId int `json:"artist_id" db:"ARTIST_ID"`
	Sequence int `json:"sequence" db:"SEQUENCE"`
	StartDate time.Time `json:"start_date" db:"START_DATE"`
	EndDate time.Time `json:"end_date" db:"END_DATE"`
	ArtistNameJa string `json:"artist_name_ja" db:"ARTIST_NAME_JA"`
	ArtistNameEn string `json:"artist_name_en" db:"ARTIST_NAME_EN"`
	DescriptionJa string `json:"description_ja" db:"DESCRIPTION_JA"`
	DescriptionEn string `json:"description_en" db:"DESCRIPTION_EN"`
}