// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package random

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

// counter is a counter that adds when each time a function is called
var counter int64

// GenerateUniqueIdentifier is a random unique ID generator, output is a string.
// Warning: Non-deterministic, not to used to generate blockchain state, for testing and client utility only
func GenerateUniqueIdentifier(prefix ...string) string {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(89999))
	if err != nil {
		panic(err)
	}

	atomic.AddInt64(&counter, 1)

	return strings.Join(prefix, "") +
		strconv.Itoa(10000+int(counter)%89999) +
		strconv.Itoa(10000000+int(time.Now().UnixNano())%89999999) +
		strconv.FormatInt(10000+randomNumber.Int64(), 10)
}
