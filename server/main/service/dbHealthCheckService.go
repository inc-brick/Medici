package service

import (
	"github.com/inc-brick/Medici/server/main/entity"
	"github.com/inc-brick/Medici/server/main/repository"
	"github.com/inc-brick/Medici/server/main/response"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

type DbHealthCheckService struct {
	Meta Service
}

func InitDbHealthCheckService(service Service) DbHealthCheckService {
	return DbHealthCheckService{Meta:service}
}

func (s DbHealthCheckService) HealthCheck(c echo.Context) error {
	artistEventRepository := repository.ArtistEventRepository{Db:s.Meta.DataAccess.Master}
	log.Info("login is called")
	artistEvent := entity.ArtistEvent{}

	maxEventId, err := artistEventRepository.SelectMaxEventId()
	if err != nil {
		log.Errorf("Failed to execute query. %s", err.Error())
	}
	artistEvent = entity.ArtistEvent{
		EventId:            maxEventId + 1,
		ArtistId:           1,
		StartDate:          time.Now(),
		EndDate:            time.Now().Add(time.Hour),
		EventNameJA:        "",
		EventNameEN:        "",
		EventDescriptionJA: "",
		EventDescriptionEN: "",
		EventImagePath:     "",
		EventStartAt:       time.Now(),
		EventEndAt:         time.Now().Add(time.Hour),
		EventPhone:         0,
		EventMailAddress:   "",
		EventUrl:           "",
	}

	artistEventRepository.Insert(artistEvent)

	if artistEvent, err = artistEventRepository.SelectByEventId(artistEvent.EventId); err != nil {
		log.Errorf("SelectByEventId is denied. %s", err.Error())
		return c.JSON(http.StatusOK, response.Response{Code: "errors.fetch.data", Data: nil})
	}
	return c.JSON(http.StatusOK, response.Response{Code: "", Data: artistEvent})
}

