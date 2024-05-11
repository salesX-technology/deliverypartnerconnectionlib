package dhl

type DHLCreateOrderAPIResponse struct {
	ManifestResponse ManifestResponse `json:"manifestResponse"`
}

type DHLHookOrderAPIResponse struct {
	TrackItemResponse struct {
		Hdr struct {
			MessageType     string `json:"messageType"`
			MessageDateTime string `json:"messageDateTime"`
			MessageVersion  string `json:"messageVersion"`
		} `json:"hdr"`
		Bd struct {
			ShipmentItems []struct {
				MasterShipmentID string      `json:"masterShipmentID"`
				ShipmentID       string      `json:"shipmentID"`
				TrackingID       string      `json:"trackingID"`
				DeliveryImage    interface{} `json:"deliveryImage"`
				OrderNumber      interface{} `json:"orderNumber"`
				HandoverID       interface{} `json:"handoverID"`
				ShippingService  struct {
					ProductCode string `json:"productCode"`
					ProductName string `json:"productName"`
				} `json:"shippingService"`
				ConsigneeAddress struct {
					Country string `json:"country"`
				} `json:"consigneeAddress"`
				Weight            string      `json:"weight"`
				DimensionalWeight interface{} `json:"dimensionalWeight"`
				WeightUnit        string      `json:"weightUnit"`
				Events            []struct {
					Status      string `json:"status"`
					Description string `json:"description"`
					DateTime    string `json:"dateTime"`
					Timezone    string `json:"timezone"`
					Address     struct {
						City     string `json:"city"`
						PostCode string `json:"postCode"`
						State    string `json:"state,omitempty"`
						Country  string `json:"country"`
					} `json:"address"`
				} `json:"events"`
				Dimensions struct {
					Length interface{} `json:"length"`
					Width  interface{} `json:"width"`
					Height interface{} `json:"height"`
				} `json:"dimensions"`
				DimensionsUnit interface{} `json:"dimensionsUnit"`
			} `json:"shipmentItems"`
			ResponseStatus struct {
				Code           string `json:"code"`
				Message        string `json:"message"`
				MessageDetails []struct {
					MessageDetail string `json:"messageDetail"`
				} `json:"messageDetails"`
			} `json:"responseStatus"`
		} `json:"bd"`
	} `json:"trackItemResponse"`
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
