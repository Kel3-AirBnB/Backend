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

	// rdb, ctx, ttl, err := databases.InitRedis(cfg)
	// if err != nil {
	// 	log.Fatalf("failed to initialize redis: %v", err)
	// }

	e := echo.New()
	routers.InitRouter(e, dbMysql)
	e.Logger.Fatal(e.Start(":8080"))
}
