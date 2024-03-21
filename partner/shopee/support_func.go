package shopee

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

func localUnixFunc() int64 {
	return time.Now().Local().Unix()
}

func secureRandomInt64() int64 {
	for i := 0; i < 3; i++ {
		buf := make([]byte, 8)

		_, err := rand.Read(buf)
		if err != nil {
			fmt.Println("Error generating random number:", err)
			time.Sleep(time.Millisecond * 100)
			continue
		}

		num := int64(binary.BigEndian.Uint64(buf))

		return num
	}

	fmt.Println("All retries secureRandomInt64 failed, returning default value")
	return 0
}

func makeSignarureGenerator(appID uint64, secret string) func(timestamp, randomInt64 int64, payload []byte) string {
	return func(timestamp, randomInt64 int64, payload []byte) string {
		originalValue := fmt.Sprintf("%d_%d_%d_%s", appID, timestamp, randomInt64, payload)
		h := hmac.New(sha256.New, []byte(secret))
		_, err := h.Write([]byte(originalValue))
		if err != nil {
			return ""
		}
		checkSign := hex.EncodeToString(h.Sum(nil))
		return checkSign
	}
}
