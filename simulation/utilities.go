// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulation

import (
	"github.com/AssetMantle/modules/helpers"
	baseSimulation "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/x/classifications/parameters/bondRate"
	"github.com/AssetMantle/schema/go/data"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/qualified"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"math/rand"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

var (
	Immutables qualified.Immutables = &baseQualified.Immutables{}
	Mutables   qualified.Mutables   = &baseQualified.Mutables{}
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

func generateGenesisProperties(r *rand.Rand) {
	Immutables = baseQualified.NewImmutables(baseSimulation.GenerateRandomMetaPropertyListWithoutData(r))
	Mutables = baseQualified.NewMutables(baseSimulation.GenerateRandomPropertyList(r))
}

func GetGenesisProperties(r *rand.Rand) (qualified.Immutables, qualified.Mutables) {
	if Immutables.(*baseQualified.Immutables).PropertyList == nil {
		generateGenesisProperties(r)
	}
	return Immutables, Mutables
}

func CalculateBondAmount(immutables qualified.Immutables, mutables qualified.Mutables) data.NumberData {
	totalWeight := sdkTypes.ZeroInt()
	for _, property := range append(immutables.GetImmutablePropertyList().GetList(), mutables.GetMutablePropertyList().GetList()...) {
		totalWeight = totalWeight.Add(property.Get().GetBondWeight())
	}
	return baseData.NewNumberData(bondRate.Parameter.MetaProperty.Data.Get().(data.NumberData).Get().Mul(totalWeight))
}

func ExecuteMessage(context sdkTypes.Context, module helpers.Module, message helpers.Message) (*sdkTypes.Result, error) {
	return module.GetTransactions().Get(message.Type()).HandleMessage(sdkTypes.WrapSDKContext(context), message)
}
