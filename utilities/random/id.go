// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
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

// GenerateUniqueIdentifierList generates a list of unique IDs with the given prefix and count.
// The returned value is a slice of strings.
// The prefix is used to customize the generated IDs.
// The count specifies the number of IDs to generate.
// Each generated ID is unique.
//
// Example:
// ids := GenerateUniqueIdentifierList("user-", 5)
//
// The above example will generate 5 unique IDs with the prefix "user-".
// The returned slice will contain ["user-123...1", "user-234...2", "user-345...3", "user-456...4", "user-567...5"].
//
// Warning: This function is non-deterministic and is not suitable for generating blockchain state.
// It is intended for testing and client utility purposes only.
//
// Dependencies:
// The function depends on the GenerateUniqueIdentifier function.
//
// Code Snippet:
//
//	func GenerateUniqueIdentifierList(prefix string, count int) []string {
//	    ids := make([]string, count)
//	    for i := 0; i < count; i++ {
//	        ids[i] = GenerateUniqueIdentifier(prefix)
//	    }
//	    return ids
//	}
func GenerateUniqueIdentifierList(prefix string, count int) []string {
	ids := make([]string, count)
	for i := 0; i < count; i++ {
		ids[i] = GenerateUniqueIdentifier(prefix)
	}
	return ids
}
