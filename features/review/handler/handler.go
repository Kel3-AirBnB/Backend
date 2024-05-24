package handler

import (
	"airbnb/app/middlewares"
	"airbnb/features/review"
	"airbnb/utils/responses"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ReviewHandler struct {
	reviewService review.ServiceInterface
}

func New(rs review.ServiceInterface) *ReviewHandler {
	return &ReviewHandler{
		reviewService: rs,
	}
}
func (rh *ReviewHandler) GetAll(c echo.Context) error {
	result, err := rh.reviewService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error read data",
		})
	}
	var allReviewsResponse []ReviewResponse
	for _, value := range result {
		allReviewsResponse = append(allReviewsResponse, ReviewResponse{
			ID:           value.ID,
			PenginapanID: value.PenginapanID,
			UserID:       value.UserID,
			Rating:       value.Rating,
			Komentar:     value.Komentar,
			Foto:         value.Foto,
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read data",
		"results": allReviewsResponse,
	})
}
func (rh *ReviewHandler) CreateReview(c echo.Context) error {
	newReview := ReviewRequest{}
	errBind := c.Bind(&newReview)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error bind"+errBind.Error(), nil))
	}

	file, handler, err := c.Request().FormFile("review_profile")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Unable to upload photo: " + err.Error(),
		})
	}
	defer file.Close()

	inputCore := RequestToCore(newReview)

	_, errInsert := rh.reviewService.Create(inputCore, file, handler.Filename)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error add data", errInsert))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error add data", errInsert))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("success add data", nil))
}
func (rh *ReviewHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get reviews id", idConv))
	}

	reviewData, err := rh.reviewService.GetReviews(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get reviews data", nil))
	}
	reviewsResponse := CoreToGorm(*reviewData)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success get review", reviewsResponse))
}
func (rh *ReviewHandler) Delete(c echo.Context) error {
	// id := c.Param("id")
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error get reviews id", idConv))
	}
	// idToken := middlewares.ExtractTokenUserId(c)
	// err := rh.reviewService.Delete(uint(idToken))
	err := rh.reviewService.Delete(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error delete data", err))
	}
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success delete data", err))
}

func (rh *ReviewHandler) UpdateReview(c echo.Context) error {
	// Ambil ID ulasan dari URL
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error converting review ID", nil))
	}

	// Bind data permintaan ke struct ReviewRequest
	updateRequest := ReviewRequest{}
	if err := c.Bind(&updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error binding update data", nil))
	}

	// Cek apakah ada file yang diunggah
	file, handler, err := c.Request().FormFile("review_profile")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Unable to upload photo: " + err.Error(),
		})
	}
	defer file.Close()

	// Ubah data permintaan menjadi Core struct
	inputCore := RequestToCore(updateRequest)

	// Panggil layanan untuk melakukan pembaruan ulasan
	_, errUpdate := rh.reviewService.UpdateById(uint(idConv), inputCore, file, handler.Filename)
	if errUpdate != nil {
		if strings.Contains(errUpdate.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("error update data", errUpdate))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error update data", errUpdate))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse("success update data", nil))
}

func (rh *ReviewHandler) GetReviewsByUserID(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c) // extract id user from jwt token
	log.Println("idtoken:", idToken)
	reviewData, err := rh.reviewService.GetReviewsByUserID(uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get review data", nil))
	}
	ReviewResponse := CoreToGorm(*reviewData)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success get review by user_id", ReviewResponse))
}

func (rh *ReviewHandler) GetReviewByPenginapanID(c echo.Context) error {
	penginapanIDParam := c.Param("penginapanID")
	penginapanID, err := strconv.Atoi(penginapanIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.JSONWebResponse("invalid penginapanID", err))
	}

	reviewData, err := rh.reviewService.GetReviewByPenginapanID(uint(penginapanID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get reviews by penginapanID", nil))
	}

	return c.JSON(http.StatusOK, responses.JSONWebResponse("success get reviews by penginapan_id", reviewData))
}
