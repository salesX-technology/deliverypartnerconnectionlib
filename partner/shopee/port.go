//go:generate bash -c "mockgen -source=port.go -package=$(go list -f '{{.Name}}') -destination=port_mock_test.go"
package shopee

type ShoppeeCreateOrderAPI interface {
	Post(headers map[string]string, request CreateOrderRequest) (CreateOrderResponse, error)
}

type ShopeePickUpTimeAPI interface {
	Post(headers map[string]string, request TimeSlotRequest) (TimeSlotResponse, error)
}

type ShopeeUpdateOrderAPI interface {
	Post(headers map[string]string, request UpdateOrderRequest) (UpdateOrderResponse, error)
}
