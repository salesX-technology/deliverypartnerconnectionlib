package courierx

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type CourierXTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	mFlashOrderCreator  *MockOrderCreator
	mShopeeOrderCreator *MockOrderCreator
	mDHLOrderCreator    *MockOrderCreator

	service *courierXService
}

func (t *CourierXTestSuite) SetupTest() {
	t.ctrl = gomock.NewController(t.T())
	t.mFlashOrderCreator = NewMockOrderCreator(t.ctrl)
	t.mShopeeOrderCreator = NewMockOrderCreator(t.ctrl)
	t.mDHLOrderCreator = NewMockOrderCreator(t.ctrl)

	t.service = NewCourierXService(map[string]OrderCreator{
		"FLASH":  t.mFlashOrderCreator,
		"SHOPEE": t.mShopeeOrderCreator,
		"DHL":    t.mDHLOrderCreator,
	})
}

func TestCourierXTestSuite(t *testing.T) {
	suite.Run(t, new(CourierXTestSuite))
}

func (t *CourierXTestSuite) TestGivenFlashOrderIsCreating_WhenCreateOrder_ThenCallAdaptorFlashCreateOrderAndReturnSuccess() {
	t.mFlashOrderCreator.EXPECT().CreateOrder(aValidOrder).Return("refID", nil)

	orderRefID, err := t.service.CreateOrder("FLASH", aValidOrder)

	t.Equal("refID", orderRefID)
	t.Nil(err)
}

func (t *CourierXTestSuite) TestGivenShopeeOrderIsCreating_WhenCreateOrder_ThenCallAdaptorShopeeCreateOrderAndReturnSuccess() {
	t.mShopeeOrderCreator.EXPECT().CreateOrder(aValidOrder).Return("refID", nil)

	orderRefID, err := t.service.CreateOrder("SHOPEE", aValidOrder)

	t.Equal("refID", orderRefID)
	t.Nil(err)
}

func (t *CourierXTestSuite) TestGivenDHLOrderIsCreating_WhenCreateOrder_ThenCallAdaptorDHLCreateOrderAndReturnSuccess() {
	t.mDHLOrderCreator.EXPECT().CreateOrder(aValidOrder).Return("refID", nil)

	orderRefID, err := t.service.CreateOrder("DHL", aValidOrder)

	t.Equal("refID", orderRefID)
	t.Nil(err)
}
