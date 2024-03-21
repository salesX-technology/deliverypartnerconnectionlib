package flash

import "testing"

func TestGivenAllFieldsOfOrderData_WhenCreateSignature_And_Key_ReturnSignature(t *testing.T) {
	articleCategory := "99"
	codEnabled := "0"
	dstCityName := "สันทราย"
	dstDetailAddress := "example detail address"
	dstName := "น้ำพริกแม่อำพร"
	dstPhone := "0123456789"
	dstPostalCode := "50210"
	dstProvinceName := "เชียงใหม่"
	expressCategory := "1"
	insured := "0"
	merchantID := "AA2315"
	srcCityName := "เมืองอุบลราชธานี"
	srcDetailAddress := "example detail address"
	srcName := "หอมรวม  create order test name"
	srcPhone := "0123456789"
	srcPostalCode := "34000"
	srcProvinceName := "อุบลราชธานี"

	weight := "1000"
	key := "key"
	nonce := "nonce"
	keyedOrderInfo := map[string]string{
		"articleCategory":  articleCategory,
		"codEnabled":       codEnabled,
		"dstCityName":      dstCityName,
		"dstDetailAddress": dstDetailAddress,
		"dstName":          dstName,
		"dstPhone":         dstPhone,
		"dstPostalCode":    dstPostalCode,
		"dstProvinceName":  dstProvinceName,
		"expressCategory":  expressCategory,
		"insured":          insured,
		"mchId":            merchantID,
		"nonceStr":         nonce,
		"srcCityName":      srcCityName,
		"srcDetailAddress": srcDetailAddress,
		"srcName":          srcName,
		"srcPhone":         srcPhone,
		"srcPostalCode":    srcPostalCode,
		"srcProvinceName":  srcProvinceName,
		"weight":           weight,
	}

	expectedSignature := "E266747637E48BE2843FA673765C08419C5CE6638A6B4D2ADA7B9AE8101ABCFB"

	signature := generateSignature(keyedOrderInfo, key)

	if signature != expectedSignature {
		t.Errorf("Expected %s but got %s", expectedSignature, signature)
	}
}

func TestGivenDSTNameFieldsOfOrderDataIsEmpty_WhenCreateSignature_And_Key_ReturnSignature(t *testing.T) {
	dstName := ""

	articleCategory := "99"
	dstCityName := "สันทราย"
	dstDetailAddress := "example detail address"
	dstPhone := "0123456789"
	dstPostalCode := "50210"
	dstProvinceName := "เชียงใหม่"
	expressCategory := "1"
	insured := "0"
	merchantID := "AA2315"
	srcCityName := "เมืองอุบลราชธานี"
	srcDetailAddress := "example detail address"
	srcName := "หอมรวม  create order test name"
	srcPhone := "0123456789"
	srcPostalCode := "34000"
	srcProvinceName := "อุบลราชธานี"
	codEnabled := "0"

	weight := "1000"

	key := "key"
	nonce := "nonce"
	keyedOrderInfo := map[string]string{
		"articleCategory":  articleCategory,
		"codEnabled":       codEnabled,
		"dstCityName":      dstCityName,
		"dstDetailAddress": dstDetailAddress,
		"dstName":          dstName,
		"dstPhone":         dstPhone,
		"dstPostalCode":    dstPostalCode,
		"dstProvinceName":  dstProvinceName,
		"expressCategory":  expressCategory,
		"insured":          insured,
		"mchId":            merchantID,
		"nonceStr":         nonce,
		"srcCityName":      srcCityName,
		"srcDetailAddress": srcDetailAddress,
		"srcName":          srcName,
		"srcPhone":         srcPhone,
		"srcPostalCode":    srcPostalCode,
		"srcProvinceName":  srcProvinceName,
		"weight":           weight,
	}

	expectedSignature := "BF9F58D0FF15352E9996327F9620AE8E13DAABF029D38AA5965E653DACCB5B23"

	signature := generateSignature(keyedOrderInfo, key)

	if signature != expectedSignature {
		t.Errorf("Expected %s but got %s", expectedSignature, signature)
	}
}
