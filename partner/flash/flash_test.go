package flash

import (
	"errors"
	"testing"

	"github.com/salesX-technology/deliverypartnerconnectionlib"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type FlashTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	mFlashCreateOrderAPI     *MockFlashCreateOrderAPI
	mFlashUpdateShipmentInfo *MockFlashUpdateShipmentInfo
	mFlashDeleteOrder        *MockFlashDeleteOrderAPI

	service *flashService
}

func (t *FlashTestSuite) SetupTest() {
	t.ctrl = gomock.NewController(t.T())
	t.mFlashCreateOrderAPI = NewMockFlashCreateOrderAPI(t.ctrl)
	t.mFlashUpdateShipmentInfo = NewMockFlashUpdateShipmentInfo(t.ctrl)
	t.mFlashDeleteOrder = NewMockFlashDeleteOrderAPI(t.ctrl)

	t.service = NewFlashService(
		t.mFlashCreateOrderAPI,
		t.mFlashUpdateShipmentInfo,
		t.mFlashDeleteOrder,
		"secret",
		"merchant",
		WithNonceGenerator(func(int) string {
			return "nonce"
		}),
		WithSignatureGenerator(func(map[string]string, string) string {
			return "signature"
		}),
	)
}

func TestFlashTestSuite(t *testing.T) {
	suite.Run(t, new(FlashTestSuite))
}

func (t *FlashTestSuite) TestGivenNonCODOrderIsCreating_WhenCreateOrder_ThenCreateSuccess() {
	t.mFlashCreateOrderAPI.EXPECT().PostForm("/open/v3/orders", map[string]string{
		"articleCategory":  "99",
		"codEnabled":       "0",
		"dstCityName":      "สันทราย",
		"dstDistrictName":  "สันทรายน้อย",
		"dstDetailAddress": "example detail address",
		"dstName":          "น้ำพริกแม่อำพร",
		"dstPhone":         "0123456789",
		"dstPostalCode":    "50210",
		"dstProvinceName":  "เชียงใหม่",
		"expressCategory":  "1",
		"insured":          "0",
		"mchId":            "merchant",
		"nonceStr":         "nonce",
		"srcCityName":      "เมืองอุบลราชธานี",
		"srcDetailAddress": "example detail address",
		"srcName":          "หอมรวม  create order test name",
		"srcDistrictName":  "ขี้เหล็ก",
		"srcPhone":         "0123456789",
		"srcPostalCode":    "34000",
		"srcProvinceName":  "อุบลราชธานี",
		"weight":           "1000",
		"sign":             "signature",
	}).Return(FlashCreateOrderAPIResponse{
		Data: FlasFlashCreateOrderAPIResponseData{
			PNO: "trackingNo",
		},
	}, nil)

	trackingNo, err := t.service.CreateOrder(deliverypartnerconnectionlib.Order{
		WeightInGram: 1000,
		IsCOD:        false,
		Sender: deliverypartnerconnectionlib.OrderAddress{
			Name:          "หอมรวม  create order test name",
			AddressDetail: "example detail address",
			SubDistrict:   "ขี้เหล็ก",
			District:      "เมืองอุบลราชธานี",
			Province:      "อุบลราชธานี",
			Phone:         "0123456789",
			PostalCode:    "34000",
		},
		Receiver: deliverypartnerconnectionlib.OrderAddress{
			Name:          "น้ำพริกแม่อำพร",
			AddressDetail: "example detail address",
			SubDistrict:   "สันทรายน้อย",
			District:      "สันทราย",
			Province:      "เชียงใหม่",
			Phone:         "0123456789",
			PostalCode:    "50210",
		},
	})

	t.Equal("trackingNo", trackingNo)
	t.NoError(err)
}

func (t *FlashTestSuite) TestGivenCODOrderIsCreating_WhenCreateOrder_ThenCreateSuccess() {
	t.mFlashCreateOrderAPI.EXPECT().PostForm("/open/v3/orders", map[string]string{
		"articleCategory":  "99",
		"codEnabled":       "1",
		"dstCityName":      "สันทราย",
		"dstDistrictName":  "สันทรายน้อย",
		"dstDetailAddress": "example detail address",
		"dstName":          "น้ำพริกแม่อำพร",
		"dstPhone":         "0123456789",
		"dstPostalCode":    "50210",
		"dstProvinceName":  "เชียงใหม่",
		"expressCategory":  "1",
		"insured":          "0",
		"mchId":            "merchant",
		"nonceStr":         "nonce",
		"srcDistrictName":  "ขี้เหล็ก",
		"srcCityName":      "เมืองอุบลราชธานี",
		"srcDetailAddress": "example detail address",
		"srcName":          "หอมรวม  create order test name",
		"srcPhone":         "0123456789",
		"srcPostalCode":    "34000",
		"srcProvinceName":  "อุบลราชธานี",
		"weight":           "1000",
		"sign":             "signature",
	}).Return(FlashCreateOrderAPIResponse{
		Data: FlasFlashCreateOrderAPIResponseData{
			PNO: "trackingNo",
		},
	}, nil)

	trackingNo, err := t.service.CreateOrder(deliverypartnerconnectionlib.Order{
		WeightInGram: 1000,
		IsCOD:        true,
		Sender: deliverypartnerconnectionlib.OrderAddress{
			Name:          "หอมรวม  create order test name",
			AddressDetail: "example detail address",
			SubDistrict:   "ขี้เหล็ก",
			District:      "เมืองอุบลราชธานี",
			Province:      "อุบลราชธานี",
			Phone:         "0123456789",
			PostalCode:    "34000",
		},
		Receiver: deliverypartnerconnectionlib.OrderAddress{
			Name:          "น้ำพริกแม่อำพร",
			AddressDetail: "example detail address",
			SubDistrict:   "สันทรายน้อย",
			District:      "สันทราย",
			Province:      "เชียงใหม่",
			Phone:         "0123456789",
			PostalCode:    "50210",
		},
	})

	t.Equal("trackingNo", trackingNo)
	t.NoError(err)
}

func (t *FlashTestSuite) TestHTTPPostFormIsFailed_WhenCreateOrder_ThenReturnError() {
	t.mFlashCreateOrderAPI.EXPECT().PostForm("/open/v3/orders", map[string]string{
		"articleCategory":  "99",
		"codEnabled":       "1",
		"dstCityName":      "สันทราย",
		"dstDistrictName":  "สันทรายน้อย",
		"dstDetailAddress": "example detail address",
		"dstName":          "น้ำพริกแม่อำพร",
		"dstPhone":         "0123456789",
		"dstPostalCode":    "50210",
		"dstProvinceName":  "เชียงใหม่",
		"expressCategory":  "1",
		"insured":          "0",
		"mchId":            "merchant",
		"nonceStr":         "nonce",
		"srcDistrictName":  "ขี้เหล็ก",
		"srcCityName":      "เมืองอุบลราชธานี",
		"srcDetailAddress": "example detail address",
		"srcName":          "หอมรวม  create order test name",
		"srcPhone":         "0123456789",
		"srcPostalCode":    "34000",
		"srcProvinceName":  "อุบลราชธานี",
		"weight":           "1000",
		"sign":             "signature",
	}).Return(FlashCreateOrderAPIResponse{}, errors.New("error"))

	trackingNo, err := t.service.CreateOrder(deliverypartnerconnectionlib.Order{
		WeightInGram: 1000,
		IsCOD:        true,
		Sender: deliverypartnerconnectionlib.OrderAddress{
			Name:          "หอมรวม  create order test name",
			AddressDetail: "example detail address",
			SubDistrict:   "ขี้เหล็ก",
			District:      "เมืองอุบลราชธานี",
			Province:      "อุบลราชธานี",
			Phone:         "0123456789",
			PostalCode:    "34000",
		},
		Receiver: deliverypartnerconnectionlib.OrderAddress{
			Name:          "น้ำพริกแม่อำพร",
			AddressDetail: "example detail address",
			SubDistrict:   "สันทรายน้อย",
			District:      "สันทราย",
			Province:      "เชียงใหม่",
			Phone:         "0123456789",
			PostalCode:    "50210",
		},
	})

	t.Empty(trackingNo)
	t.EqualError(err, "failed to create order with flash: error")
}

func (t *FlashTestSuite) TestGivenNonCODOrderIsUpdating_WhenUpdateOrder_ThenUpdateSuccess() {
	t.mFlashUpdateShipmentInfo.EXPECT().PostForm(
		"/open/v1/orders/modify",
		map[string]string{
			"mchId":            "merchant",
			"nonceStr":         "nonce",
			"sign":             "signature",
			"pno":              "trackingNo",
			"dstCityName":      "สันทราย",
			"dstDistrictName":  "สันทรายน้อย",
			"dstDetailAddress": "example detail address",
			"dstName":          "น้ำพริกแม่อำพร",
			"dstPhone":         "0123456789",
			"dstPostalCode":    "50210",
			"dstProvinceName":  "เชียงใหม่",
			"srcDistrictName":  "ขี้เหล็ก",
			"srcCityName":      "เมืองอุบลราชธานี",
			"srcDetailAddress": "example detail address",
			"srcName":          "หอมรวม  create order test name",
			"srcPhone":         "0123456789",
			"srcPostalCode":    "34000",
			"srcProvinceName":  "อุบลราชธานี",
			"weight":           "1000",
			"insured":          "0",
			"codEnabled":       "0",
		},
	).Return(FlashOrderUpdateAPIResponse{}, nil)

	err := t.service.UpdateOrder("trackingNo", deliverypartnerconnectionlib.Order{
		WeightInGram: 1000,
		IsCOD:        false,
		Sender: deliverypartnerconnectionlib.OrderAddress{
			Name:          "หอมรวม  create order test name",
			AddressDetail: "example detail address",
			SubDistrict:   "ขี้เหล็ก",
			District:      "เมืองอุบลราชธานี",
			Province:      "อุบลราชธานี",
			Phone:         "0123456789",
			PostalCode:    "34000",
		},
		Receiver: deliverypartnerconnectionlib.OrderAddress{
			Name:          "น้ำพริกแม่อำพร",
			AddressDetail: "example detail address",
			SubDistrict:   "สันทรายน้อย",
			District:      "สันทราย",
			Province:      "เชียงใหม่",
			Phone:         "0123456789",
			PostalCode:    "50210",
		},
	})

	t.NoError(err)
}

func (t *FlashTestSuite) TestGivenFlashOrderIsExist_WhenDeleteOrder_ThenReturnSuccess() {
	t.mFlashDeleteOrder.EXPECT().PostForm("/open/v1/orders/trackingNo/cancel", map[string]string{
		"mchId":    "merchant",
		"nonceStr": "nonce",
		"sign":     "signature",
	}).Return(FlashOrderDeleteAPIResponse{}, nil)

	err := t.service.DeleteOrder("trackingNo")
	t.NoError(err)
}

func (t *FlashTestSuite) TestGivenFlashAPIIsBroken_WhenUpdateOrder_ThenReturnError() {
	t.mFlashUpdateShipmentInfo.EXPECT().PostForm(
		gomock.Any(), gomock.Any()).Return(FlashOrderUpdateAPIResponse{}, errors.New("error"))

	err := t.service.UpdateOrder("trackingNo", deliverypartnerconnectionlib.Order{})
	t.Error(err)
}

func (t *FlashTestSuite) TestGivenFlashAPIIsBroken_WhenDeleteOrder_ThenReturnError() {
	t.mFlashDeleteOrder.EXPECT().PostForm(
		gomock.Any(), gomock.Any()).Return(FlashOrderDeleteAPIResponse{}, errors.New("error"))

	err := t.service.DeleteOrder("trackingNo")
	t.Error(err)
}
