//go:generate bash -c "mockgen -source=port.go -package=$(go list -f '{{.Name}}') -destination=port_mock_test.go"

package flash

type FlashCreateOrderAPI interface {
	PostForm(fullurl string, form map[string]string) (FlashCreateOrderAPIResponse, error)
}

type FlashUpdateShipmentInfo interface {
	PostForm(fullurl string, form map[string]string) (FlashOrderUpdateAPIResponse, error)
}
