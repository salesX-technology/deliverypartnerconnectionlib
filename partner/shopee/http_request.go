package shopee

type CreateOrderRequest struct {
	UserID     uint64  `json:"user_id"`
	UserSecret string  `json:"user_secret"`
	Orders     []Order `json:"orders"`
}

type Order struct {
	OrderID         string          `json:"order_id"`
	BaseInfo        BaseInfo        `json:"base_info"`
	SenderInfo      SenderInfo      `json:"sender_info"`
	DeliverInfo     DeliverInfo     `json:"deliver_info"`
	FulfillmentInfo FulfillmentInfo `json:"fulfillment_info"`
	ParcelInfo      ParcelInfo      `json:"parcel_info"`
	TrackingNo      string          `json:"tracking_no"`
}

type BaseInfo struct {
	ServiceType int `json:"service_type"`
}

type FulfillmentInfo struct {
	PaymentRole         int `json:"payment_role"`
	CODCollection       int `json:"cod_collection"`
	InsuranceCollection int `json:"insurance_collection"`
	CollectType         int `json:"collect_type"`
	PickUpTime          int `json:"pickup_time"`
	PickupTimeRangeID   int `json:"pickup_time_range_id"`
}

type SenderInfo struct {
	SenderName          string `json:"sender_name"`
	SenderDetailAddress string `json:"sender_detail_address"`
	SenderState         string `json:"sender_state"`
	SenderCity          string `json:"sender_city"`
	SenderDistrict      string `json:"sender_district"`
	SenderPostCode      string `json:"sender_post_code"`
	SenderPhone         string `json:"sender_phone"`
}

type DeliverInfo struct {
	DeliverName          string `json:"deliver_name"`
	DeliverDetailAddress string `json:"deliver_detail_address"`
	DeliverState         string `json:"deliver_state"`
	DeliverCity          string `json:"deliver_city"`
	DeliverDistrict      string `json:"deliver_district"`
	DeliverPostCode      string `json:"deliver_post_code"`
	DeliverPhone         string `json:"deliver_phone"`
}

type ParcelInfo struct {
	ParcelWeight   float64 `json:"parcel_weight"`
	ParcelItemName string  `json:"parcel_item_name"`
}

type TimeSlotRequest struct {
	UserID      uint64 `json:"user_id"`
	UserSecret  string `json:"user_secret"`
	ServiceType int    `json:"service_type"`
}

type TimeSlotResponse struct {
	RetCode int            `json:"ret_code"`
	Data    []TimeSlotData `json:"data"`
}

type TimeSlotData struct {
	PickupTime int          `json:"pickup_time"`
	Slots      []PickupSlot `json:"slots"`
}

type PickupSlot struct {
	PickupTimeRangeID int `json:"pickup_time_range_id"`
}

type UpdateOrderRequest struct {
	UserID     uint64  `json:"user_id"`
	UserSecret string  `json:"user_secret"`
	Orders     []Order `json:"orders"`
}

type CancelOrderRequest struct {
	UserID         uint64   `json:"user_id"`
	UserSecret     string   `json:"user_secret"`
	TrackingNoList []string `json:"tracking_no_list"`
}

type HookOrderRequest struct {
	UserID         uint64   `json:"user_id"`
	UserSecret     string   `json:"user_secret"`
	TrackingNoList []string `json:"tracking_no_list"`
}
