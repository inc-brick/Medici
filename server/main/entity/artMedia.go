package entity

import "time"

type ArtMedia struct {
	MediaId int `json:"media_id" db:"MEDIA_ID"`
	ArtistId int `json:"artist_id" db:"ARTIST_ID"`
	StartDate time.Time `json:"start_date" db:"START_DATE"`
	EndDate time.Time `json:"end_date" db:"END_DATE"`
	MediaType int `json:"media_type" db:"MEDIA_TYPE"`
	MediaNameJa string `json:"media_name_ja" db:"MEDIA_NAME_JA"`
	MediaNameEn string `json:"media_name_en" db:"MEDIA_NAME_EN"`
	MediaDescriptionJa string `json:"media_description_ja" db:"MEDIA_DESCRIPTION_JA"`
	MediaDescriptionEn string `json:"media_description_en" db:"MEDIA_DESCRIPTION_EN"`
	MediaUrl int `json:"media_url" db:"MEDIA_URL"`
}