//go:generate bash -c "mockgen -source=port.go -package=$(go list -f '{{.Name}}') -destination=port_mock_test.go"

package courierx

type OrderCreator interface {
	CreateOrder(order Order) (trackingNo string, err error)
}

type OrderUpdator interface {
	UpdateOrder(trackingNo string, order Order) error
}
