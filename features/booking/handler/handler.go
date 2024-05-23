package handler

import (
	"airbnb/app/middlewares"
	"airbnb/features/booking"
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
	newBooking.StatusPembayaran = "Belum Dibayar"
	errInsert := h.bookingService.Create(GormToCore(newBooking))
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

	harga := bookingData.TotalTransaksi
	fmt.Println("[Handler Layer] harga: ", harga)
	projectResponse := InvoiceResponse(*bookingData, *homeData)
	fmt.Println("[Handler Layer] Total Transaksi: ", projectResponse.TotalTransaksi)
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
		allHistory = append(allHistory, BookingResponse{
			ID: value.ID,
			// NamaPenginapan:   value.NamaPenginapan,
			CheckIn:          value.CheckIn,
			CheckOut:         value.CheckOut,
			TotalTransaksi:   value.TotalTransaksi,
			JenisTransaksi:   value.JenisTransaksi,
			StatusPembayaran: value.StatusPembayaran,
		})
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success read data", allHistory))
}
