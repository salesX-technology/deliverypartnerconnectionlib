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
	PickupAccountID string         `json:"pickupAccountId,omitempty"`
	SoldToAccountID string         `json:"soldToAccountId,omitempty"`
	PickupDateTime  string         `json:"pickupDateTime,omitempty"`
	HandoverMethod  int            `json:"handoverMethod,omitempty"`
	PickupAddress   *DHLADdress    `json:"pickupAddress,omitempty"`
	SipperAddress   *DHLADdress    `json:"shipperAddress,omitempty"`
	ShipmentItems   []ShipmentItem `json:"shipmentItems,omitempty"`
	Label           *Label         `json:"label,omitempty"`
}

type ShipmentItem struct {
	Currency         string      `json:"currency,omitempty"`
	TotalWeight      int         `json:"totalWeight,omitempty"`
	TotalWeightUOM   string      `json:"totalWeightUOM,omitempty"`
	ShipmentID       string      `json:"shipmentID,omitempty"`
	ProductCode      string      `json:"productCode,omitempty"`
	ConsigneeAddress *DHLADdress `json:"consigneeAddress,omitempty"`
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
	HDR HDR `json:"hdr"`
	BD  BD  `json:"bd"`
}

type DHLDeleteOrderAPIRequestShipmentItem struct {
	ShipmentID string `json:"shipmentID"`
}

type DHLUpdateOrderAPIRequest struct {
	LabelRequest LabelRequest `json:"labelRequest"`
}

type LabelRequest struct {
	HDR HDR `json:"hdr"`
	BD  BD  `json:"bd"`
}

type Label struct {
	PageSize string `json:"pageSize"`
	Format   string `json:"format"`
	Layout   string `json:"layout"`
}
