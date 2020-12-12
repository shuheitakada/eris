package pkg

import (
	"math"
	"math/rand"
	"time"
)

func CreateRandomNumber(digit int) int {
	rand.Seed(time.Now().UnixNano())
	max := int(math.Pow10(digit) - 1)
	subtrahendMax := 9 * int(math.Pow10(digit-1))
	subtrahend := rand.Int63n(int64(subtrahendMax))
	return max - int(subtrahend)
}
