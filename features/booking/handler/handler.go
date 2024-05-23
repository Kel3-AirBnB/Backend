package handler

import (
	"airbnb/app/middlewares"
	"airbnb/features/booking"
	"airbnb/utils/helper"
	"airbnb/utils/responses"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookingHandler struct {
	bookingService booking.ServiceInterface
	helperService  helper.HelperInterface
}

func New(bs booking.ServiceInterface, hs helper.HelperInterface) *BookingHandler {
	return &BookingHandler{
		bookingService: bs,
		helperService:  hs,
	}
}

func (h *BookingHandler) Create(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get user id", idConv))
	}

	idToken := middlewares.ExtractTokenUserId(c)
	newBooking := BookingRequest{}
	errBind := c.Bind(&newBooking)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind data: "+errBind.Error(), nil))
	}
	newBooking.UserID = uint(idToken)
	newBooking.PenginapanID = uint(idConv)
	// newBooking.TotalTransaksi =
	newBooking.StatusPembayaran = "Belum Dibayar"
	errInsert := h.bookingService.Create(GormToCore(newBooking))
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error add data", errInsert))
	}
	return c.JSON(http.StatusCreated, responses.JSONWebResponse("success add data", nil))
}

func (h *BookingHandler) BookById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get project id", idConv))
	}

	idToken := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
	log.Println("idtoken:", idToken)

	bookingData, err := h.bookingService.GetBookingById(uint(idConv), uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get bookingData data", nil))
	}

	homeData, errhomeData := h.bookingService.GetHomeById(bookingData.PenginapanID)
	if errhomeData != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get homeData data", nil))
	}

	result, err := h.bookingService.Payment(bookingData, homeData, idConv, idToken)

	projectResponse := BookingResponses(*bookingData, *homeData)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success get detail project", projectResponse))
}

func (h *BookingHandler) GetBookById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get project id", idConv))
	}

	idToken := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
	log.Println("idtoken:", idToken)

	bookingData, err := h.bookingService.GetBookingById(uint(idConv), uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get project data", nil))
	}

	projectResponse := SelectResponses(*bookingData)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success get detail project", projectResponse))
}
