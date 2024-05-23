package routers

import (
	"airbnb/app/configs"

	"airbnb/app/middlewares"
	"airbnb/features/review/data"
	"airbnb/features/review/handler"
	"airbnb/features/review/service"

	_userData "airbnb/features/user/data"
	_userHandler "airbnb/features/user/handler"
	_userService "airbnb/features/user/service"

	_bookingData "airbnb/features/booking/data"
	_bookingHandler "airbnb/features/booking/handler"
	_bookingService "airbnb/features/booking/service"
	"airbnb/utils/encrypts"
	"airbnb/utils/helper"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB, s3 *s3.S3, cfg *configs.AppConfig, s3Bucket string) {
	hashService := encrypts.NewHashService()
	helperService := helper.NewHelperService()
	userData := _userData.New(db, s3)
	userService := _userService.New(userData, hashService, s3, s3Bucket)
	userHandlerAPI := _userHandler.New(userService, hashService)

	bookingData := _bookingData.New(db)
	bookingService := _bookingService.New(bookingData, userData)
	bookingHanlderAPI := _bookingHandler.New(bookingService, helperService)

	//review
	reviewData := data.New(db)
	reviewService := service.New(reviewData, s3, s3Bucket)
	reviewHandlerAPI := handler.New(reviewService)

	e.POST("/users", userHandlerAPI.Register)
	e.POST("/login", userHandlerAPI.Login)
	e.GET("/profile", userHandlerAPI.Profile, middlewares.JWTMiddleware())
	e.GET("/users/:id", userHandlerAPI.GetById)
	e.PUT("/users", userHandlerAPI.UpdateUserById)
	e.DELETE("/users", userHandlerAPI.Delete)

	e.POST("/booking/:id", bookingHanlderAPI.Create)
	e.GET("/booking/:id", bookingHanlderAPI.GetBookById)
	e.GET("/bookings/:id", bookingHanlderAPI.BookById)
	//review
	e.GET("/reviews", reviewHandlerAPI.GetAll)
	e.POST("/reviews", reviewHandlerAPI.CreateReview)
	e.GET("/reviews/:id", reviewHandlerAPI.GetById)

}
