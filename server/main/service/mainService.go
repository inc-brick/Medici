package service

import (
	"github.com/inc-brick/Medici/server/main/db"
)

type Service struct {
	DataAccess *db.Db
}

func InitService(env string) Service {
	return Service{DataAccess: db.Init(env)}
}

func (s Service) DbHealthCheckService() DbHealthCheckService {
	return InitDbHealthCheckService(s)
}

func (s Service) ArtistService() ArtistService {
	return InitArtistService(s)
}
