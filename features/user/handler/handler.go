package handler

import (
	"airbnb/app/middlewares"
	"airbnb/features/user"
	"airbnb/utils/encrypts"
	"airbnb/utils/responses"
	"log"
	"net/http"
	"strconv"
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
	_, errInsert := uh.userService.Create(inputCore, file, handler.Filename)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error add data", errInsert))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error add data", errInsert))
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

func (uh *UserHandler) Profile(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
	log.Println("idtoken:", idToken)
	userData, err := uh.userService.GetProfile(uint(idToken)) // Ambil data pengguna dari Redis
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get user data", nil))
	}
	userResponse := CoreToGorm(*userData)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success get profile", userResponse))
}

func (uh *UserHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get user id", idConv))
	}

	userData, err := uh.userService.GetProfile(uint(idConv)) // Ambil data pengguna dari Redis
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get user data", nil))
	}
	userResponse := CoreToGorm(*userData)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success get profile", userResponse))
}

func (uh *UserHandler) UpdateUserById(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	updatedUser := UserRequest{}
	errBind := c.Bind(&updatedUser)
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
	inputCore := RequestToCore(updatedUser)
	//uh.userService.Create(inputCore, file, handler.Filename)
	//inputNewCore := RequestToCore(updatedUser) // mapping  dari request ke core
	_, errUpdate := uh.userService.UpdateById(uint(idToken), inputCore, file, handler.Filename)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error update data", errUpdate))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success update data", errUpdate))
}

func (uh *UserHandler) Delete(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	err := uh.userService.Delete(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error delete data", err))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success delete data", err))
}
