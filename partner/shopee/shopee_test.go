package shopee

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type ShopeeTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	mShopeeCreateOrderAPI *MockShoppeeCreateOrderAPI
	mShopeeUpdateOrderAPI *MockShopeeUpdateOrderAPI
	mShopeeCancelOrderAPI *MockShopeeCancelOrderAPI
	mTimeSlotAPI          *MockShopeePickUpTimeAPI

	service *shopeeService
}

func (t *ShopeeTestSuite) SetupTest() {
	t.ctrl = gomock.NewController(t.T())

	t.mShopeeCreateOrderAPI = NewMockShoppeeCreateOrderAPI(t.ctrl)
	t.mShopeeUpdateOrderAPI = NewMockShopeeUpdateOrderAPI(t.ctrl)
	t.mShopeeCancelOrderAPI = NewMockShopeeCancelOrderAPI(t.ctrl)
	t.mTimeSlotAPI = NewMockShopeePickUpTimeAPI(t.ctrl)

	userID := uint64(99999)
	t.service = NewShopeeService(
		10000,
		"app_secret",
		userID,
		"user_secret",
		t.mShopeeCreateOrderAPI,
		t.mShopeeUpdateOrderAPI,
		t.mShopeeCancelOrderAPI,
		t.mTimeSlotAPI,
		WithCheckSignFunc(func(randomInt64, timestamp int64, payload []byte) string {
			return "check-sign"
		}),
		WithUnixTimeFunc(func() int64 {
			// timestamp func
			return 12345
		}),
		WithRandomFunc(func() int64 {
			// random func
			return 98765
		}))
}

func TestShopeeTestSuite(t *testing.T) {
	suite.Run(t, new(ShopeeTestSuite))
}

func (t *ShopeeTestSuite) TearDownTest() {
	t.ctrl.Finish()
}

func (t *ShopeeTestSuite) TestGivenShopeeOrderIsCreating_WhenCreateOrder_ThenCallAdaptorShopeeCreateOrderAndReturnSuccess() {
	userID := uint64(99999)
	pickupTimeSlot := 1701981200

	t.mTimeSlotAPI.EXPECT().Post("/open/api/v1/order/get_pickup_time", map[string]string{
		"Content-Type": "application/json",
		"app-id":       "10000",
		"check-sign":   "check-sign",
		"timestamp":    "12345",
		"random-num":   "98765",
	}, TimeSlotRequest{
		UserID:      userID,
		UserSecret:  "user_secret",
		ServiceType: 1,
	}).Return(TimeSlotResponse{
		Data: []TimeSlotData{
			{
				PickupTime: pickupTimeSlot,
				Slots:      []PickupSlot{{PickupTimeRangeID: 1}},
			},
		},
		RetCode: 0,
	}, nil)

	t.mShopeeCreateOrderAPI.EXPECT().Post(
		"/open/api/v1/order/batch_create_order",
		map[string]string{
			"Content-Type": "application/json",
			"app-id":       "10000",
			"check-sign":   "check-sign",
			"timestamp":    "12345",
			"random-num":   "98765",
		}, CreateOrderRequest{
			UserID:     userID,
			UserSecret: "user_secret",
			Orders: []Order{
				{
					OrderID: "order_id",
					BaseInfo: BaseInfo{
						ServiceType: 1,
					},
					FulfillmentInfo: FulfillmentInfo{
						PaymentRole:         1,
						CODCollection:       0,
						InsuranceCollection: 0,
						CollectType:         1,
						PickUpTime:          pickupTimeSlot,
						PickupTimeRangeID:   1,
					},
					SenderInfo: SenderInfo{
						SenderName:          aValidNonCODOrder.Sender.Name,
						SenderDetailAddress: aValidNonCODOrder.Sender.AddressDetail,
						SenderState:         aValidNonCODOrder.Sender.Province,
						SenderCity:          aValidNonCODOrder.Sender.District,
						SenderDistrict:      aValidNonCODOrder.Sender.SubDistrict,
						SenderPostCode:      aValidNonCODOrder.Sender.PostalCode,
						SenderPhone:         aValidNonCODOrder.Sender.Phone,
					},
					DeliverInfo: DeliverInfo{
						DeliverName:          aValidNonCODOrder.Receiver.Name,
						DeliverDetailAddress: aValidNonCODOrder.Receiver.AddressDetail,
						DeliverState:         aValidNonCODOrder.Receiver.Province,
						DeliverCity:          aValidNonCODOrder.Receiver.District,
						DeliverDistrict:      aValidNonCODOrder.Receiver.SubDistrict,
						DeliverPostCode:      aValidNonCODOrder.Receiver.PostalCode,
						DeliverPhone:         aValidNonCODOrder.Receiver.Phone,
					},
					ParcelInfo: ParcelInfo{
						ParcelWeight:   1,
						ParcelItemName: "parcel",
					},
				},
			},
		}).Return(CreateOrderResponse{Data: Data{
		Orders: []OrderResponse{
			{
				TrackingNo: "tracking_no",
			},
		},
	}}, nil)

	orderID, err := t.service.CreateOrder(aValidNonCODOrder)

	t.Equal("tracking_no", orderID)
	t.Nil(err)
}

func (t *ShopeeTestSuite) TestGivenCreateShopeeOrder_WhenReturnCodeIsNot0_ThenReturnError() {
	userID := uint64(99999)
	pickupTimeSlot := 1701981200
	t.mTimeSlotAPI.EXPECT().Post(
		"/open/api/v1/order/get_pickup_time",
		map[string]string{
			"Content-Type": "application/json",
			"app-id":       "10000",
			"check-sign":   "check-sign",
			"timestamp":    "12345",
			"random-num":   "98765",
		}, TimeSlotRequest{
			UserID:      userID,
			UserSecret:  "user_secret",
			ServiceType: 1,
		}).Return(TimeSlotResponse{
		Data: []TimeSlotData{
			{
				PickupTime: pickupTimeSlot,
				Slots:      []PickupSlot{{PickupTimeRangeID: 1}},
			},
		},
		RetCode: 0,
	}, nil)

	t.mShopeeCreateOrderAPI.EXPECT().Post(
		"/open/api/v1/order/batch_create_order",
		map[string]string{
			"Content-Type": "application/json",
			"app-id":       "10000",
			"check-sign":   "check-sign",
			"timestamp":    "12345",
			"random-num":   "98765",
		}, CreateOrderRequest{
			UserID:     userID,
			UserSecret: "user_secret",
			Orders: []Order{
				{
					OrderID: "order_id",
					BaseInfo: BaseInfo{
						ServiceType: 1,
					},
					FulfillmentInfo: FulfillmentInfo{
						PaymentRole:         1,
						CODCollection:       0,
						InsuranceCollection: 0,
						CollectType:         1,
						PickUpTime:          pickupTimeSlot,
						PickupTimeRangeID:   1,
					},
					SenderInfo: SenderInfo{
						SenderName:          aValidNonCODOrder.Sender.Name,
						SenderDetailAddress: aValidNonCODOrder.Sender.AddressDetail,
						SenderState:         aValidNonCODOrder.Sender.Province,
						SenderCity:          aValidNonCODOrder.Sender.District,
						SenderDistrict:      aValidNonCODOrder.Sender.SubDistrict,
						SenderPostCode:      aValidNonCODOrder.Sender.PostalCode,
						SenderPhone:         aValidNonCODOrder.Sender.Phone,
					},
					DeliverInfo: DeliverInfo{
						DeliverName:          aValidNonCODOrder.Receiver.Name,
						DeliverDetailAddress: aValidNonCODOrder.Receiver.AddressDetail,
						DeliverState:         aValidNonCODOrder.Receiver.Province,
						DeliverCity:          aValidNonCODOrder.Receiver.District,
						DeliverDistrict:      aValidNonCODOrder.Receiver.SubDistrict,
						DeliverPostCode:      aValidNonCODOrder.Receiver.PostalCode,
						DeliverPhone:         aValidNonCODOrder.Receiver.Phone,
					},
					ParcelInfo: ParcelInfo{
						ParcelWeight:   1,
						ParcelItemName: "parcel",
					},
				},
			},
		}).Return(CreateOrderResponse{
		RetCode: 9999,
	}, nil)

	orderID, err := t.service.CreateOrder(aValidNonCODOrder)

	t.Empty(orderID)
	t.EqualError(err, "shopee create order failed with ret_code: 9999")
}

func (t *ShopeeTestSuite) TestGivenCreateShopeeOrder_WhenNoOrderInResponse_ThenReturnError() {
	userID := uint64(99999)
	pickupTimeSlot := 1701981200
	t.mTimeSlotAPI.EXPECT().Post(
		"/open/api/v1/order/get_pickup_time",
		map[string]string{
			"Content-Type": "application/json",
			"app-id":       "10000",
			"check-sign":   "check-sign",
			"timestamp":    "12345",
			"random-num":   "98765",
		}, TimeSlotRequest{
			UserID:      userID,
			UserSecret:  "user_secret",
			ServiceType: 1,
		}).Return(TimeSlotResponse{
		Data: []TimeSlotData{
			{
				PickupTime: pickupTimeSlot,
				Slots:      []PickupSlot{{PickupTimeRangeID: 1}},
			},
		},
		RetCode: 0,
	}, nil)

	t.mShopeeCreateOrderAPI.EXPECT().Post(
		"/open/api/v1/order/batch_create_order",
		map[string]string{
			"Content-Type": "application/json",
			"app-id":       "10000",
			"check-sign":   "check-sign",
			"timestamp":    "12345",
			"random-num":   "98765",
		}, CreateOrderRequest{
			UserID:     userID,
			UserSecret: "user_secret",
			Orders: []Order{
				{
					OrderID: "order_id",
					BaseInfo: BaseInfo{
						ServiceType: 1,
					},
					FulfillmentInfo: FulfillmentInfo{
						PaymentRole:         1,
						CODCollection:       0,
						InsuranceCollection: 0,
						CollectType:         1,
						PickUpTime:          pickupTimeSlot,
						PickupTimeRangeID:   1,
					},
					SenderInfo: SenderInfo{
						SenderName:          aValidNonCODOrder.Sender.Name,
						SenderDetailAddress: aValidNonCODOrder.Sender.AddressDetail,
						SenderState:         aValidNonCODOrder.Sender.Province,
						SenderCity:          aValidNonCODOrder.Sender.District,
						SenderDistrict:      aValidNonCODOrder.Sender.SubDistrict,
						SenderPostCode:      aValidNonCODOrder.Sender.PostalCode,
						SenderPhone:         aValidNonCODOrder.Sender.Phone,
					},
					DeliverInfo: DeliverInfo{
						DeliverName:          aValidNonCODOrder.Receiver.Name,
						DeliverDetailAddress: aValidNonCODOrder.Receiver.AddressDetail,
						DeliverState:         aValidNonCODOrder.Receiver.Province,
						DeliverCity:          aValidNonCODOrder.Receiver.District,
						DeliverDistrict:      aValidNonCODOrder.Receiver.SubDistrict,
						DeliverPostCode:      aValidNonCODOrder.Receiver.PostalCode,
						DeliverPhone:         aValidNonCODOrder.Receiver.Phone,
					},
					ParcelInfo: ParcelInfo{
						ParcelWeight:   1,
						ParcelItemName: "parcel",
					},
				},
			},
		}).Return(CreateOrderResponse{
		RetCode: 0,
	}, nil)

	orderID, err := t.service.CreateOrder(aValidNonCODOrder)

	t.Empty(orderID)
	t.EqualError(err, "shopee create order failed with no order in response")
}

func (t *ShopeeTestSuite) TestGivenShopeeOrderIsUpdating_WhenUpdateOrder_ThenCallAdaptorShopeeUpdateOrderAndReturnSuccess() {
	userID := uint64(99999)

	pickupTimeSlot := 1701981200
	t.mTimeSlotAPI.EXPECT().Post(
		"/open/api/v1/order/get_pickup_time",
		map[string]string{
			"Content-Type": "application/json",
			"app-id":       "10000",
			"check-sign":   "check-sign",
			"timestamp":    "12345",
			"random-num":   "98765",
		}, TimeSlotRequest{
			UserID:      userID,
			UserSecret:  "user_secret",
			ServiceType: 1,
		}).Return(TimeSlotResponse{
		Data: []TimeSlotData{
			{
				PickupTime: pickupTimeSlot,
				Slots:      []PickupSlot{{PickupTimeRangeID: 1}},
			},
		},
		RetCode: 0,
	}, nil)

	t.mShopeeUpdateOrderAPI.EXPECT().Post(
		"/open/api/v1/order/batch_update_order",
		map[string]string{
			"Content-Type": "application/json",
			"app-id":       "10000",
			"check-sign":   "check-sign",
			"timestamp":    "12345",
			"random-num":   "98765",
		}, UpdateOrderRequest{
			UserID:     userID,
			UserSecret: "user_secret",
			Orders: []Order{
				{
					TrackingNo: "tracking_no",
					OrderID:    "order_id",
					BaseInfo: BaseInfo{
						ServiceType: 1,
					},
					FulfillmentInfo: FulfillmentInfo{
						PaymentRole:         1,
						CODCollection:       0,
						InsuranceCollection: 0,
						CollectType:         1,
						PickUpTime:          pickupTimeSlot,
						PickupTimeRangeID:   1,
					},
					SenderInfo: SenderInfo{
						SenderName:          aValidNonCODOrder.Sender.Name,
						SenderDetailAddress: aValidNonCODOrder.Sender.AddressDetail,
						SenderState:         aValidNonCODOrder.Sender.Province,
						SenderCity:          aValidNonCODOrder.Sender.District,
						SenderDistrict:      aValidNonCODOrder.Sender.SubDistrict,
						SenderPostCode:      aValidNonCODOrder.Sender.PostalCode,
						SenderPhone:         aValidNonCODOrder.Sender.Phone,
					},
					DeliverInfo: DeliverInfo{
						DeliverName:          aValidNonCODOrder.Receiver.Name,
						DeliverDetailAddress: aValidNonCODOrder.Receiver.AddressDetail,
						DeliverState:         aValidNonCODOrder.Receiver.Province,
						DeliverCity:          aValidNonCODOrder.Receiver.District,
						DeliverDistrict:      aValidNonCODOrder.Receiver.SubDistrict,
						DeliverPostCode:      aValidNonCODOrder.Receiver.PostalCode,
						DeliverPhone:         aValidNonCODOrder.Receiver.Phone,
					},
					ParcelInfo: ParcelInfo{
						ParcelWeight:   1,
						ParcelItemName: "parcel",
					},
				},
			},
		}).Return(UpdateOrderResponse{
		RetCode: 0,
	}, nil)

	err := t.service.UpdateOrder("tracking_no", aValidNonCODOrder)
	t.NoError(err)
}

func (t *ShopeeTestSuite) TestGivenOrderIsExist_WhenCancelOrderSuccess_ThenReturnNoError() {
	t.mShopeeCancelOrderAPI.EXPECT().Post(
		"/open/api/v1/order/batch_cancel_order",
		map[string]string{
			"Content-Type": "application/json",
			"app-id":       "10000",
			"check-sign":   "check-sign",
			"timestamp":    "12345",
			"random-num":   "98765",
		},
		CancelOrderRequest{
			UserID:         99999,
			UserSecret:     "user_secret",
			TrackingNoList: []string{"tracking_no"},
		}).
		Return(CancelOrderResponse{}, nil)

	err := t.service.DeleteOrder("tracking_no")
	t.NoError(err)
}
