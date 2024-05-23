package helper

import (
	"fmt"
	"strconv"
)

type HelperInterface interface {
	GetTotalDay(checkin string, checkout string) (int, error)
	//GetTotalDay(inpit int) (int, error)
}

type helper struct{}

func NewHelperService() HelperInterface {
	return &helper{}
}

func (h *helper) GetTotalDay(checkin string, checkout string) (int, error) {
	if len(checkin) < 2 {
		return 0, fmt.Errorf("input string too short")
	}
	if len(checkout) < 2 {
		return 0, fmt.Errorf("input string too short")
	}
	fmt.Println("[Helper Layer]")
	fmt.Println("[Helper Layer] CheckIn: ", checkin)
	fmt.Println("[Helper Layer] CheckOut: ", checkout)
	intCheckin, _ := strconv.Atoi(checkin[8:])
	intCheckout, _ := strconv.Atoi(checkout[8:])
	fmt.Println("[Helper Layer] intCheckin: ", intCheckin)
	fmt.Println("[Helper Layer] intCheckout: ", intCheckout)
	TotalDay := intCheckout - intCheckin
	fmt.Println("[Helper Layer] TotalDay: ", TotalDay)
	return TotalDay, nil
}
