// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math/rand"
	"strconv"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/AssetMantle/schema/x/ids"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
)

func GenerateRandomID(r *rand.Rand) ids.ID {
	return baseIDs.NewStringID(simulationTypes.RandStringOfLength(r, r.Intn(99)))
}

func GenerateRandomIDWithDec(r *rand.Rand) ids.ID {
	return baseIDs.NewStringID(sdkTypes.MustNewDecFromStr(strconv.FormatInt(r.Int63(), 10)).String())
}

func GenerateRandomIDWithInt64(r *rand.Rand) ids.ID {
	return baseIDs.NewStringID(strconv.FormatInt(r.Int63(), 10))
}
