package handler

import (
	"airbnb/features/review"
	"net/http"

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
