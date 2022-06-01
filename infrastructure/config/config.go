package config

import (
	"log"

	"github.com/wilmerpino/mutant/infrastructure/utils"
	"github.com/wilmerpino/mutant/public/config"
)

type DatabaseConfig struct {
	Host,
	User,
	Name,
	Password,
	Port string
}

type EnviromentConfig struct {
	AppName,
	LevelLog,
	PortServer string
	Database DatabaseConfig
}

func (v *EnviromentConfig) InitVariables() {
	var err error
	v.AppName, err = utils.ValidString("APP_NAME", config.AppName)
	if err != nil {
		log.Fatal(err.Error())
	}
	v.LevelLog, err = utils.ValidString("LOG_LEVEL", config.LevelLog)
	if err != nil {
		log.Fatal(err.Error())
	}
	v.PortServer, _ = utils.ValidString("PORT_SERVER", config.PortServer)

	v.Database.Host, err = utils.ValidString("BD_HOST", config.DBHost)
	if err != nil {
		log.Fatal(err.Error())
	}
	v.Database.User, err = utils.ValidString("BD_USER", config.DBUser)
	if err != nil {
		log.Fatal(err.Error())
	}
	v.Database.Name, err = utils.ValidString("BD_NAME", config.DBName)
	if err != nil {
		log.Fatal(err.Error())
	}
	v.Database.Password, err = utils.ValidString("BD_PASSWORD", config.DBPassword)
	if err != nil {
		log.Fatal(err.Error())
	}
	v.Database.Port, err = utils.ValidString("BD_PORT", config.DBPort)
	if err != nil {
		log.Fatal(err.Error())
	}
}
