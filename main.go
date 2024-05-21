package main

import (
	configs "airbnb/app/configs"
	databases "airbnb/app/database"
	"airbnb/app/migrations"
	"airbnb/app/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.InitConfig()
	dbMysql := databases.InitDBMysql(cfg)
	migrations.InitMigrations(dbMysql)
	s3Client, _ := databases.InitS3(cfg)
	e := echo.New()
	routers.InitRouter(e, dbMysql, s3Client, cfg)
	e.Logger.Fatal(e.Start(":8080"))
}
