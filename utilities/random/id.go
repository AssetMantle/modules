package random

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"sync/atomic"
	"time"
)

// counter is a counter that adds when each time a function is called
var counter int64

// GenerateID is a random unique ID generator, output is a string.
// Warning: Non-deterministic, do not use to generate blockchain state
func GenerateID(prefix string) string {
	randomNumber, Error := rand.Int(rand.Reader, big.NewInt(89999))
	if Error != nil {
		panic(Error)
	}

	atomic.AddInt64(&counter, 1)

	return prefix +
		strconv.Itoa(10000+int(counter)%89999) +
		strconv.Itoa(10000000+int(time.Now().UnixNano())%89999999) +
		strconv.FormatInt(10000+randomNumber.Int64(), 10)
}
