package handler

import (
	"airbnb/app/middlewares"
	"airbnb/features/booking"
	"airbnb/features/user"
	"airbnb/utils/helper"
	"airbnb/utils/responses"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookingHandler struct {
	bookingService booking.ServiceInterface
	helperService  helper.HelperInterface
	userService    user.ServiceInterface
}

func New(bs booking.ServiceInterface, hs helper.HelperInterface, us user.ServiceInterface) *BookingHandler {
	return &BookingHandler{
		bookingService: bs,
		helperService:  hs,
		userService:    us,
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
	newBooking.StatusPembayaran = "Belum Dibayar"

	homeData, errhomeData := h.bookingService.GetHomeById(newBooking.PenginapanID)
	if errhomeData != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get homeData data", nil))
	}

	errInsert := h.bookingService.Create(GormToCore(newBooking), newBooking.CheckIn, newBooking.CheckOut, homeData.Harga)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error add data", errInsert))
	}
	return c.JSON(http.StatusCreated, responses.JSONWebResponse("success add data", nil))
}

func (h *BookingHandler) BookById(c echo.Context) error {
	updatePayment := PaymentRequest{}
	errBind := c.Bind(&updatePayment)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind"+errBind.Error(), nil))
	}

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

	inputCore := RequestToCore(updatePayment)

	_, errInsert := h.bookingService.Payment(idConv, idToken, inputCore, bookingData.CheckIn, bookingData.CheckOut, homeData.Harga)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error update data", errInsert))
	}

	PaymentResponse := PaymentResponse(*bookingData, *homeData)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success pay booking", PaymentResponse))
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

func (h *BookingHandler) GetInvoiceById(c echo.Context) error {
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

	homeData, errhomeData := h.bookingService.GetHomeById(bookingData.PenginapanID)
	if errhomeData != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get homeData data", nil))
	}

	projectResponse := InvoiceResponse(*bookingData, *homeData)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success get detail project", projectResponse))
}

func (h *BookingHandler) GetAllHistoryUser(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
	log.Println("idtoken:", idToken)

	result, err := h.bookingService.GetAll(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error read data", nil))
	}

	var allHistory []BookingResponse
	for _, value := range result {
		homeData, errhomeData := h.bookingService.GetHomeById(value.PenginapanID)

		if errhomeData != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get homeData data", nil))
		}
		allHistory = append(allHistory, InvoiceResponse(value, *homeData))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success read data", allHistory))
}

func (h *BookingHandler) GetHistoryHost(c echo.Context) error {
	fmt.Println("[Handler] GetHistoryHost")
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get project id", idConv))
	}
	fmt.Println("[Handler] ID House", idConv)

	idToken := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
	log.Println("idtoken:", idToken)

	result, err := h.bookingService.GetBookingByHomestay(uint(idConv), uint(idToken)) //rumah - token
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error read data", nil))
	}

	var allHistoryHomeById []HistoryHomeStayResponse
	for _, v := range result {
		fmt.Println("[Handler] Range Result", len(result))
		fmt.Println("[Handler] allHistoryHomeById", v.PenginapanID)
		fmt.Println("[Handler] userID", v.UserID)
		homeData, errhomeData := h.bookingService.GetHomeById(v.PenginapanID)
		fmt.Println("[Handler] homeData", homeData)
		if errhomeData != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get homeData data", nil))
		}
		userData, errUserData := h.userService.GetProfile(v.UserID)

		if errUserData != nil {
			return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get user data", nil))
		}
		allHistoryHomeById = append(allHistoryHomeById, HistoryHomeStayResponses(v, *homeData, *userData))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success read data", allHistoryHomeById))
}
