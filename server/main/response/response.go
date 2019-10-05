package response

type FetchArtistInfo struct {
	ArtistId int `json:"artistId"`
	ArtistNameJPN string `json:"artistNameJPN"`
	ArtistNameENG string `json:"artistNameENG"`
	Description string `json:"description"`
	Works []Work `json:"works"`
	Events []Events `json:"events"`
}

type Work struct {
	Id int `json:"id"`
	Name string `json:"name"`
	CreatedYear int `json:"createdYear"`
	Height int `json:"height"`
	Width int `json:"width"`
	Material string `json:"material"`
	Url string `json:"url"`
}

type Events struct {
	Id int `json:"id"`
	Url string `json:"url"`
}