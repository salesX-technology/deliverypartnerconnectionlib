package deliverypartnerconnectionlib

type deliveryPartnerConnectionLib struct {
	partnerCreateOrderAdaptor map[string]OrderCreator
	partnerUpdateOrderAdaptor map[string]OrderUpdator
	partnerDeleteOrderAdaptor map[string]OrderDeleter
}

func New(
	partnerCreateOrderAdaptor map[string]OrderCreator,
	partnerUpdateOrderAdaptor map[string]OrderUpdator,
	partnerDeleteOrderAdaptor map[string]OrderDeleter,
) *deliveryPartnerConnectionLib {
	return &deliveryPartnerConnectionLib{
		partnerCreateOrderAdaptor: partnerCreateOrderAdaptor,
		partnerUpdateOrderAdaptor: partnerUpdateOrderAdaptor,
		partnerDeleteOrderAdaptor: partnerDeleteOrderAdaptor,
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

func (c *deliveryPartnerConnectionLib) DeleteOrder(partner, trackingNo string) error {
	err := c.partnerDeleteOrderAdaptor[partner].DeleteOrder(trackingNo)
	return err
}
