package entity

type ComArtistPrivateInfo struct {
	ArtistPrivateId int `json:"artist_private_id" db:"artist_private_id"`
	ArtistId int `json:"artist_id" db:"artist_id"`
	PhoneNumber int `json:"phone_number" db:"phone_number"`
	MailAddress string `json:"mail_address" db:"mail_address"`
	PostalCode string `json:"postal_code" db:"postal_code"`
	Address string `json:"address" db:"address"`
	WorkPlacePostalCode string `json:"work_place_postal_code" db:"work_place_postal_code"`
	WorkPlaceAddress string `json:"work_place_address" db:"work_place_address"`
	FormalName string `json:"formal_name" db:"formal_name"`
}

