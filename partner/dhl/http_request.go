package dhl

type DHLCreateOrderAPIRequest struct {
	ManifestRequest ManifestRequest `json:"manifestRequest"`
}

type ManifestRequest struct {
	HDR HDR `json:"hdr"`
	BD  BD  `json:"bd"`
}

type HDR struct {
	MessageType     string `json:"messageType"`
	MessageDateTime string `json:"messageDateTime"`
	MessageVersion  string `json:"messageVersion"`
	AccessToken     string `json:"accessToken"`
}

type BD struct {
	PickupAccountID string         `json:"pickupAccountId"`
	SoldToAccountID string         `json:"soldToAccountId"`
	PickupDateTime  string         `json:"pickupDateTime"`
	HandoverMethod  int            `json:"handoverMethod"`
	PickupAddress   DHLADdress     `json:"pickupAddress"`
	SipperAddress   DHLADdress     `json:"shipperAddress"`
	ShipmentItems   []ShipmentItem `json:"shipmentItems"`
}

type ShipmentItem struct {
	Currency         string     `json:"currency"`
	TotalWeight      int        `json:"totalWeight"`
	TotalWeightUOM   string     `json:"totalWeightUOM"`
	ShipmentID       string     `json:"shipmentID"`
	ProductCode      string     `json:"productCode"`
	ConsigneeAddress DHLADdress `json:"consigneeAddress"`
}

type ShipmentPiece struct {
	PieceID int `json:"pieceID"`
}

type DHLADdress struct {
	Name     string `json:"name"`
	Address1 string `json:"address1"`
	Country  string `json:"country"`
	State    string `json:"state"`
	District string `json:"district"`
	PostCode string `json:"postCode"`
}

type DHLAuthenticationAPIRequest struct {
}

type DHLDeleteOrderAPIRequest struct {
	DeleteShipmentReq DHLDeleteOrderAPIRequestDeleteShipmentRequest `json:"deleteShipmentReq"`
}

type DHLDeleteOrderAPIRequestDeleteShipmentRequest struct {
	HDR DHLDeleteOrderAPIRequestHDR `json:"hdr"`
	BD  DHLDeleteOrderAPIRequestBD  `json:"bd"`
}

type DHLDeleteOrderAPIRequestHDR struct {
	MessageType     string `json:"messageType"`
	MessageDateTime string `json:"messageDateTime"`
	AccessToken     string `json:"accessToken"`
	MessageVersion  string `json:"messageVersion"`
}

type DHLDeleteOrderAPIRequestBD struct {
	PickupAccountID string                                 `json:"pickupAccountId"`
	SoldToAccountID string                                 `json:"soldToAccountId"`
	ShipmentItems   []DHLDeleteOrderAPIRequestShipmentItem `json:"shipmentItems"`
}

type DHLDeleteOrderAPIRequestShipmentItem struct {
	ShipmentID string `json:"shipmentID"`
}
