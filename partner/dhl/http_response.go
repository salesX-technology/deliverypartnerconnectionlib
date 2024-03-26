package dhl

type DHLCreateOrderAPIResponse struct {
	ManifestResponse ManifestResponse `json:"manifestResponse"`
}

type ManifestResponse struct {
	BD DHLCreateOrderAPIResponseBD `json:"bd"`
}

type DHLCreateOrderAPIResponseBD struct {
	ShipmentItems []DHLCreateOrderAPIResponseBDShipmentItem `json:"shipmentItems"`
}

type DHLCreateOrderAPIResponseBDShipmentItem struct {
	DeliveryConfirmationNo string `json:"deliveryConfirmationNo"`
}

type DHLAuthenticationAPIResponse struct {
	AccessTokenResponse AccessTokenResponse `json:"accessTokenResponse"`
}

type AccessTokenResponse struct {
	Token string `json:"token"`
}

type DHLDeleteOrderAPIResponse struct {
}

type DHLUpdateOrderAPIResponse struct {
}
