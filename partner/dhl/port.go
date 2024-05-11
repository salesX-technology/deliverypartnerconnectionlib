//go:generate bash -c "mockgen -source=port.go -package=$(go list -f '{{.Name}}') -destination=port_mock_test.go"
package dhl

type Authenticator interface {
	Authenticate() (accessToken string, err error)
}

type DHLOrderCreatorAPI interface {
	Post(endpoint string, headers map[string]string, request DHLCreateOrderAPIRequest) (DHLCreateOrderAPIResponse, error)
}

type DHLAuthenticationAPI interface {
	Get(endpoint string, headers map[string]string, request DHLAuthenticationAPIRequest) (DHLAuthenticationAPIResponse, error)
}

type DHLOrderDeletorAPI interface {
	Post(endpoint string, headers map[string]string, request DHLDeleteOrderAPIRequest) (DHLDeleteOrderAPIResponse, error)
}

type DHLOrderUpdatorAPI interface {
	Post(endpoint string, headers map[string]string, request DHLUpdateOrderAPIRequest) (DHLUpdateOrderAPIResponse, error)
}

type DHLHookOrderAPI interface {
	PostHook(endpoint string, headers map[string]string, request DHLHookOrderAPIRequest) (DHLHookOrderAPIResponse, error)
}
