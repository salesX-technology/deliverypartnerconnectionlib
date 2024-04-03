package shopee

type CreateOrderResponse struct {
	RetCode int  `json:"ret_code"`
	Data    Data `json:"data"`
}

type Data struct {
	Orders []OrderResponse `json:"orders"`
}

type OrderResponse struct {
	OrderID              string  `json:"order_id"`
	TrackingNo           string  `json:"tracking_no"`
	TrackingLink         string  `json:"tracking_link"`
	RFirstSortCode       string  `json:"r_first_sort_code"`
	RThirdSortCode       string  `json:"r_third_sort_code"`
	ReturnFirstSortCode  string  `json:"return_first_sort_code"`
	EstimatedShippingFee float64 `json:"estimated_shipping_fee"`
	BasicShippingFee     float64 `json:"basic_shipping_fee"`
	CODServiceFee        float64 `json:"cod_service_fee"`
	InsuranceServiceFee  float64 `json:"insurance_service_fee"`
	VATFee               float64 `json:"vat_fee"`
}

type UpdateOrderResponse struct {
	RetCode int `json:"ret_code"`
}

type CancelOrderResponse struct {
}

type ShopeeCreateOrderResponse struct {
	OrderID              string  `json:"order_id"`
	TrackingNo           string  `json:"tracking_no"`
	TrackingLink         string  `json:"tracking_link"`
	RFirstSortCode       string  `json:"r_first_sort_code"`
	RThirdSortCode       string  `json:"r_third_sort_code"`
	ReturnFirstSortCode  string  `json:"return_first_sort_code"`
	EstimatedShippingFee float64 `json:"estimated_shipping_fee"`
	BasicShippingFee     float64 `json:"basic_shipping_fee"`
	CODServiceFee        float64 `json:"cod_service_fee"`
	InsuranceServiceFee  float64 `json:"insurance_service_fee"`
	VATFee               float64 `json:"vat_fee"`
}

type HookOrderResponse struct {
	RetCode int  `json:"ret_code"`
	Data    Data `json:"data"`
}

type HookData struct {
	Orders []HookOrderResponse `json:"orders"`
}

type HookResponse struct {
	OrderID      string `json:"order_id"`
	OrderIDLink  string `json:"order_id_link"`
	TrackingNo   string `json:"tracking_no"`
	TrackingLink string `json:"tracking_link"`
}
