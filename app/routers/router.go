package routers

import (
	"airbnb/app/configs"
	"airbnb/features/review/data"
	"airbnb/features/review/handler"
	"airbnb/features/review/service"
	_userData "airbnb/features/user/data"
	_userHandler "airbnb/features/user/handler"
	_userService "airbnb/features/user/service"
	"airbnb/utils/encrypts"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB, s3 *s3.S3, cfg *configs.AppConfig) {
	hashService := encrypts.NewHashService()
	userData := _userData.New(db, s3)

	userService := _userService.New(userData, hashService, s3, cfg)

	userHandlerAPI := _userHandler.New(userService, hashService)

	//review
	reviewData := data.New(db)
	reviewService := service.New(reviewData)
	reviewHandlerAPI := handler.New(reviewService)

	e.POST("/users", userHandlerAPI.Register)

	//review
	e.GET("/reviews", reviewHandlerAPI.GetAll)

}
