package handler

import (
	"airbnb/features/review"
	"airbnb/utils/responses"
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

	reviewData, err := rh.reviewService.GetReviews(uint(idConv)) // Ambil data pengguna dari Redis
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse("error get reviews data", nil))
	}
	reviewsResponse := CoreToGorm(*reviewData)
	return c.JSON(http.StatusOK, responses.JSONWebResponse("success get review", reviewsResponse))
}
