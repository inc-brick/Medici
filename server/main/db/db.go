package db

import (
	"fmt"
	"github.com/inc-brick/Medici/server/main/constant"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type Db struct {
	 Master *sqlx.DB
	 User *sqlx.DB
}

func Init(env string) *Db {
	conf := ConfigInit(env)
	dsMaster := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", conf.UserName, conf.Password, conf.Host, conf.Port, constant.DB_MASTER)
	dsUser := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", conf.UserName, conf.Password, conf.Host, conf.Port, constant.DB_USER)
	master, err := sqlx.Open("mysql", dsMaster)
	if err != nil {
		log.Fatalf("failed to connect to DB: %s.", err.Error())
	}
	user, err := sqlx.Open("mysql", dsUser)
	if err != nil {
		log.Fatalf("failed to connect to DB: %s.", err.Error())
	}
	return &Db{
		Master: master,
		User:   user,
	}
}