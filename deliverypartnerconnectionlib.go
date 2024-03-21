package courierx

type courierXService struct {
	partnerCreateOrderAdaptor map[string]OrderCreator
}

func NewCourierXService(partnerCreateOrderAdaptor map[string]OrderCreator) *courierXService {
	return &courierXService{
		partnerCreateOrderAdaptor: partnerCreateOrderAdaptor,
	}
}

func (c *courierXService) CreateOrder(partner string, order Order) (string, error) {
	orderRefID, err := c.partnerCreateOrderAdaptor[partner].CreateOrder(order)
	return orderRefID, err
}
