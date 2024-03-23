package flash

import (
	"fmt"

	"github.com/salesX-technology/deliverypartnerconnectionlib"
)

var (
	_ deliverypartnerconnectionlib.OrderCreator = (*flashService)(nil)
	_ deliverypartnerconnectionlib.OrderUpdator = (*flashService)(nil)
	_ deliverypartnerconnectionlib.OrderDeleter = (*flashService)(nil)
)

func NewFlashService(
	flashCreateOrderAPI FlashCreateOrderAPI,
	flashUpdateOrderAPI FlashUpdateShipmentInfo,
	flashDeleteOrderaPI FlashDeleteOrderAPI,
	secretKey string,
	merchantID string,
	options ...FlashServiceOption,
) *flashService {
	fs := &flashService{
		flashCreateOrderAPI: flashCreateOrderAPI,
		flashUpdateOrderAPI: flashUpdateOrderAPI,
		flashDeleteOrderaPI: flashDeleteOrderaPI,
		secretKey:           secretKey,
		merchantID:          merchantID,
		nonceGenerator:      generateNonceStr,
		signatureGenerator:  generateSignature,
	}

	for _, option := range options {
		option(fs)
	}

	return fs
}

type FlashServiceOption func(*flashService)

func WithNonceGenerator(ng nonceGeneratorFunc) FlashServiceOption {
	return func(fs *flashService) {
		fs.nonceGenerator = ng
	}
}

func WithSignatureGenerator(sg signatureGeneratorFunc) FlashServiceOption {
	return func(fs *flashService) {
		fs.signatureGenerator = sg
	}
}

type flashService struct {
	flashCreateOrderAPI FlashCreateOrderAPI
	flashUpdateOrderAPI FlashUpdateShipmentInfo
	flashDeleteOrderaPI FlashDeleteOrderAPI
	secretKey           string
	merchantID          string
	baseURL             string
	nonceGenerator      nonceGeneratorFunc
	signatureGenerator  signatureGeneratorFunc
}

type nonceGeneratorFunc func(int) string
type signatureGeneratorFunc func(map[string]string, string) string

func (f *flashService) CreateOrder(order deliverypartnerconnectionlib.Order) (string, error) {
	articleCategory := "99"
	expressCategory := "1"
	insured := "0"
	codEnabled := "0"
	if order.IsCOD {
		codEnabled = "1"
	}

	nonceStr := f.nonceGenerator(32)

	keyedOrderInfo := map[string]string{
		"articleCategory":  articleCategory,
		"codEnabled":       codEnabled,
		"dstDistrictName":  order.Receiver.SubDistrict,
		"dstCityName":      order.Receiver.District,
		"dstDetailAddress": order.Receiver.AddressDetail,
		"dstName":          order.Receiver.Name,
		"dstPhone":         order.Receiver.Phone,
		"dstPostalCode":    order.Receiver.PostalCode,
		"dstProvinceName":  order.Receiver.Province,
		"expressCategory":  expressCategory,
		"insured":          insured,
		"mchId":            f.merchantID,
		"nonceStr":         nonceStr,
		"srcDistrictName":  order.Sender.SubDistrict,
		"srcCityName":      order.Sender.District,
		"srcDetailAddress": order.Sender.AddressDetail,
		"srcName":          order.Sender.Name,
		"srcPhone":         order.Sender.Phone,
		"srcPostalCode":    order.Sender.PostalCode,
		"srcProvinceName":  order.Sender.Province,
		"weight":           fmt.Sprintf("%d", order.WeightInGram),
	}

	plainSignature := f.signatureGenerator(keyedOrderInfo, f.secretKey)

	keyedOrderInfo["sign"] = plainSignature

	response, err := f.flashCreateOrderAPI.PostForm("/open/v3/orders", keyedOrderInfo)
	if err != nil {
		return "", fmt.Errorf("failed to create order with flash: %s", err)
	}

	return response.Data.PNO, nil
}

func (f *flashService) UpdateOrder(trackingNo string, order deliverypartnerconnectionlib.Order) error {
	insured := "0"
	codEnabled := "0"
	if order.IsCOD {
		codEnabled = "1"
	}

	nonceStr := f.nonceGenerator(32)

	keyedOrderInfo := map[string]string{
		"mchId":            f.merchantID,
		"nonceStr":         nonceStr,
		"pno":              trackingNo,
		"dstDistrictName":  order.Receiver.SubDistrict,
		"dstCityName":      order.Receiver.District,
		"dstDetailAddress": order.Receiver.AddressDetail,
		"dstName":          order.Receiver.Name,
		"dstPhone":         order.Receiver.Phone,
		"dstPostalCode":    order.Receiver.PostalCode,
		"dstProvinceName":  order.Receiver.Province,
		"insured":          insured,
		"srcDistrictName":  order.Sender.SubDistrict,
		"srcCityName":      order.Sender.District,
		"srcDetailAddress": order.Sender.AddressDetail,
		"srcName":          order.Sender.Name,
		"srcPhone":         order.Sender.Phone,
		"srcPostalCode":    order.Sender.PostalCode,
		"srcProvinceName":  order.Sender.Province,
		"weight":           fmt.Sprintf("%d", order.WeightInGram),
		"codEnabled":       codEnabled,
	}

	plainSignature := f.signatureGenerator(keyedOrderInfo, f.secretKey)

	keyedOrderInfo["sign"] = plainSignature

	f.flashUpdateOrderAPI.PostForm(f.baseURL+"/open/v1/orders/modify", keyedOrderInfo)

	return nil
}

func (f *flashService) DeleteOrder(trackingNo string) error {
	nonceStr := f.nonceGenerator(32)
	deleteForm := map[string]string{
		"mchId":    f.merchantID,
		"nonceStr": nonceStr,
	}

	plainSignature := f.signatureGenerator(deleteForm, f.secretKey)

	deleteForm["sign"] = plainSignature

	_, _ = f.flashDeleteOrderaPI.PostForm("/open/v1/orders/"+trackingNo+"/cancel", deleteForm)

	return nil
}
