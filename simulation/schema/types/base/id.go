// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"cosmossdk.io/math"
	"math/rand"
	"strconv"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
)

// validStringIDChars matches the schema regex: [A-Za-z0-9_]{0,30}
const validStringIDChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"

// RandValidStringID generates a random string that passes StringID.ValidateBasic().
// Length: 1-30, charset: [A-Za-z0-9_].
func RandValidStringID(r *rand.Rand) string {
	length := r.Intn(30) + 1
	b := make([]byte, length)
	for i := range b {
		b[i] = validStringIDChars[r.Intn(len(validStringIDChars))]
	}
	return string(b)
}

func GenerateRandomID(r *rand.Rand) ids.StringID {
	return baseIDs.NewStringID(RandValidStringID(r))
}

func GenerateRandomIDWithDec(r *rand.Rand) ids.ID {
	return baseIDs.NewStringID(math.LegacyMustNewDecFromStr(strconv.FormatInt(r.Int63(), 10)).String())
}

func GenerateRandomIDWithInt64(r *rand.Rand) ids.ID {
	return baseIDs.NewStringID(strconv.FormatInt(r.Int63(), 10))
}
