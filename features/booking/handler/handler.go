package handler

import (
	"airbnb/features/booking"
)

type BookingHandler struct {
	bookingService booking.ServiceInterface
}

func New(bs booking.ServiceInterface) *BookingHandler {
	return &BookingHandler{
		bookingService: bs,
	}
}

// func (h *BookingHandler) Create(c echo.Context) error {
// 	newBooking := BookingRequest{}
// 	errBind := c.Bind(&newBooking)
// 	if errBind != nil {
// 		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind data: "+errBind.Error(), nil))
// 	}
// 	idToken := middlewares.ExtractTokenUserId(c)
// 	newBooking.UserID = uint(idToken)
// 	errInsert := h.projectService.Create(GormToCore(newBooking))
// 	if errInsert != nil {
// 		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error add data", errInsert))
// 	}
// 	return c.JSON(http.StatusCreated, responses.JSONWebResponse("success add data", nil))
// }
