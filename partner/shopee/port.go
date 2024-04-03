//go:generate bash -c "mockgen -source=port.go -package=$(go list -f '{{.Name}}') -destination=port_mock_test.go"
package shopee

type ShoppeeCreateOrderAPI interface {
	Post(endpoint string, headers map[string]string, request CreateOrderRequest) (CreateOrderResponse, error)
}

type ShopeePickUpTimeAPI interface {
	Post(endpoint string, headers map[string]string, request TimeSlotRequest) (TimeSlotResponse, error)
}

type ShopeeUpdateOrderAPI interface {
	Post(endpoint string, headers map[string]string, request UpdateOrderRequest) (UpdateOrderResponse, error)
}

type ShopeeCancelOrderAPI interface {
	Post(endpoint string, headers map[string]string, request CancelOrderRequest) (CancelOrderResponse, error)
}

type ShopeeHookOrderAPI interface {
	Post(endpoint string, headers map[string]string, request HookOrderRequest) (HookOrderResponse, error)
}
