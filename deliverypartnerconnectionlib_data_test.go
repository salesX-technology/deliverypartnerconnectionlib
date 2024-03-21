package courierx

var aValidOrder = Order{
	WeightInGram: 1000,
	IsCOD:        false,
	Sender: OrderAddress{
		Name:          "sender_name",
		AddressDetail: "sender_address",
		District:      "sender_district",
		Province:      "sender_province",
		Phone:         "sender_phone",
		PostalCode:    "sender_postal_code",
	},
	Receiver: OrderAddress{
		Name:          "receiver_name",
		AddressDetail: "receiver_address",
		District:      "receiver_district",
		Province:      "receiver_province",
		Phone:         "receiver_phone",
		PostalCode:    "receiver_postal_code",
	},
}
