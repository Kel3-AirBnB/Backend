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

func (uh *UserHandler) Register(c echo.Context) error {
	log.Print("[Handler Layer]")

	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind"+errBind.Error(), nil))
	}

	file, handler, err := c.Request().FormFile("profilepicture")
	if err != nil {
		log.Print("[Handler Layer] Error getting file from request:", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Unable to upload photo: " + err.Error(),
		})
	}
	defer file.Close()
	inputCore := RequestToCore(newUser)
	uh.userService.Create(inputCore, file, handler.Filename)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error add data", err))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error add data", err))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("success add data", nil))
}
