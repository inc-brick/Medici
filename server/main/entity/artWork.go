package entity

import "time"

type ArtWork struct {
	WorkId int `json:"work_id" db:"WORK_ID"`
	ArtistId int `json:"artist_id" db:"ARTIST_ID"`
	StartDate time.Time `json:"start_date" db:"START_DATE"`
	EndDate time.Time `json:"end_date" db:"END_DATE"`
	WorkTitleJa string `json:"work_title_ja" db:"WORK_TITLE_JA"`
	WorkTitleEn string `json:"work_title_en" db:"WORK_TITLE_EN"`
	WorkDescriptionJa string `json:"work_description_ja" db:"WORK_DESCRIPTION_JA"`
	WorkDescriptionEn string `json:"work_description_en" db:"WORK_DESCRIPTION_EN"`
	WorkImagePath string `json:"work_image_path" db:"WORK_IMAGE_PATH"`
	CreateYear int `json:"create_year" db:"CREATE_YEAR"`
	Size string `json:"size" db:"SIZE"`
	SizeXPosition int `json:"size_x_position" db:"SIZE_X_POSITION"`
	SizeYPosition int `json:"size_y_position" db:"SIZE_Y_POSITION"`
	SizeZPosition int `json:"size_z_position" db:"SIZE_Z_POSITION"`
	Material string `json:"material" db:"MATERIAL"`
}

