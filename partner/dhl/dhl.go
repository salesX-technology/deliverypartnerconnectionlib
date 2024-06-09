package dhl

import (
	"fmt"
	"time"

	"github.com/salesX-technology/deliverypartnerconnectionlib"
)

var (
	_ deliverypartnerconnectionlib.OrderCreator = (*dhlService)(nil)
	_ deliverypartnerconnectionlib.OrderUpdator = (*dhlService)(nil)
	_ deliverypartnerconnectionlib.OrderDeleter = (*dhlService)(nil)
)

type dhlService struct {
	authorizer         Authenticator
	dhlOrderCreatorAPI DHLOrderCreatorAPI
	dhlOrderDeletorAPI DHLOrderDeletorAPI
	dhlOrderUpdatorAPI DHLOrderUpdatorAPI
	dhlHookOrderAPI    DHLHookOrderAPI
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
	dhlOrderUpdatorAPI DHLOrderUpdatorAPI,
	dhlOrderDeletorAPI DHLOrderDeletorAPI,
	dhlHookOrderAPI DHLHookOrderAPI,
	dhlAPIConfig DHLAPIConfig,
	options ...DHLServiceOption,
) *dhlService {
	svc := &dhlService{
		authorizer:         authorizer,
		dhlOrderCreatorAPI: dhlOrderCreatorAPI,
		dhlOrderDeletorAPI: dhlOrderDeletorAPI,
		dhlOrderUpdatorAPI: dhlOrderUpdatorAPI,
		dhlHookOrderAPI:    dhlHookOrderAPI,
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

func (f *dhlService) CreateOrder(order deliverypartnerconnectionlib.Order) (map[string]interface{}, error) {
	var resDHLOrderCreateOrder map[string]interface{}

	accessToken, err := f.authorizer.Authenticate()
	if err != nil {
		return resDHLOrderCreateOrder, err
	}

	orderDateTime := f.nowFunc().Format("2006-01-02T15:04:05-07:00")

	if order.IsCOD {
		_, err = f.dhlOrderCreatorAPI.Post(
			"/rest/v3/Shipment",
			map[string]string{
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
						PickupAddress: &DHLADdress{
							Name:     order.Sender.Name,
							Address1: order.Sender.AddressDetail,
							Country:  "TH",
							State:    order.Sender.Province,
							District: order.Sender.District,
							PostCode: order.Sender.PostalCode,
							Phone:    order.Sender.Phone,
						},
						SipperAddress: &DHLADdress{
							Name:     order.Receiver.Name,
							Address1: order.Receiver.AddressDetail,
							Country:  "TH",
							State:    order.Receiver.Province,
							District: order.Receiver.District,
							PostCode: order.Receiver.PostalCode,
							Phone:    order.Receiver.Phone,
						},
						ShipmentItems: []ShipmentItem{
							{
								Currency:       "THB",
								TotalWeight:    order.WeightInGram,
								TotalWeightUOM: "g",
								ShipmentID:     order.ID,
								ProductCode:    "PDO",
								CodValue:       order.CODValue,
								ReturnMode:     "02",
								ConsigneeAddress: &DHLADdress{
									Name:     order.Receiver.Name,
									Address1: order.Receiver.AddressDetail,
									Country:  "TH",
									State:    order.Receiver.Province,
									District: order.Receiver.District,
									PostCode: order.Receiver.PostalCode,
									Phone:    order.Receiver.Phone,
								},
							},
						},
					},
				},
			})

	} else {
		_, err = f.dhlOrderCreatorAPI.Post(
			"/rest/v3/Shipment",
			map[string]string{
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
						PickupAddress: &DHLADdress{
							Name:     order.Sender.Name,
							Address1: order.Sender.AddressDetail,
							Country:  "TH",
							State:    order.Sender.Province,
							District: order.Sender.District,
							PostCode: order.Sender.PostalCode,
							Phone:    order.Sender.Phone,
						},
						SipperAddress: &DHLADdress{
							Name:     order.Receiver.Name,
							Address1: order.Receiver.AddressDetail,
							Country:  "TH",
							State:    order.Receiver.Province,
							District: order.Receiver.District,
							PostCode: order.Receiver.PostalCode,
							Phone:    order.Receiver.Phone,
						},
						ShipmentItems: []ShipmentItem{
							{
								Currency:       "THB",
								TotalWeight:    order.WeightInGram,
								TotalWeightUOM: "g",
								ShipmentID:     order.ID,
								ProductCode:    "PDO",
								ReturnMode:     "02",
								ConsigneeAddress: &DHLADdress{
									Name:     order.Receiver.Name,
									Address1: order.Receiver.AddressDetail,
									Country:  "TH",
									State:    order.Receiver.Province,
									District: order.Receiver.District,
									PostCode: order.Receiver.PostalCode,
									Phone:    order.Receiver.Phone,
								},
							},
						},
					},
				},
			})

	}

	if err != nil {
		return resDHLOrderCreateOrder, err
	}

	resDHLOrderCreateOrder = map[string]interface{}{
		"trackingNo": order.ID,
	}

	return resDHLOrderCreateOrder, nil
}

func (f *dhlService) UpdateOrder(trackingNo string, order deliverypartnerconnectionlib.Order) error {
	accessToken, err := f.authorizer.Authenticate()
	if err != nil {
		return err
	}

	orderDateTime := f.nowFunc().Format("2006-01-02T15:04:05-07:00")
	_, err = f.dhlOrderUpdatorAPI.Post(
		"/rest/v2/Label/Edit",
		map[string]string{
			"Content-Type": "application/json",
		}, DHLUpdateOrderAPIRequest{
			LabelRequest: LabelRequest{
				HDR: HDR{
					MessageType:     "EDITSHIPMENT",
					MessageDateTime: orderDateTime,
					MessageVersion:  "1.4",
					AccessToken:     accessToken,
				},
				BD: BD{
					PickupAccountID: f.DHLAPIConfig.PickupAccountID,
					SoldToAccountID: f.DHLAPIConfig.SoldToAccountID,
					HandoverMethod:  handoverMethod,
					PickupDateTime:  orderDateTime,
					PickupAddress: &DHLADdress{
						Name:     order.Sender.Name,
						Address1: order.Sender.AddressDetail,
						Country:  "TH",
						State:    order.Sender.Province,
						District: order.Sender.District,
						PostCode: order.Sender.PostalCode,
						Phone:    order.Sender.Phone,
					},
					SipperAddress: &DHLADdress{
						Name:     order.Receiver.Name,
						Address1: order.Receiver.AddressDetail,
						Country:  "TH",
						State:    order.Receiver.Province,
						District: order.Receiver.District,
						PostCode: order.Receiver.PostalCode,
						Phone:    order.Receiver.Phone,
					},
					ShipmentItems: []ShipmentItem{
						{
							Currency:       "THB",
							TotalWeight:    order.WeightInGram,
							TotalWeightUOM: "g",
							ShipmentID:     trackingNo,
							ProductCode:    "PDO",
							ReturnMode:     "02",
							ConsigneeAddress: &DHLADdress{
								Name:     order.Receiver.Name,
								Address1: order.Receiver.AddressDetail,
								Country:  "TH",
								State:    order.Receiver.Province,
								District: order.Receiver.District,
								PostCode: order.Receiver.PostalCode,
								Phone:    order.Receiver.Phone,
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
		})

	if err != nil {
		return err
	}

	return nil
}

// DeleteOrder implements deliverypartnerconnectionlib.OrderDeleter.
func (f *dhlService) DeleteOrder(trackingNo string) error {
	accessToken, err := f.authorizer.Authenticate()
	if err != nil {
		return err
	}

	transactionDateTime := f.nowFunc().Format("2006-01-02T15:04:05-07:00")

	_, err = f.dhlOrderDeletorAPI.Post(
		"/rest/v2/Label/Delete",
		map[string]string{
			"Content-Type": "application/json",
		}, DHLDeleteOrderAPIRequest{
			DeleteShipmentReq: DHLDeleteOrderAPIRequestDeleteShipmentRequest{
				HDR: HDR{
					MessageType:     "DELETESHIPMENT",
					MessageDateTime: transactionDateTime,
					AccessToken:     accessToken,
					MessageVersion:  "1.0",
				},
				BD: BD{
					SoldToAccountID: f.DHLAPIConfig.SoldToAccountID,
					PickupAccountID: f.DHLAPIConfig.PickupAccountID,
					ShipmentItems: []ShipmentItem{
						{
							ShipmentID: trackingNo,
						},
					},
				},
			},
		})
	if err != nil {
		return err
	}

	return nil
}

func (f *dhlService) CreateReceived(order deliverypartnerconnectionlib.Order) (map[string]interface{}, error) {
	var resDHLOrderCreateOrder map[string]interface{}

	accessToken, err := f.authorizer.Authenticate()
	if err != nil {
		return resDHLOrderCreateOrder, err
	}

	orderDateTime := f.nowFunc().Format("2006-01-02T15:04:05-07:00")
	_, err = f.dhlOrderCreatorAPI.Post(
		"/rest/v3/Shipment",
		map[string]string{
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
					PickupAddress: &DHLADdress{
						Name:     order.Sender.Name,
						Address1: order.Sender.AddressDetail,
						Country:  "TH",
						State:    order.Sender.Province,
						District: order.Sender.District,
						PostCode: order.Sender.PostalCode,
						Phone:    order.Sender.Phone,
					},
					SipperAddress: &DHLADdress{
						Name:     order.Receiver.Name,
						Address1: order.Receiver.AddressDetail,
						Country:  "TH",
						State:    order.Receiver.Province,
						District: order.Receiver.District,
						PostCode: order.Receiver.PostalCode,
						Phone:    order.Receiver.Phone,
					},
					ShipmentItems: []ShipmentItem{
						{
							Currency:       "THB",
							TotalWeight:    order.WeightInGram,
							TotalWeightUOM: "g",
							ShipmentID:     "DHL_PUR_" + order.ID,
							ProductCode:    "PDO",
							ConsigneeAddress: &DHLADdress{
								Name:     order.Receiver.Name,
								Address1: order.Receiver.AddressDetail,
								Country:  "TH",
								State:    order.Receiver.Province,
								District: order.Receiver.District,
								PostCode: order.Receiver.PostalCode,
								Phone:    order.Receiver.Phone,
							},
						},
					},
				},
			},
		})

	if err != nil {
		return resDHLOrderCreateOrder, err
	}

	resDHLOrderCreateOrder = map[string]interface{}{
		"trackingNo": order.ID,
	}

	return resDHLOrderCreateOrder, nil
}

func (f *dhlService) HookOrder(trackingNoLists []string) (map[string]interface{}, error) {

	dhlURL := "/rest/v3/Tracking"

	accessToken, err := f.authorizer.Authenticate()
	if err != nil {
		return nil, err
	}

	res, err := f.dhlHookOrderAPI.PostHook(
		dhlURL,
		map[string]string{
			"Content-Type": "application/json",
		}, DHLHookOrderAPIRequest{
			TrackItemRequest: ManifestRequest{
				HDR: HDR{
					MessageType:     "TRACKITEM",
					MessageDateTime: time.Now().Format("2006-01-02T15:04:05-07:00"),
					MessageVersion:  "1.0",
					AccessToken:     accessToken,
				},
				BD: BD{
					PickupAccountID:         f.DHLAPIConfig.PickupAccountID,
					SoldToAccountID:         f.DHLAPIConfig.SoldToAccountID,
					Epod:                    "Y",
					TrackingReferenceNumber: trackingNoLists,
				},
			},
		})

	if err != nil {
		return nil, fmt.Errorf("failed to hook order: %w", err)
	}
	resMap := make(map[string]interface{})
	resMap["trackingNo"] = res.TrackItemResponse.Bd.ShipmentItems

	return resMap, nil
}
