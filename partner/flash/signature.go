package flash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func generateSignature(keyedOrderInfo map[string]string, secret string) string {
	keys := make([]string, 0, len(keyedOrderInfo))
	for k := range keyedOrderInfo {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var builder strings.Builder
	for _, k := range keys {
		if v := keyedOrderInfo[k]; v != "" {
			if builder.Len() > 0 {
				builder.WriteString("&")
			}
			builder.WriteString(fmt.Sprintf("%s=%s", k, v))
		}
	}

	builder.WriteString("&key=" + secret)

	plainSignature := builder.String()
	h := sha256.New()
	h.Write([]byte(plainSignature))
	signature := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))

	return signature
}
