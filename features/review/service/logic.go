package service

import "airbnb/features/review"

type reviewService struct {
	reviewData review.DataInterface
}

// GetAll implements review.ServiceInterface.
func (r *reviewService) GetAll() ([]review.Core, error) {
	return r.reviewData.SelectAll()
}

func New(rd review.DataInterface) review.ServiceInterface {
	return &reviewService{
		reviewData: rd,
	}
}
