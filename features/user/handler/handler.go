package handler

import (
	"airbnb/features/user"
	"airbnb/utils/encrypts"
	"airbnb/utils/responses"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.ServiceInterface
	hashService encrypts.HashInterface
}

func New(us user.ServiceInterface, hash encrypts.HashInterface) *UserHandler {
	return &UserHandler{
		userService: us,
		hashService: hash,
	}
}

// func (uh *UserHandler) Register(c echo.Context) error {
// 	log.Print("[Handler Layer]")
// 	newUser := UserRequest{} // membaca data dari request body
// 	errBind := c.Bind(&newUser)
// 	if errBind != nil {
// 		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind"+errBind.Error(), nil))
// 	}

// 	file, handler, err := c.Request().FormFile("profilepicture")
// 	log.Print("[Handler Layer] file", file)
// 	log.Print("[Handler Layer] handler", handler)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "Unable to upload photo: " + err.Error(),
// 		})
// 	}
// 	defer file.Close()

// 	inputCore := RequestToCore(newUser) // mapping  dari request ke core
// 	log.Print("[Handler Layer] inputCore", inputCore)

// 	photoURL, err := uh.userService.Create(inputCore, file, handler.Filename)
// 	//errInsert := uh.userService.Create(inputCore) // memanggil/mengirimkan data ke method service layer
// 	log.Print("[Handler Layer] photoURL", photoURL)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "validation") {
// 			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error add data", err))
// 		}
// 		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error add data", err))
// 	}
// 	return c.JSON(http.StatusCreated, responses.JSONWebResponse("success add data", photoURL))
// }

func (uh *UserHandler) Register(c echo.Context) error {
	log.Print("[Handler Layer]")

	// Bind request body to UserRequest
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind"+errBind.Error(), nil))
	}

	// Get file from request
	file, handler, err := c.Request().FormFile("profilepicture")
	if err != nil {
		log.Print("[Handler Layer] Error getting file from request:", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Unable to upload photo: " + err.Error(),
		})
	}
	defer file.Close()

	log.Print("[Handler Layer] File:", file)
	log.Print("[Handler Layer] Handler:", handler)

	// Convert UserRequest to Core
	inputCore := RequestToCore(newUser)
	log.Print("[Handler Layer] Input Core:", inputCore)

	// Call userService.Create
	photoURL, err := uh.userService.Create(inputCore, file, handler.Filename)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error add data", err))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error add data", err))
	}

	log.Print("[Handler Layer] Photo URL:", photoURL)

	// Return response
	return c.JSON(http.StatusCreated, responses.JSONWebResponse("success add data", photoURL))
}
