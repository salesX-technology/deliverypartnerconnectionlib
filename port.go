//go:generate bash -c "mockgen -source=port.go -package=$(go list -f '{{.Name}}') -destination=port_mock_test.go"

package deliverypartnerconnectionlib

type OrderCreator interface {
	//trackingNo string ปรับให้ออกเป็น array อะไรก็ได้
	// example.CreateOrderResponse ปรับให้ออกเป็น string อะไรก็ได้

	//CreateOrder(order Order) (trackingNo string, err error)
	CreateOrder(order Order) (trackingNo map[string]interface{}, err error)
}

type OrderUpdator interface {
	UpdateOrder(trackingNo string, order Order) error
}

type OrderDeleter interface {
	DeleteOrder(trackingNo string) error
}
