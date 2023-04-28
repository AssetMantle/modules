// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulation

import (
	"context"
	"github.com/AssetMantle/modules/x/classifications/parameters/bondRate"
	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/qualified"
	"math/rand"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func RandomBool(r *rand.Rand) bool {
	return r.Intn(2) == 0
}

func GenerateRandomAddresses(r *rand.Rand) []sdkTypes.AccAddress {
	randomAccounts := simulationTypes.RandomAccounts(r, r.Intn(99))
	addresses := make([]sdkTypes.AccAddress, len(randomAccounts))

	for i, account := range randomAccounts {
		addresses[i] = account.Address
	}

	return addresses
}

func CalculateBondAmount(context context.Context, immutables qualified.Immutables, mutables qualified.Mutables) data.NumberData {
	totalWeight := sdkTypes.ZeroInt()
	for _, property := range append(immutables.GetImmutablePropertyList().GetList(), mutables.GetMutablePropertyList().GetList()...) {
		totalWeight = totalWeight.Add(property.Get().GetBondWeight())
	}
	return baseData.NewNumberData(bondRate.Parameter.MetaProperty.Data.Get().(data.NumberData).Get().Mul(totalWeight))
}
