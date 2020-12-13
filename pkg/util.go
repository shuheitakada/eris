package pkg

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CreateRandomNumber(digit int) int {
	max := int(math.Pow10(digit) - 1)
	subtrahendMax := 9 * int(math.Pow10(digit-1))
	subtrahend := rand.Intn(subtrahendMax)
	return max - subtrahend
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"

func CreateRandomString(digit int) string {
	b := make([]byte, digit)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
