package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

// RandomString generated a random string with n length
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates a random owner with length of 6
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates random amount of money between 0 and 1000
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency returns an exact currency between 3
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
