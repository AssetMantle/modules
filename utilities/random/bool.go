// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

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
