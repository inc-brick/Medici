package request

type contactsParam struct {
	ArtistId int `json:"artist_id"`
	WorkId int `json:"work_id"`
	Method int `json:"method"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}