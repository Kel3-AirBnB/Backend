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

	_homestayData "airbnb/features/homestay/data"
	_homestayHandler "airbnb/features/homestay/handler"
	_homestayService "airbnb/features/homestay/service"

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
	bookingService := _bookingService.New(bookingData, userData, helperService)
	bookingHanlderAPI := _bookingHandler.New(bookingService, helperService)

	homestayData := _homestayData.New(db)
	homestayService := _homestayService.NewHomestayService(homestayData, s3, s3Bucket)
	homestayHandlerAPI := _homestayHandler.NewHomestayHandler(homestayService)

	//review
	reviewData := data.New(db)
	reviewService := service.New(reviewData, s3, userData, s3Bucket, cfg.VALIDATLOCALORSERVER)
	reviewHandlerAPI := handler.New(reviewService)

	e.POST("/users", userHandlerAPI.Register)
	e.POST("/login", userHandlerAPI.Login)
	e.GET("/profile", userHandlerAPI.Profile, middlewares.JWTMiddleware())
	e.GET("/users/:id", userHandlerAPI.GetById)
	e.PUT("/users", userHandlerAPI.UpdateUserById)
	e.DELETE("/users", userHandlerAPI.Delete)

	e.GET("/homestay", homestayHandlerAPI.GetAll)
	e.POST("/homestay", homestayHandlerAPI.CreateHomestay)
	e.GET("/homestay/:id", homestayHandlerAPI.GetById)
	e.GET("/homestay", homestayHandlerAPI.Delete)

	e.POST("/booking/:id", bookingHanlderAPI.Create)           //membuat pesanan
	e.GET("/booking/:id", bookingHanlderAPI.GetBookById)       //cek status pesanan dengan id pesanan dan user / jwt
	e.POST("/payment/:id", bookingHanlderAPI.BookById)         //melakukan pembayaran
	e.GET("/payment/:id", bookingHanlderAPI.GetBookById)       //Mendapatkan invoices
	e.GET("/historyuser", bookingHanlderAPI.GetAllHistoryUser) //Mendapatkan hitory user

	//review
	e.GET("/reviews", reviewHandlerAPI.GetAll)
	e.POST("/reviews", reviewHandlerAPI.CreateReview)
	e.GET("/reviews/:id", reviewHandlerAPI.GetById)
	e.DELETE("/reviews/:id", reviewHandlerAPI.Delete)
	e.PUT("/reviews/:id", reviewHandlerAPI.UpdateReview)
	e.GET("/reviewsuserid", reviewHandlerAPI.GetReviewsByUserID, middlewares.JWTMiddleware())
	e.GET("/reviewshomestay/:penginapanID", reviewHandlerAPI.GetReviewByPenginapanID)
}
