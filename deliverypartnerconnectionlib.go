package deliverypartnerconnectionlib

type DeliveryPartnerConnectionLib struct {
	partnerCreateOrderAdaptor map[string]OrderCreator
	partnerUpdateOrderAdaptor map[string]OrderUpdator
	partnerDeleteOrderAdaptor map[string]OrderDeleter
}

func New(
	partnerCreateOrderAdaptor map[string]OrderCreator,
	partnerUpdateOrderAdaptor map[string]OrderUpdator,
	partnerDeleteOrderAdaptor map[string]OrderDeleter,
) *DeliveryPartnerConnectionLib {
	return &DeliveryPartnerConnectionLib{
		partnerCreateOrderAdaptor: partnerCreateOrderAdaptor,
		partnerUpdateOrderAdaptor: partnerUpdateOrderAdaptor,
		partnerDeleteOrderAdaptor: partnerDeleteOrderAdaptor,
	}
}

func (c *DeliveryPartnerConnectionLib) CreateOrder(partner string, order Order) (map[string]interface{}, error) {
	orderRefID, err := c.partnerCreateOrderAdaptor[partner].CreateOrder(order)
	return orderRefID, err
}

func (c *DeliveryPartnerConnectionLib) UpdateOrder(partner, trackingNo string, order Order) error {
	err := c.partnerUpdateOrderAdaptor[partner].UpdateOrder(trackingNo, order)
	return err
}

func (c *DeliveryPartnerConnectionLib) DeleteOrder(partner, trackingNo string) error {
	err := c.partnerDeleteOrderAdaptor[partner].DeleteOrder(trackingNo)
	return err
}
