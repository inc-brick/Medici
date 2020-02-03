package entity

import "time"

type ArtRelatedItem struct {
	ItemId int `json:"item_id" db:"ITEM_ID"`
	ArtistId int `json:"artist_id" db:"ARTIST_ID"`
	StartDate time.Time `json:"start_date" db:"START_DATE"`
	EndDate time.Time `json:"end_date" db:"END_DATE"`
	ItemTitleJa string `json:"item_title_ja" db:"ITEM_TITLE_JA"`
	ItemTitleEn string `json:"item_title_en" db:"ITEM_TITLE_EN"`
	ItemDescriptionJa string `json:"item_description_ja" db:"ITEM_DESCRIPTION_JA"`
	ItemDescriptionEn string `json:"item_description_en" db:"ITEM_DESCRIPTION_EN"`
	ItemImagePath string `json:"item_image_path" db:"ITEM_IMAGE_PATH"`
	CreateYear int `json:"create_year" db:"CREATE_YEAR"`
	Size string `json:"size" db:"SIZE"`
	SizeXPosition int `json:"size_x_position" db:"SIZE_X_POSITION"`
	SizeYPosition int `json:"size_y_position" db:"SIZE_Y_POSITION"`
	SizeZPosition int `json:"size_z_position" db:"SIZE_Z_POSITION"`
	Material string `json:"material" db:"MATERIAL"`
}