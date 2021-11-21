package utils

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

// RandomInt generates number between min and max 
func RandomInt(min, max int64) int64 {
    return min + rand.Int63n(max-min+1)
}
const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomString generates randome string
func RandomString(n int) string {
	var stringBuilder strings.Builder
	alphabetLength := len(alphabet)

	for i := 0; i < n; i++ {
		byteRamdom := alphabet[rand.Intn(alphabetLength)]
		stringBuilder.WriteByte(byteRamdom)
	}

	return stringBuilder.String()
}

// RandomOwner generates Owner
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates Money
func RandomMoney() int64 {
	return RandomInt(0, 100000)
}
// RandomCurrency generates currency
func RandomCurrency() string {
	currencies := []string {"USD", "EUR", "CAD", "VND"}
	lengthCurrency := len(currencies)
	currency := currencies[rand.Intn(lengthCurrency)] 
	return currency
}
