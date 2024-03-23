//go:generate bash -c "mockgen -source=port.go -package=$(go list -f '{{.Name}}') -destination=port_mock_test.go"
package dhl

type Authenticator interface {
	Authenticate() (accessToken string, err error)
}

type DHLOrderCreatorAPI interface {
	Post(headers map[string]string, request DHLCreateOrderAPIRequest) (DHLCreateOrderAPIResponse, error)
}

type DHLAuthenticationAPI interface {
	Get(headers map[string]string, queryParam string, request DHLAuthenticationAPIRequest) (DHLAuthenticationAPIResponse, error)
}

type DHLOrderDeletorAPI interface {
	Post(headers map[string]string, request DHLDeleteOrderAPIRequest) (DHLDeleteOrderAPIResponse, error)
}
