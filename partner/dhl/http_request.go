package dhl

type DHLCreateOrderAPIRequest struct {
	ManifestRequest ManifestRequest `json:"manifestRequest"`
}

type DHLHookOrderAPIRequest struct {
	TrackItemRequest ManifestRequest `json:"trackItemRequest"`
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
	PickupAccountID         string         `json:"pickupAccountId,omitempty"`
	SoldToAccountID         string         `json:"soldToAccountId,omitempty"`
	PickupDateTime          string         `json:"pickupDateTime,omitempty"`
	HandoverMethod          int            `json:"handoverMethod,omitempty"`
	ShipmentContent         string         `json:"shipmentContent,omitempty"`
	PickupAddress           *DHLADdress    `json:"pickupAddress,omitempty"`
	SipperAddress           *DHLADdress    `json:"shipperAddress,omitempty"`
	ShipmentItems           []ShipmentItem `json:"shipmentItems,omitempty"`
	Label                   *Label         `json:"label,omitempty"`
	Epod                    string         `json:"ePODRequired,omitempty"`
	TrackingReferenceNumber []string       `json:"trackingReferenceNumber,omitempty"`
}

type ShipmentItem struct {
	Currency         string            `json:"currency,omitempty"`
	TotalWeight      int               `json:"totalWeight,omitempty"`
	TotalWeightUOM   string            `json:"totalWeightUOM,omitempty"`
	ShipmentID       string            `json:"shipmentID,omitempty"`
	ProductCode      string            `json:"productCode,omitempty"`
	CodValue         float64           `json:"codValue,omitempty"`
	TotalValue       float64           `json:"totalValue,omitempty"`
	ConsigneeAddress *DHLADdress       `json:"consigneeAddress,omitempty"`
	ShipmentContents []ShipmentContent `json:"shipmentContents,omitempty"`
	ReturnMode       string            `json:"returnMode,omitempty"`
}

type ShipmentContent struct {
	SkuNumber         string  `json:"skuNumber,omitempty"`
	Description       string  `json:"description,omitempty"`
	DescriptionImport string  `json:"descriptionImport,omitempty"`
	DescriptionExport string  `json:"descriptionExport,omitempty"`
	ItemValue         float64 `json:"itemValue,omitempty"`
	ItemQuantity      int     `json:"itemQuantity,omitempty"`
	GrossWeight       int     `json:"grossWeight,omitempty"`
	NetWeight         int     `json:"netWeight,omitempty"`
	WeightUOM         string  `json:"weightUOM,omitempty"`
	ContentIndicator  string  `json:"contentIndicator,omitempty"`
	CountryOfOrigin   string  `json:"countryOfOrigin,omitempty"`
	HsCode            string  `json:"hsCode,omitempty"`
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
	Phone    string `json:"phone"`
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
