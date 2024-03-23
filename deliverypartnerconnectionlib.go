package deliverypartnerconnectionlib

type deliveryPartnerConnectionLib struct {
	partnerCreateOrderAdaptor map[string]OrderCreator
	partnerUpdateOrderAdaptor map[string]OrderUpdator
}

func New(partnerCreateOrderAdaptor map[string]OrderCreator, partnerUpdateOrderAdaptor map[string]OrderUpdator) *deliveryPartnerConnectionLib {
	return &deliveryPartnerConnectionLib{
		partnerCreateOrderAdaptor: partnerCreateOrderAdaptor,
		partnerUpdateOrderAdaptor: partnerUpdateOrderAdaptor,
	}
}

func (c *deliveryPartnerConnectionLib) CreateOrder(partner string, order Order) (string, error) {
	orderRefID, err := c.partnerCreateOrderAdaptor[partner].CreateOrder(order)
	return orderRefID, err
}

func (c *deliveryPartnerConnectionLib) UpdateOrder(partner, trackingNo string, order Order) error {
	err := c.partnerUpdateOrderAdaptor[partner].UpdateOrder(trackingNo, order)
	return err
}
