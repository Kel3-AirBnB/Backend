package routers

import (
	_userData "airbnb/features/user/data"
	_userHandler "airbnb/features/user/handler"
	_userService "airbnb/features/user/service"
	"airbnb/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	hashService := encrypts.NewHashService()
	userData := _userData.New(db)

	userService := _userService.New(userData, hashService)

	userHandlerAPI := _userHandler.New(userService, hashService)

	e.POST("/users", userHandlerAPI.Register)
}
