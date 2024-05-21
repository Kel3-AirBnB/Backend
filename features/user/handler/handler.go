package handler

import (
	"airbnb/features/user"
	"airbnb/utils/encrypts"
	"airbnb/utils/responses"
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

	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind"+errBind.Error(), nil))
	}

	file, handler, err := c.Request().FormFile("profilepicture")
	if err != nil {
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

func (uh *UserHandler) Login(c echo.Context) error {
	var reqLoginData = LoginRequest{}
	errBind := c.Bind(&reqLoginData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind"+errBind.Error(), nil))
	}
	result, token, err := uh.userService.Login(reqLoginData.Email, reqLoginData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error login", result))
	}
	result.Token = token
	var resultResponse = ResponseLogin(result)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success login", resultResponse))
}
