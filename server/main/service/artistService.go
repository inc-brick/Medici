package service

import (
	"github.com/inc-brick/Medici/server/main/response"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

type ArtistService struct {
	Meta Service
}

func InitArtistService(service Service) ArtistService {
	return ArtistService{Meta:service}
}

func (s ArtistService) GetArtistInfo(c echo.Context) error {
	log.Info("fetchArtistInfo is called")
	artistId, err := strconv.Atoi(c.Param("artistId"))
	if err != nil {
		log.Fatal("Could not parse artistId. artistId: %s ¥nCause is... %s", artistId, err.Error())
		return c.JSON(http.StatusOK, response.Response{
			Code: "errors.data.invalid",
			Data: nil,
		})
	}
	var res response.GetArtistInfoResponse
	res.ArtistId = artistId

	var works = make([]response.Work, 0)
	for i := 1; i < artistId + 1; i++ {
		works = append(works, response.Work{
			Id:          i,
			Name:        "WorkName_" + strconv.Itoa(i),
			CreatedYear: 2020,
			Height:      1000,
			Width:       800,
			Material:    "アクリル",
			Url:         "https://google.com",
		})
	}

	var events = make([]response.Event, 0)
	for i := 1; i < artistId + 1; i++ {
		events = append(events, response.Event{
			Id:  i,
			Url: "https://google.com",
		})
	}

	var medias = make([]response.Media, 0)
	for i := 1; i < artistId + 1; i++ {
		medias = append(medias, response.Media{
			Id:  i,
			Url: "https://google.com",
		})
	}

	artistInfo := response.GetArtistInfoResponse{
		ArtistId:      artistId,
		ArtistNameJPN: "ArtistNameJPN_" + strconv.Itoa(artistId),
		ArtistNameENG: "ArtistNameENG_" + strconv.Itoa(artistId),
		Description:   "Description_" + strconv.Itoa(artistId),
		Works:         works,
		Events:        events,
		Medias:        medias,
	}
	// ======== 処理内容 ========
	// ComArtistMstから情報を取得
	// ArtWorkMstから情報を取得
	// ArtArtistEventTrnから情報を取得
	// MediaTrnから情報を取得
	// response作って返す

	//if err := master.Get(&user, "SELECT * FROM `users` WHERE `username` = ?", req.UserName); err != nil {
	//	log.Errorf("Failed to execute query. %s", err.Error())
	//	return c.JSON(http.StatusUnauthorized, "Can not find user")
	//}
	return c.JSON(http.StatusOK, response.Response{
		Code: "",
		Data: artistInfo,
	})
}

