package shopee

type CreateOrderResponse struct {
	RetCode int  `json:"ret_code"`
	Data    Data `json:"data"`
}

type Data struct {
	Orders []OrderResponse `json:"orders"`
}

type OrderResponse struct {
	TrackingNo string `json:"tracking_no"`
}

type UpdateOrderResponse struct {
	RetCode int `json:"ret_code"`
}

type CancelOrderResponse struct {
}
