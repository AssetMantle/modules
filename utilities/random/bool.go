package random

import (
	"math"
	"math/rand"
	"time"
)

func GenerateRandomBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(math.MaxInt)%2 == 0
}
