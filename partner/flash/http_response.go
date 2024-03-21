package flash

type FlashCreateOrderAPIResponse struct {
	Data FlasFlashCreateOrderAPIResponseData `json:"data"`
}

type FlasFlashCreateOrderAPIResponseData struct {
	PNO string `json:"pno"`
}

type FlashOrderUpdateAPIResponse struct {
}
