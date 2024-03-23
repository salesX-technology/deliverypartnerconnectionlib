package dhl

import "github.com/salesX-technology/deliverypartnerconnectionlib"

var aValidNonCODOrder = deliverypartnerconnectionlib.Order{
	ID:           "order_id",
	WeightInGram: 1000,
	IsCOD:        false,
	Sender: deliverypartnerconnectionlib.OrderAddress{
		Name:          "sender_name",
		AddressDetail: "sender_address",
		SubDistrict:   "sender_sub_district",
		District:      "sender_district",
		Province:      "sender_province",
		Phone:         "sender_phone",
		PostalCode:    "sender_postal_code",
	},
	Receiver: deliverypartnerconnectionlib.OrderAddress{
		Name:          "receiver_name",
		AddressDetail: "receiver_address",
		SubDistrict:   "receiver_sub_district",
		District:      "receiver_district",
		Province:      "receiver_province",
		Phone:         "receiver_phone",
		PostalCode:    "receiver_postal_code",
	},
}
