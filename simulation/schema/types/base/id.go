// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math/rand"
	"strconv"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

func GenerateRandomID(r *rand.Rand) ids.ID {
	return baseIDs.NewStringID(simulation.RandStringOfLength(r, r.Intn(99)))
}

func GenerateRandomIDWithDec(r *rand.Rand) ids.ID {
	return baseIDs.NewStringID(sdkTypes.MustNewDecFromStr(strconv.FormatInt(r.Int63(), 10)).String())
}

func GenerateRandomIDWithInt64(r *rand.Rand) ids.ID {
	return baseIDs.NewStringID(strconv.FormatInt(r.Int63(), 10))
}
