package deliverypartnerconnectionlib

type Order struct {
	ID           string
	WeightInGram int
	IsCOD        bool
	Sender       OrderAddress
	Receiver     OrderAddress
	CODValue     float64
	TotalValue   float64
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
