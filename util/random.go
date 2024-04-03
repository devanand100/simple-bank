package util

import (
	"math/rand"
)

var chars = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	result := make([]byte, n)

	for i := 0; i < n; i++ {
		result[i] = chars[RandomInt(0, int64(len(chars)-1))]
	}

	return string(result)
}

func RandomCurrency() string {
	currency := []string{"INR", "YAN", "USD"}
	return currency[RandomInt(0, int64(len(currency)-1))]
}
