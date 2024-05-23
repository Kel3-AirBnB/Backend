package main

import (
	configs "airbnb/app/configs"
	databases "airbnb/app/database"
	"airbnb/app/middlewares"
	"airbnb/app/migrations"
	"airbnb/app/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := configs.InitConfig()
	dbMysql := databases.InitDBMysql(cfg)
	migrations.InitMigrations(dbMysql)
	s3Client, s3Bucket := databases.InitS3(cfg)

	e := echo.New()

	e.Use(middlewares.RemoveTrailingSlash)
	e.Use(middleware.CORSWithConfig(middlewares.CORSConfig()))

	routers.InitRouter(e, dbMysql, s3Client, cfg, s3Bucket)
	e.Logger.Fatal(e.Start(":8080"))
}
