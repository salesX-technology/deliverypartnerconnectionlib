package deliverypartnerconnectionlib

type Order struct {
	ID           string
	WeightInGram int
	IsCOD        bool
	Sender       OrderAddress
	Receiver     OrderAddress
	CODValue     float64
	TotalValue   float64
	SubItemTypes []*SubItemType
}

type OrderAddress struct {
	Name          string
	AddressDetail string
	SubDistrict   string
	District      string
	Province      string
	Phone         string
	PostalCode    string
}

type SubItemType struct {
	ItemName       string `json:"itemName"`       //max 200
	ItemWeightSize string `json:"itemWeightSize"` //max 128
	ItemColor      string `json:"itemColor"`      //max 128
	ItemQuantity   int    `json:"itemQuantity"`   //min 1 max 999
}
