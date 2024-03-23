package shopee

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/salesX-technology/deliverypartnerconnectionlib"
)

var (
	_ deliverypartnerconnectionlib.OrderCreator = (*shopeeService)(nil)
	_ deliverypartnerconnectionlib.OrderUpdator = (*shopeeService)(nil)
)

type shopeeService struct {
	appID                uint64
	appSecret            string
	userID               uint64
	userSecret           string
	checkSignFunc        func(timestamp, randomInt64 int64, payload []byte) string
	unixFunc             func() int64
	randomFunc           func() int64
	shopeeCreateOrderAPI ShoppeeCreateOrderAPI
	shopeeUpdateOrderAPI ShopeeUpdateOrderAPI
	shopeeTimeSlotAPI    ShopeePickUpTimeAPI
}

type ShopeeServiceOption func(*shopeeService)

func NewShopeeService(
	appID uint64,
	appSecret string,
	userID uint64,
	userSecret string,
	shopeeCreateOrderAPI ShoppeeCreateOrderAPI,
	shopeeUpdateOrderAPI ShopeeUpdateOrderAPI,
	shopeeTimeSlotAPI ShopeePickUpTimeAPI,
	options ...ShopeeServiceOption,
) *shopeeService {
	svc := &shopeeService{
		appID:                appID,
		appSecret:            appSecret,
		userID:               userID,
		userSecret:           userSecret,
		checkSignFunc:        makeSignarureGenerator(appID, appSecret),
		unixFunc:             localUnixFunc,
		randomFunc:           secureRandomInt64,
		shopeeCreateOrderAPI: shopeeCreateOrderAPI,
		shopeeUpdateOrderAPI: shopeeUpdateOrderAPI,
		shopeeTimeSlotAPI:    shopeeTimeSlotAPI,
	}

	for _, option := range options {
		option(svc)
	}

	return svc
}

func WithCheckSignFunc(checkSignFunc func(randomInt64, timestamp int64, payload []byte) string) ShopeeServiceOption {
	return func(fs *shopeeService) {
		fs.checkSignFunc = checkSignFunc
	}
}

func WithRandomFunc(randomFunc func() int64) ShopeeServiceOption {
	return func(fs *shopeeService) {
		fs.randomFunc = randomFunc
	}
}

func WithUnixTimeFunc(unixFunc func() int64) ShopeeServiceOption {
	return func(fs *shopeeService) {
		fs.unixFunc = unixFunc
	}
}

func (f *shopeeService) CreateOrder(order deliverypartnerconnectionlib.Order) (string, error) {
	randomNumForRequest := f.randomFunc()

	timeStamp := f.unixFunc()

	timeSlotRequest := TimeSlotRequest{
		UserID:      f.userID,
		UserSecret:  f.userSecret,
		ServiceType: 1,
	}
	pickupTimeSlotRequestBodyBytes, _ := json.Marshal(timeSlotRequest)

	timeSlotResponse, _ := f.shopeeTimeSlotAPI.Post("/open/api/v1/order/get_pickup_time", map[string]string{
		"Content-Type": "application/json",
		"check-sign":   f.checkSignFunc(timeStamp, randomNumForRequest, pickupTimeSlotRequestBodyBytes),
		"app-id":       strconv.FormatUint(f.appID, 10),
		"timestamp":    strconv.FormatInt(timeStamp, 10),
		"random-num":   strconv.FormatInt(randomNumForRequest, 10),
	}, timeSlotRequest)

	shopeeCreateOrderRequestBody := CreateOrderRequest{
		UserID:     f.userID,
		UserSecret: f.userSecret,
		Orders: []Order{
			{
				OrderID: order.ID,
				BaseInfo: BaseInfo{
					ServiceType: 1,
				},
				FulfillmentInfo: FulfillmentInfo{
					PaymentRole:         1,
					CODCollection:       0,
					InsuranceCollection: 0,
					CollectType:         1,
					PickUpTime:          timeSlotResponse.Data[0].PickupTime,
					PickupTimeRangeID:   timeSlotResponse.Data[0].Slots[0].PickupTimeRangeID,
				},
				SenderInfo: SenderInfo{
					SenderName:          order.Sender.Name,
					SenderDetailAddress: order.Sender.AddressDetail,
					SenderState:         order.Sender.Province,
					SenderCity:          order.Sender.District,
					SenderDistrict:      order.Sender.SubDistrict,
					SenderPostCode:      order.Sender.PostalCode,
					SenderPhone:         order.Sender.Phone,
				},
				DeliverInfo: DeliverInfo{
					DeliverName:          order.Receiver.Name,
					DeliverDetailAddress: order.Receiver.AddressDetail,
					DeliverState:         order.Receiver.Province,
					DeliverCity:          order.Receiver.District,
					DeliverDistrict:      order.Receiver.SubDistrict,
					DeliverPostCode:      order.Receiver.PostalCode,
					DeliverPhone:         order.Receiver.Phone,
				},
				ParcelInfo: ParcelInfo{
					ParcelWeight:   order.WeightInGram / 1000,
					ParcelItemName: "parcel",
				},
			},
		},
	}

	shopeeCreateOrderRequestBodyBytes, _ := json.Marshal(shopeeCreateOrderRequestBody)
	timeStamp = f.unixFunc()

	response, _ := f.shopeeCreateOrderAPI.Post(
		"/open/api/v1/order/batch_create_order",
		map[string]string{
			"Content-Type": "application/json",
			"check-sign":   f.checkSignFunc(timeStamp, randomNumForRequest, shopeeCreateOrderRequestBodyBytes),
			"app-id":       strconv.FormatUint(f.appID, 10),
			"timestamp":    strconv.FormatInt(timeStamp, 10),
			"random-num":   strconv.FormatInt(randomNumForRequest, 10),
		}, shopeeCreateOrderRequestBody,
	)

	if response.RetCode != 0 {
		return "", fmt.Errorf("shopee create order failed with ret_code: %d", response.RetCode)
	}

	if len(response.Data.Orders) == 0 {
		return "", fmt.Errorf("shopee create order failed with no order in response")
	}

	return response.Data.Orders[0].TrackingNo, nil
}

func GenerateCheckSign(appId uint64, secret string, timestamp, random int64, payload []byte) (string, error) {
	originalValue := fmt.Sprintf("%d_%d_%d_%s", appId, timestamp, random, payload)
	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(originalValue))
	if err != nil {
		return "", err
	}
	checkSign := hex.EncodeToString(h.Sum(nil))
	return checkSign, nil
}

func (f *shopeeService) UpdateOrder(trackingNo string, order deliverypartnerconnectionlib.Order) error {
	randomNumForRequest := f.randomFunc()

	timeSlotRequest := TimeSlotRequest{
		UserID:      f.userID,
		UserSecret:  f.userSecret,
		ServiceType: 1,
	}
	pickupTimeSlotRequestBodyBytes, _ := json.Marshal(timeSlotRequest)
	timeStamp := f.unixFunc()
	timeSlotResponse, _ := f.shopeeTimeSlotAPI.Post(
		"/open/api/v1/order/get_pickup_time",
		map[string]string{
			"Content-Type": "application/json",
			"check-sign":   f.checkSignFunc(timeStamp, randomNumForRequest, pickupTimeSlotRequestBodyBytes),
			"app-id":       strconv.FormatUint(f.appID, 10),
			"timestamp":    strconv.FormatInt(timeStamp, 10),
			"random-num":   strconv.FormatInt(randomNumForRequest, 10),
		}, timeSlotRequest)

	updateOrderRequest := UpdateOrderRequest{
		UserID:     f.userID,
		UserSecret: f.userSecret,
		Orders: []Order{
			{
				TrackingNo: trackingNo,
				OrderID:    order.ID,
				BaseInfo: BaseInfo{
					ServiceType: 1,
				},
				FulfillmentInfo: FulfillmentInfo{
					PaymentRole:         1,
					CODCollection:       0,
					InsuranceCollection: 0,
					CollectType:         1,
					PickUpTime:          timeSlotResponse.Data[0].PickupTime,
					PickupTimeRangeID:   timeSlotResponse.Data[0].Slots[0].PickupTimeRangeID,
				},
				SenderInfo: SenderInfo{
					SenderName:          order.Sender.Name,
					SenderDetailAddress: order.Sender.AddressDetail,
					SenderState:         order.Sender.Province,
					SenderCity:          order.Sender.District,
					SenderDistrict:      order.Sender.SubDistrict,
					SenderPostCode:      order.Sender.PostalCode,
					SenderPhone:         order.Sender.Phone,
				},
				DeliverInfo: DeliverInfo{
					DeliverName:          order.Receiver.Name,
					DeliverDetailAddress: order.Receiver.AddressDetail,
					DeliverState:         order.Receiver.Province,
					DeliverCity:          order.Receiver.District,
					DeliverDistrict:      order.Receiver.SubDistrict,
					DeliverPostCode:      order.Receiver.PostalCode,
					DeliverPhone:         order.Receiver.Phone,
				},
				ParcelInfo: ParcelInfo{
					ParcelWeight:   order.WeightInGram / 1000,
					ParcelItemName: "parcel",
				},
			},
		},
	}

	updateOrderRequestBytes, _ := json.Marshal(updateOrderRequest)
	timeStamp = f.unixFunc()

	_, _ = f.shopeeUpdateOrderAPI.Post(
		"/open/api/v1/order/batch_update_order",
		map[string]string{
			"Content-Type": "application/json",
			"check-sign":   f.checkSignFunc(timeStamp, randomNumForRequest, updateOrderRequestBytes),
			"app-id":       strconv.FormatUint(f.appID, 10),
			"timestamp":    strconv.FormatInt(timeStamp, 10),
			"random-num":   strconv.FormatInt(randomNumForRequest, 10),
		}, updateOrderRequest,
	)

	return nil
}
