package configs

import (
	"os"
	"strconv"
)

// var (
// 	JWT_SECRET string
// 	ctx        = context.Background()
// )

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     int
	DB_NAME     string

	RD_ADDRESS  string
	RD_PASSWORD string
	RD_DATABASE int
}

func ReadEnv() *AppConfig {
	var app = AppConfig{}
	app.DB_USERNAME = os.Getenv("DBUSER")
	app.DB_PASSWORD = os.Getenv("DBPASS")
	app.DB_HOSTNAME = os.Getenv("DBHOST")
	portConv, errConv := strconv.Atoi(os.Getenv("DBPORT"))
	if errConv != nil {
		panic("error convert dbport")
	}
	app.DB_PORT = portConv
	app.DB_NAME = os.Getenv("DBNAME")
	//JWT_SECRET = os.Getenv("JWTSECRET")

	//redis
	app.RD_ADDRESS = os.Getenv("RDADDR")
	app.RD_PASSWORD = os.Getenv("RDPASS")
	dbConv, errConv := strconv.Atoi(os.Getenv("RDDB"))
	if errConv != nil {
		panic("error convert redis database")
	}
	app.RD_DATABASE = dbConv

	return &app
}

func InitConfig() *AppConfig {
	return ReadEnv()
}