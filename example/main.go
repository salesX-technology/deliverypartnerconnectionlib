package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/salesX-technology/deliverypartnerconnectionlib"
	"github.com/salesX-technology/deliverypartnerconnectionlib/httpclient"
	"github.com/salesX-technology/deliverypartnerconnectionlib/partner/flash"
	"github.com/salesX-technology/deliverypartnerconnectionlib/partner/shopee"
)

func main() {
	// shopeeCreateOrderExample()
	// fmt.Print("\nshopeeCreateOrderExample\n")
	// shopeeCreateOrderExample()
	// fmt.Print("\nflashCreateOrderExample\n")
	// flashCreateOrderExample()

	// z := time.Now().Unix()
	// r := rand.Int63()
	// x, _ := GenerateCheckSign(100190, "57e08ead78cf63721eed92911f2dfe8a1a1152ebc880877ceae96e406c16dbab", z, r, []byte("test"))
	// y := makeSignarureGenerator(100190, "57e08ead78cf63721eed92911f2dfe8a1a1152ebc880877ceae96e406c16dbab")(r, z, []byte("test"))
	// fmt.Println(x)
	// fmt.Println(y)
	// flashUpdateOrderExample()
	shopeeUpdateOrderExample()
}

func shopeeCreateOrderExample() {
	shopeeTimeSlotAPI := httpclient.NewHTTPPoster[shopee.TimeSlotRequest, shopee.TimeSlotResponse](http.DefaultClient, "https://test-stable.spx.co.th/open/api/v1/order/get_pickup_time", map[string]string{})
	shopeeCreateOrderPoster := httpclient.NewHTTPPoster[shopee.CreateOrderRequest, shopee.CreateOrderResponse](http.DefaultClient, "https://test-stable.spx.co.th/open/api/v1/order/batch_create_order", map[string]string{})
	shopeeUpdateOrderPoster := httpclient.NewHTTPPoster[shopee.UpdateOrderRequest, shopee.UpdateOrderResponse](http.DefaultClient, "https://test-stable.spx.co.th/open/api/v1/order/batch_update_order", map[string]string{})

	svc := shopee.NewShopeeService(100190, "57e08ead78cf63721eed92911f2dfe8a1a1152ebc880877ceae96e406c16dbab", 36439626319285, "b32776af-28c0-4283-971c-92ac48c01afe", shopeeCreateOrderPoster, shopeeUpdateOrderPoster, shopeeTimeSlotAPI)
	trackingNo, err := svc.CreateOrder(courierx.Order{
		WeightInGram: 1000,
		IsCOD:        false,
		Sender: courierx.OrderAddress{
			Name:          "John Wick",
			AddressDetail: "dashi",
			District:      "อำเภอเมืองบึงกาฬ",
			Province:      "จังหวัดบึงกาฬ",
			Phone:         "66898765432",
			PostalCode:    "38000",
		},
		Receiver: courierx.OrderAddress{
			Name:          "น้ำพริกแม่อำพร",
			AddressDetail: "sdfsdf",
			District:      "อำเภอเมืองบึงกาฬ",
			Province:      "จังหวัดบึงกาฬ",
			Phone:         "0812345679",
			PostalCode:    "50210",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("shopee trackingNo: %s\n", trackingNo)
}

func shopeeUpdateOrderExample() {
	shopeeTimeSlotAPI := httpclient.NewHTTPPoster[shopee.TimeSlotRequest, shopee.TimeSlotResponse](http.DefaultClient, "https://test-stable.spx.co.th/open/api/v1/order/get_pickup_time", map[string]string{})
	shopeeCreateOrderPoster := httpclient.NewHTTPPoster[shopee.CreateOrderRequest, shopee.CreateOrderResponse](http.DefaultClient, "https://test-stable.spx.co.th/open/api/v1/order/batch_create_order", map[string]string{})
	shopeeUpdateOrderPoster := httpclient.NewHTTPPoster[shopee.UpdateOrderRequest, shopee.UpdateOrderResponse](http.DefaultClient, "https://test-stable.spx.co.th/open/api/v1/order/batch_update_order", map[string]string{})

	svc := shopee.NewShopeeService(100190, "57e08ead78cf63721eed92911f2dfe8a1a1152ebc880877ceae96e406c16dbab", 36439626319285, "b32776af-28c0-4283-971c-92ac48c01afe", shopeeCreateOrderPoster, shopeeUpdateOrderPoster, shopeeTimeSlotAPI)
	trackingNo, err := svc.CreateOrder(courierx.Order{
		WeightInGram: 1000,
		IsCOD:        false,
		Sender: courierx.OrderAddress{
			Name:          "John Wick",
			AddressDetail: "dashi",
			District:      "อำเภอเมืองบึงกาฬ",
			Province:      "จังหวัดบึงกาฬ",
			Phone:         "66898765432",
			PostalCode:    "38000",
		},
		Receiver: courierx.OrderAddress{
			Name:          "น้ำพริกแม่อำพร",
			AddressDetail: "sdfsdf",
			District:      "อำเภอเมืองบึงกาฬ",
			Province:      "จังหวัดบึงกาฬ",
			Phone:         "0812345679",
			PostalCode:    "50210",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("shopee trackingNo: %s\n", trackingNo)

	svc.UpdateOrder(trackingNo, courierx.Order{
		WeightInGram: 1000,
		IsCOD:        false,
		Sender: courierx.OrderAddress{
			Name:          "NewJohn Wick",
			AddressDetail: "Newdashi",
			District:      "อำเภอเมืองบึงกาฬ",
			Province:      "จังหวัดบึงกาฬ",
			Phone:         "66898765432",
			PostalCode:    "38000",
		},
		Receiver: courierx.OrderAddress{
			Name:          "Newน้ำพริกแม่อำพร",
			AddressDetail: "New sdfsdf",
			District:      "อำเภอเมืองบึงกาฬ",
			Province:      "จังหวัดบึงกาฬ",
			Phone:         "0812345679",
			PostalCode:    "50210",
		},
	})
}

func shopeeCreateOrderExample2() {
	var appId = uint64(100190)
	var appSecret = "57e08ead78cf63721eed92911f2dfe8a1a1152ebc880877ceae96e406c16dbab"
	var timestamp = time.Now().Unix()

	var randomNum = rand.Int63()
	reqData := &HttpReqData{
		UserId:      36439626319285,
		UserSecret:  "b32776af-28c0-4283-971c-92ac48c01afe",
		ServiceType: 1,
	}
	payload, err := json.Marshal(reqData)
	if err != nil {
		return
	}
	// Generate check sign

	checkSign := makeSignarureGenerator(appId, appSecret)(randomNum, timestamp, payload)
	// checkSign, err := GenerateCheckSign(appId, appSecret, timestamp, randomNum, payload)
	// if err != nil {
	// 	return
	// }
	fmt.Println(checkSign)

	// Do http request
	httpReq, err := http.NewRequest(http.MethodPost, "https://test-stable.spx.co.th/open/api/v1/order/get_pickup_time", bytes.NewReader(payload))
	if err != nil {
		return
	}
	httpReq.Header.Set("app-id", strconv.FormatUint(appId, 10))
	httpReq.Header.Set("check-sign", checkSign)
	httpReq.Header.Set("timestamp", strconv.FormatInt(timestamp, 10))
	httpReq.Header.Set("random-num", strconv.FormatInt(randomNum, 10))
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	print(string(body))
}

func dhlCreateOrderExample() {

}

func flashUpdateOrderExample() {
	fCreate := httpclient.NewHTTPFormPoster[flash.FlashCreateOrderAPIResponse](http.DefaultClient)
	fUpdate := httpclient.NewHTTPFormPoster[flash.FlashOrderUpdateAPIResponse](http.DefaultClient)
	fs := flash.NewFlashService(fCreate, fUpdate, "8db711e11b3fe34f793444d6f2b4679be9da45446fbb82b84e40e90e1868ed75", "AA2315", "https://open-api-tra.flashexpress.com")
	trackingNo, err := fs.CreateOrder(courierx.Order{
		WeightInGram: 1000,
		IsCOD:        false,
		Sender: courierx.OrderAddress{
			Name:          "หอมรวม  create order test name",
			AddressDetail: "example detail address",
			District:      "เมืองอุบลราชธานี",
			Province:      "อุบลราชธานี",
			Phone:         "0123456789",
			PostalCode:    "34000",
		},
		Receiver: courierx.OrderAddress{
			Name:          "น้ำพริกแม่อำพร",
			AddressDetail: "example detail address",
			District:      "สันทราย",
			Province:      "เชียงใหม่",
			Phone:         "0123456789",
			PostalCode:    "50210",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("create order flash trackingNo: %s\n", trackingNo)
	fs.UpdateOrder(trackingNo, courierx.Order{
		WeightInGram: 1000,
		IsCOD:        false,
		Sender: courierx.OrderAddress{
			Name:          "new หอมรวม  create order test name",
			AddressDetail: "new example detail address",
			District:      "เมืองอุบลราชธานี",
			Province:      "อุบลราชธานี",
			Phone:         "0812345679",
			PostalCode:    "34000",
		},
		Receiver: courierx.OrderAddress{
			Name:          "น้ำพริกแม่อำพร",
			AddressDetail: "example detail address",
			District:      "สันทราย",
			Province:      "เชียงใหม่",
			Phone:         "0898765432",
			PostalCode:    "50210",
		},
	})
}

func flashCreateOrderExample() {
	fCreate := httpclient.NewHTTPFormPoster[flash.FlashCreateOrderAPIResponse](http.DefaultClient)
	fUpdate := httpclient.NewHTTPFormPoster[flash.FlashOrderUpdateAPIResponse](http.DefaultClient)
	fs := flash.NewFlashService(fCreate, fUpdate, "8db711e11b3fe34f793444d6f2b4679be9da45446fbb82b84e40e90e1868ed75", "AA2315", "https://open-api-tra.flashexpress.com")
	trackingNo, err := fs.CreateOrder(courierx.Order{
		WeightInGram: 1000,
		IsCOD:        false,
		Sender: courierx.OrderAddress{
			Name:          "หอมรวม  create order test name",
			AddressDetail: "example detail address",
			District:      "เมืองอุบลราชธานี",
			Province:      "อุบลราชธานี",
			Phone:         "0123456789",
			PostalCode:    "34000",
		},
		Receiver: courierx.OrderAddress{
			Name:          "น้ำพริกแม่อำพร",
			AddressDetail: "example detail address",
			District:      "สันทราย",
			Province:      "เชียงใหม่",
			Phone:         "0123456789",
			PostalCode:    "50210",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("flash trackingNo: %s\n", trackingNo)
}

type HttpReqData struct {
	UserId      uint64 `json:"user_id"`
	UserSecret  string `json:"user_secret"`
	ServiceType uint32 `json:"service_type"`
}

func GenerateCheckSign(appId uint64, secret string, timestamp, random int64, payload []byte) (string, error) {
	originalValue := fmt.Sprintf("%d_%d_%d_%s", appId, timestamp, random, payload)
	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(originalValue))
	if err != nil {
		return "", err
	}
	checkSign := hex.EncodeToString(h.Sum(nil))
	return checkSign, nil
}

func makeSignarureGenerator(appID uint64, secret string) func(timestamp, randomInt64 int64, payload []byte) string {
	return func(randomInt64 int64, timestamp int64, payload []byte) string {
		originalValue := fmt.Sprintf("%d_%d_%d_%s", appID, timestamp, randomInt64, payload)
		h := hmac.New(sha256.New, []byte(secret))
		_, err := h.Write([]byte(originalValue))
		if err != nil {
			fmt.Println("Error writing to hmac:", err)
			return ""
		}

		return hex.EncodeToString(h.Sum(nil))
	}
}
