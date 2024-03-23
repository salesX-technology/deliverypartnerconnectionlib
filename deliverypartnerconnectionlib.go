package deliverypartnerconnectionlib

type deliveryPartnerConnectionLib struct {
	partnerCreateOrderAdaptor map[string]OrderCreator
}

func New(partnerCreateOrderAdaptor map[string]OrderCreator) *deliveryPartnerConnectionLib {
	return &deliveryPartnerConnectionLib{
		partnerCreateOrderAdaptor: partnerCreateOrderAdaptor,
	}
}

func (c *deliveryPartnerConnectionLib) CreateOrder(partner string, order Order) (string, error) {
	orderRefID, err := c.partnerCreateOrderAdaptor[partner].CreateOrder(order)
	return orderRefID, err
}
