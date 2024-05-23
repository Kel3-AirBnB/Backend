package handler

import (
	"airbnb/app/middlewares"
	homestay "airbnb/features/homestay"
	"airbnb/utils/responses"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type HomestayHandler struct {
	homestayService homestay.ServiceInterface
}

func NewHomestayHandler(hs homestay.ServiceInterface) *HomestayHandler {
	return &HomestayHandler{
		homestayService: hs,
	}
}

func (h *HomestayHandler) GetAll(c echo.Context) error {
	result, err := h.homestayService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error read data",
		})
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success read data", result))
}

func (h *HomestayHandler) CreateHomestay(c echo.Context) error {
	newHomestay := HomeStayRequest{}
	errBind := c.Bind(&newHomestay)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind "+errBind.Error(), nil))
	}

	file, handler, err := c.Request().FormFile("homestay_profile")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Unable to upload photo: " + err.Error(),
		})
	}
	defer file.Close()
	inputCore := RequestToCore(newHomestay)
	_, errInsert := h.homestayService.Create(inputCore, file, handler.Filename)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error add data", errInsert))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error add data", errInsert))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("success add data", nil))
}

func (h *HomestayHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get homestay id", nil))
	}

	homestayData, err := h.homestayService.GethomeStayid(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get homestay data", nil))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success get homestay", homestayData))
}

func (h *HomestayHandler) Delete(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	err := h.homestayService.Delete(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error delete data", err))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success delete data", nil))
}
