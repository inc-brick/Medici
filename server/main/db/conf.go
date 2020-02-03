package db

import "github.com/inc-brick/Medici/server/main/constant"

type DataBaseConfig struct {
	Host string
	Port string
	UserName string
	Password string
}

func ConfigInit(env string) DataBaseConfig {
	return DataBaseConfig{
		Host:     constant.Host[env],
		Port:     constant.Port[env],
		UserName: constant.UserName[env],
		Password: constant.Password[env],
	}
}