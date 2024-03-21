package dhl

import "github.com/salesX-technology/deliverypartnerconnectionlib"

type dhlService struct {
}

func NewDHLService() *dhlService {
	return &dhlService{}
}

func (f *dhlService) CreateOrder(order courierx.Order) (string, error) {
	return "", nil
}
