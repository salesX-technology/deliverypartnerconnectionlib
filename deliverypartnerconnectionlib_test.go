package deliverypartnerconnectionlib

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type DeliveryPartnerConnectionTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	mFlashOrderCreator  *MockOrderCreator
	mShopeeOrderCreator *MockOrderCreator
	mDHLOrderCreator    *MockOrderCreator

	mFlashOrderUpdator *MockOrderUpdator

	mAnyOrderDeleter *MockOrderDeleter

	service *deliveryPartnerConnectionLib
}

func (t *DeliveryPartnerConnectionTestSuite) SetupTest() {
	t.ctrl = gomock.NewController(t.T())
	t.mFlashOrderCreator = NewMockOrderCreator(t.ctrl)
	t.mShopeeOrderCreator = NewMockOrderCreator(t.ctrl)
	t.mDHLOrderCreator = NewMockOrderCreator(t.ctrl)
	t.mFlashOrderUpdator = NewMockOrderUpdator(t.ctrl)
	t.mAnyOrderDeleter = NewMockOrderDeleter(t.ctrl)

	t.service = New(map[string]OrderCreator{
		"FLASH":  t.mFlashOrderCreator,
		"SHOPEE": t.mShopeeOrderCreator,
		"DHL":    t.mDHLOrderCreator,
	}, map[string]OrderUpdator{
		"FLASH": t.mFlashOrderUpdator,
	},
		map[string]OrderDeleter{
			"ANY": t.mAnyOrderDeleter,
		},
	)
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, new(DeliveryPartnerConnectionTestSuite))
}

func (t *DeliveryPartnerConnectionTestSuite) TestGivenFlashOrderIsCreating_WhenCreateOrder_ThenCallAdaptorFlashCreateOrderAndReturnSuccess() {
	t.mFlashOrderCreator.EXPECT().CreateOrder(aValidOrder).Return("refID", nil)

	orderRefID, err := t.service.CreateOrder("FLASH", aValidOrder)

	t.Equal("refID", orderRefID)
	t.Nil(err)
}

func (t *DeliveryPartnerConnectionTestSuite) TestGivenShopeeOrderIsCreating_WhenCreateOrder_ThenCallAdaptorShopeeCreateOrderAndReturnSuccess() {
	t.mShopeeOrderCreator.EXPECT().CreateOrder(aValidOrder).Return("refID", nil)

	orderRefID, err := t.service.CreateOrder("SHOPEE", aValidOrder)

	t.Equal("refID", orderRefID)
	t.Nil(err)
}

func (t *DeliveryPartnerConnectionTestSuite) TestGivenDHLOrderIsCreating_WhenCreateOrder_ThenCallAdaptorDHLCreateOrderAndReturnSuccess() {
	t.mDHLOrderCreator.EXPECT().CreateOrder(aValidOrder).Return("refID", nil)

	orderRefID, err := t.service.CreateOrder("DHL", aValidOrder)

	t.Equal("refID", orderRefID)
	t.Nil(err)
}

func (t *DeliveryPartnerConnectionTestSuite) TestGivenFlashOrderIsUpdating_WhenUpdateOrder_ThenCallAdaptorFlashUpdateOrderAndReturnSuccess() {
	t.mFlashOrderUpdator.EXPECT().UpdateOrder("trackingNo", aValidOrder).Return(nil)
	err := t.service.UpdateOrder("FLASH", "trackingNo", aValidOrder)
	t.Nil(err)
}

func (t *DeliveryPartnerConnectionTestSuite) TestGivenShopeeOrderIsUpdating_WhenUpdateOrder_ThenReturnError() {
	t.mFlashOrderUpdator.EXPECT().UpdateOrder("trackingNo", aValidOrder).Return(errors.New("error"))
	err := t.service.UpdateOrder("FLASH", "trackingNo", aValidOrder)
	t.Error(err)
}

func (t *DeliveryPartnerConnectionTestSuite) TestGivenAnyOrderIsDeleting_WhenDeleteOrder_ThenCallAdaptorAnyDeleteOrderAndReturnSuccess() {
	t.mAnyOrderDeleter.EXPECT().DeleteOrder("trackingNo").Return(nil)
	err := t.service.DeleteOrder("ANY", "trackingNo")
	t.Nil(err)
}

func (t *DeliveryPartnerConnectionTestSuite) TestGivenAnyOrderIsDeleting_WhenDeleteOrder_ThenCallAdaptorAnyDeleteOrderAndReturnError() {
	t.mAnyOrderDeleter.EXPECT().DeleteOrder("trackingNo").Return(errors.New("error"))
	err := t.service.DeleteOrder("ANY", "trackingNo")
	t.Error(err)
}
