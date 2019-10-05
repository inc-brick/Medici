package entity

import "time"

type ComArtistPrivateInfo struct {
	ArtistPrivateId int `json:"artist_private_id" db:"ARTIST_PRIVATE_ID"`
	ArtistId int `json:"artist_id" db:"ARTIST_ID"`
	PhoneNumber int `json:"phone_number" db:"PHONE_NUMBER"`
	MailAddress string `json:"mail_address" db:"MAIL_ADDRESS"`
	PostalCode string `json:"postal_code" db:"POSTAL_CODE"`
	Address string `json:"address" db:"ADDRESS"`
	WorkPlacePostalCode string `json:"work_place_postal_code" db:"WORK_PLACE_POSTAL_CODE"`
	WorkPlaceAddress string `json:"work_place_address" db:"WORK_PLACE_ADDRESS"`
	FormalName string `json:"formal_name" db:"FORMAL_NAME"`
}

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

type ArtArtistEvent struct {
	EventId int `json:"event_id" db:"EVENT_ID"`
	ArtistId int `json:"artist_id" db:"ARTIST_ID"`
	StartDate time.Time `json:"start_date" db:"START_DATE"`
	EndDate time.Time `json:"end_date" db:"END_DATE"`
	EventNameJA string `json:"event_name_ja" db:"EVENT_NAME_JA"`
	EventNameEN string `json:"event_name_en" db:"EVENT_NAME_EN"`
	EventDescriptionJA string `json:"event_description_ja" db:"EVENT_DESCRIPTION_JA"`
	EventDescriptionEN string `json:"event_description_en" db:"EVENT_DESCRIPTION_EN"`
	EventImagePath string `json:"event_image_path" db:"EVENT_IMAGE_PATH"`
	EventStartAt time.Time `json:"event_start_at" db:"EVENT_START_AT"`
	EventEndAt time.Time `json:"event_end_at" db:"EVENT_END_AT"`
	EventPhone int `json:"event_phone" db:"EVENT_PHONE"`
	EventMailAddress string `json:"event_mail_address" db:"EVENT_MAIL_ADDRESS"`
	EventUrl string `json:"event_url" db:"EVENT_URL"`
}

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
