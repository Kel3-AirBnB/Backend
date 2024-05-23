package helper

import (
	"fmt"
	"strconv"
)

type HelperInterface interface {
	GetTotalDay(checkin string, checkout string) (uint, error)
	//GetTotalDay(inpit int) (int, error)
}

type helper struct{}

func NewHelperService() HelperInterface {
	return &helper{}
}

func (h *helper) GetTotalDay(checkin string, checkout string) (uint, error) {
	if len(checkin) < 2 {
		return 0, fmt.Errorf("input string too short")
	}
	if len(checkout) < 2 {
		return 0, fmt.Errorf("input string too short")
	}
	uintCheckin, _ := strconv.Atoi(checkin[:2])
	uintCheckout, _ := strconv.Atoi(checkout[:2])
	// TotalDay :=
	return uint(uintCheckout) - uint(uintCheckin), nil
}
