package dhl

import (
	"testing"
	"time"

	"errors"

	"github.com/stretchr/testify/suite"
	gomock "go.uber.org/mock/gomock"
)

type DHLServiceTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	mAuthenticator      *MockAuthenticator
	mDHLOrderCreatorAPI *MockDHLOrderCreatorAPI
	mDHLOrderDeletorAPI *MockDHLOrderDeletorAPI
	mDHLOrderUpdatorAPI *MockDHLOrderUpdatorAPI

	service *dhlService
}

// setup test
func (t *DHLServiceTestSuite) SetupTest() {
	t.ctrl = gomock.NewController(t.T())
	t.mAuthenticator = NewMockAuthenticator(t.ctrl)
	t.mDHLOrderCreatorAPI = NewMockDHLOrderCreatorAPI(t.ctrl)
	t.mDHLOrderDeletorAPI = NewMockDHLOrderDeletorAPI(t.ctrl)
	t.mDHLOrderUpdatorAPI = NewMockDHLOrderUpdatorAPI(t.ctrl)

	t.service = NewDHLService(t.mAuthenticator,
		t.mDHLOrderCreatorAPI,
		t.mDHLOrderUpdatorAPI,
		t.mDHLOrderDeletorAPI,
		DHLAPIConfig{
			PickupAccountID: "PickupAccountID",
			SoldToAccountID: "SoldToAccountID",
		},
		WithNowFunc(func() time.Time {
			return time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)
		}))
}

func (t *DHLServiceTestSuite) TearDownTest() {
	t.ctrl.Finish()
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, new(DHLServiceTestSuite))
}

func (t *DHLServiceTestSuite) TestGivenOrderIsCreating_WhenCreateOrder_ThenReturnSuccess() {
	t.mAuthenticator.EXPECT().Authenticate().Return("accessToken", nil)
	t.mDHLOrderCreatorAPI.EXPECT().Post("/rest/v3/Shipment", map[string]string{
		"Content-Type": "application/json",
	}, DHLCreateOrderAPIRequest{
		ManifestRequest: ManifestRequest{
			HDR: HDR{
				MessageType:     "SHIPMENT",
				MessageDateTime: "2021-01-01T00:00:00+07:00",
				MessageVersion:  "1.0",
				AccessToken:     "accessToken",
			},
			BD: BD{
				PickupAccountID: "PickupAccountID",
				SoldToAccountID: "SoldToAccountID",
				HandoverMethod:  2,
				PickupDateTime:  "2021-01-01T00:00:00+07:00",
				PickupAddress: &DHLADdress{
					Name:     aValidNonCODOrder.Sender.Name,
					Address1: aValidNonCODOrder.Sender.AddressDetail,
					Country:  "TH",
					State:    aValidNonCODOrder.Sender.Province,
					District: aValidNonCODOrder.Sender.District,
					PostCode: aValidNonCODOrder.Sender.PostalCode,
				},
				SipperAddress: &DHLADdress{
					Name:     aValidNonCODOrder.Receiver.Name,
					Address1: aValidNonCODOrder.Receiver.AddressDetail,
					Country:  "TH",
					State:    aValidNonCODOrder.Receiver.Province,
					District: aValidNonCODOrder.Receiver.District,
					PostCode: aValidNonCODOrder.Receiver.PostalCode,
				},
				ShipmentItems: []ShipmentItem{
					{
						Currency:       "THB",
						TotalWeight:    1000,
						TotalWeightUOM: "g",
						ShipmentID:     aValidNonCODOrder.ID,
						ProductCode:    "PDO",
						ConsigneeAddress: &DHLADdress{
							Name:     aValidNonCODOrder.Receiver.Name,
							Address1: aValidNonCODOrder.Receiver.AddressDetail,
							Country:  "TH",
							State:    aValidNonCODOrder.Receiver.Province,
							District: aValidNonCODOrder.Receiver.District,
							PostCode: aValidNonCODOrder.Receiver.PostalCode,
						},
					},
				},
			},
		},
	}).Return(DHLCreateOrderAPIResponse{
		ManifestResponse: ManifestResponse{
			BD: DHLCreateOrderAPIResponseBD{
				ShipmentItems: []DHLCreateOrderAPIResponseBDShipmentItem{
					{
						DeliveryConfirmationNo: "DeliveryConfirmationNo",
					},
				},
			},
		},
	}, nil)

	trackingNo, err := t.service.CreateOrder(aValidNonCODOrder)
	t.Equal(aValidNonCODOrder.ID, trackingNo)
	t.Nil(err)
}

func (t *DHLServiceTestSuite) TestGivenOrderIsDeleting_WhenDeleteOrder_ThenReturnSuccess() {
	t.mAuthenticator.EXPECT().Authenticate().Return("accessToken", nil)

	t.mDHLOrderDeletorAPI.EXPECT().Post(
		"/rest/v2/Label/Delete",
		map[string]string{
			"Content-Type": "application/json",
		}, DHLDeleteOrderAPIRequest{
			DeleteShipmentReq: DHLDeleteOrderAPIRequestDeleteShipmentRequest{
				HDR: HDR{
					MessageType:     "DELETESHIPMENT",
					MessageDateTime: "2021-01-01T00:00:00+07:00",
					AccessToken:     "accessToken",
					MessageVersion:  "1.0",
				},
				BD: BD{
					SoldToAccountID: "SoldToAccountID",
					PickupAccountID: "PickupAccountID",
					ShipmentItems: []ShipmentItem{
						{
							ShipmentID: "trackingNo",
						},
					},
				},
			},
		}).Return(DHLDeleteOrderAPIResponse{}, nil)

	err := t.service.DeleteOrder("trackingNo")
	t.Nil(err)
}

func (t *DHLServiceTestSuite) TestGivenOrderIsCreated_WhenUpdateOrder_ThenReturnSuccess() {
	t.mAuthenticator.EXPECT().Authenticate().Return("accessToken", nil)
	t.mDHLOrderUpdatorAPI.EXPECT().Post("/rest/v2/Label/Edit", map[string]string{
		"Content-Type": "application/json",
	}, DHLUpdateOrderAPIRequest{
		LabelRequest: LabelRequest{
			HDR: HDR{
				MessageType:     "EDITSHIPMENT",
				MessageDateTime: "2021-01-01T00:00:00+07:00",
				MessageVersion:  "1.4",
				AccessToken:     "accessToken",
			},
			BD: BD{
				PickupAccountID: "PickupAccountID",
				SoldToAccountID: "SoldToAccountID",
				HandoverMethod:  2,
				PickupDateTime:  "2021-01-01T00:00:00+07:00",
				PickupAddress: &DHLADdress{
					Name:     aValidNonCODOrder.Sender.Name,
					Address1: aValidNonCODOrder.Sender.AddressDetail,
					Country:  "TH",
					State:    aValidNonCODOrder.Sender.Province,
					District: aValidNonCODOrder.Sender.District,
					PostCode: aValidNonCODOrder.Sender.PostalCode,
				},
				SipperAddress: &DHLADdress{
					Name:     aValidNonCODOrder.Receiver.Name,
					Address1: aValidNonCODOrder.Receiver.AddressDetail,
					Country:  "TH",
					State:    aValidNonCODOrder.Receiver.Province,
					District: aValidNonCODOrder.Receiver.District,
					PostCode: aValidNonCODOrder.Receiver.PostalCode,
				},
				ShipmentItems: []ShipmentItem{
					{
						Currency:       "THB",
						TotalWeight:    1000,
						TotalWeightUOM: "g",
						ShipmentID:     "trackingNo",
						ProductCode:    "PDO",
						ConsigneeAddress: &DHLADdress{
							Name:     aValidNonCODOrder.Receiver.Name,
							Address1: aValidNonCODOrder.Receiver.AddressDetail,
							Country:  "TH",
							State:    aValidNonCODOrder.Receiver.Province,
							District: aValidNonCODOrder.Receiver.District,
							PostCode: aValidNonCODOrder.Receiver.PostalCode,
						},
					},
				},
				Label: &Label{
					PageSize: "400x600",
					Format:   "PDF",
					Layout:   "1x1",
				},
			},
		},
	}).Return(DHLUpdateOrderAPIResponse{}, nil)

	err := t.service.UpdateOrder("trackingNo", aValidNonCODOrder)
	t.Nil(err)
}

func (t *DHLServiceTestSuite) TestGivenAuthenticatorIsBroken_WhenCreateOrder_ThenReturnError() {
	someError := errors.New("some error")
	t.mAuthenticator.EXPECT().Authenticate().Return("", someError)

	_, err := t.service.CreateOrder(aValidNonCODOrder)
	t.Equal(someError, err)
}

func (t *DHLServiceTestSuite) TestGivenAuthenticatorIsBroken_WhenUpdateOrder_ThenReturnError() {
	someError := errors.New("some error")
	t.mAuthenticator.EXPECT().Authenticate().Return("", someError)

	err := t.service.UpdateOrder("tracking", aValidNonCODOrder)
	t.Equal(someError, err)
}

func (t *DHLServiceTestSuite) TestGivenAuthenticatorIsBroken_WhenDeleteOrder_ThenReturnError() {
	someError := errors.New("some error")
	t.mAuthenticator.EXPECT().Authenticate().Return("", someError)

	err := t.service.DeleteOrder("tracking")
	t.Equal(someError, err)
}

func (t *DHLServiceTestSuite) TestGivenDHLAPIIsBroken_WhenCreateOrder_ThenReturnError() {
	someError := errors.New("some error")
	t.mAuthenticator.EXPECT().Authenticate().Return("token", nil)
	t.mDHLOrderCreatorAPI.EXPECT().Post(gomock.Any(), gomock.Any(), gomock.Any()).Return(DHLCreateOrderAPIResponse{}, errors.New("some error"))

	_, err := t.service.CreateOrder(aValidNonCODOrder)
	t.Equal(someError, err)
}

func (t *DHLServiceTestSuite) TestGivenDHLAPIIsBroken_WhenUpdateOrder_ThenReturnError() {
	someError := errors.New("some error")
	t.mAuthenticator.EXPECT().Authenticate().Return("token", nil)
	t.mDHLOrderUpdatorAPI.EXPECT().Post(gomock.Any(), gomock.Any(), gomock.Any()).Return(DHLUpdateOrderAPIResponse{}, errors.New("some error"))

	err := t.service.UpdateOrder("tracking", aValidNonCODOrder)
	t.Equal(someError, err)
}

func (t *DHLServiceTestSuite) TestGivenDHLAPIIsBroken_WhenDeleteOrder_ThenReturnError() {
	someError := errors.New("some error")
	t.mAuthenticator.EXPECT().Authenticate().Return("token", nil)
	t.mDHLOrderDeletorAPI.EXPECT().Post(gomock.Any(), gomock.Any(), gomock.Any()).Return(DHLDeleteOrderAPIResponse{}, errors.New("some error"))

	err := t.service.DeleteOrder("tracking")
	t.Equal(someError, err)
}
