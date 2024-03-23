package dhl

import (
	"time"

	"github.com/salesX-technology/deliverypartnerconnectionlib"
)

var (
	_ deliverypartnerconnectionlib.OrderCreator = (*dhlService)(nil)
	_ deliverypartnerconnectionlib.OrderUpdator = (*dhlService)(nil)
)

type dhlService struct {
	authorizer         Authenticator
	dhlOrderCreatorAPI DHLOrderCreatorAPI
	DHLAPIConfig       DHLAPIConfig
	nowFunc            func() time.Time
}

type DHLServiceOption func(*dhlService)

type DHLAPIConfig struct {
	PickupAccountID string
	SoldToAccountID string
}

func NewDHLService(
	authorizer Authenticator,
	dhlOrderCreatorAPI DHLOrderCreatorAPI,
	dhlAPIConfig DHLAPIConfig,
	options ...DHLServiceOption,
) *dhlService {
	svc := &dhlService{
		authorizer:         authorizer,
		dhlOrderCreatorAPI: dhlOrderCreatorAPI,
		DHLAPIConfig:       dhlAPIConfig,
		nowFunc: func() time.Time {
			return time.Now().Local()
		},
	}

	for _, option := range options {
		option(svc)
	}

	return svc
}

func WithNowFunc(nowFunc func() time.Time) DHLServiceOption {
	return func(svc *dhlService) {
		svc.nowFunc = nowFunc
	}
}

func (f *dhlService) CreateOrder(order deliverypartnerconnectionlib.Order) (string, error) {
	accessToken, _ := f.authorizer.Authenticate()

	orderDateTime := f.nowFunc().Format("2006-01-02T15:04:05-07:00")
	resp, _ := f.dhlOrderCreatorAPI.Post(map[string]string{
		"Content-Type": "application/json",
	}, DHLCreateOrderAPIRequest{
		ManifestRequest: ManifestRequest{
			HDR: HDR{
				MessageType:     "SHIPMENT",
				MessageDateTime: orderDateTime,
				MessageVersion:  "1.0",
				AccessToken:     accessToken,
			},
			BD: BD{
				PickupAccountID: f.DHLAPIConfig.PickupAccountID,
				SoldToAccountID: f.DHLAPIConfig.SoldToAccountID,
				HandoverMethod:  handoverMethod,
				PickupDateTime:  orderDateTime,
				PickupAddress: DHLADdress{
					Name:     order.Sender.Name,
					Address1: order.Sender.AddressDetail,
					Country:  "TH",
					State:    order.Sender.Province,
					District: order.Sender.District,
					PostCode: order.Sender.PostalCode,
				},
				SipperAddress: DHLADdress{
					Name:     order.Receiver.Name,
					Address1: order.Receiver.AddressDetail,
					Country:  "TH",
					State:    order.Receiver.Province,
					District: order.Receiver.District,
					PostCode: order.Receiver.PostalCode,
				},
				ShipmentItems: []ShipmentItem{
					{
						Currency:       "THB",
						TotalWeight:    order.WeightInGram,
						TotalWeightUOM: "g",
						ShipmentID:     "THHSU" + order.ID,
						ProductCode:    "PDO",
						ConsigneeAddress: DHLADdress{
							Name:     order.Receiver.Name,
							Address1: order.Receiver.AddressDetail,
							Country:  "TH",
							State:    order.Receiver.Province,
							District: order.Receiver.District,
							PostCode: order.Receiver.PostalCode,
						},
					},
				},
			},
		},
	})

	return resp.ManifestResponse.BD.ShipmentItems[0].DeliveryConfirmationNo, nil
}

func (f *dhlService) UpdateOrder(trackingNo string, order deliverypartnerconnectionlib.Order) error {
	panic("unimplemented")
}
