package flash

type CreateOrderResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TID     string `json:"tid"`
}
