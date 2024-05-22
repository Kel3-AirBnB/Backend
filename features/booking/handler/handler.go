package handler

import (
	"airbnb/app/middlewares"
	"airbnb/features/booking"
	"airbnb/utils/responses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookingHandler struct {
	bookingService booking.ServiceInterface
}

func New(bs booking.ServiceInterface) *BookingHandler {
	return &BookingHandler{
		bookingService: bs,
	}
}

func (h *BookingHandler) Create(c echo.Context) error {
	fmt.Println("--- [Handler Layer] ---")
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get user id", idConv))
	}
	fmt.Println("[Handler Layer] idparam:", idConv)

	idToken := middlewares.ExtractTokenUserId(c)
	fmt.Println("[Handler Layer] idtoken:", idToken)

	newBooking := BookingRequest{}
	errBind := c.Bind(&newBooking)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind data: "+errBind.Error(), nil))
	}
	newBooking.UserID = uint(idToken)
	newBooking.PenginapanID = uint(idConv)
	errInsert := h.bookingService.Create(GormToCore(newBooking))
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error add data", errInsert))
	}
	return c.JSON(http.StatusCreated, responses.JSONWebResponse("success add data", nil))
}
