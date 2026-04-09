// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulation

import (
	"cosmossdk.io/math"
	"fmt"
	goGoProto "github.com/cosmos/gogoproto/proto"
	"math/rand"
	"strings"

	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/qualified"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/AssetMantle/modules/helpers"
	baseSimulation "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/x/classifications/parameters/bond_rate"
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
	totalWeight := math.ZeroInt()
	for _, property := range append(immutables.GetImmutablePropertyList().Get(), mutables.GetMutablePropertyList().Get()...) {
		totalWeight = totalWeight.Add(property.Get().GetBondWeight())
	}
	return baseData.NewNumberData(bond_rate.Parameter.GetMetaProperty().GetData().Get().(data.NumberData).Get().Mul(totalWeight))
}

func ExecuteMessage(context sdkTypes.Context, module helpers.Module, message helpers.Message) (*sdkTypes.Result, error) {
	// Derive service path from proto message name.
	// Proto name: "AssetMantle.modules.x.assets.transactions.burn.Message"
	// Service path: "/assets/burn"
	msgName := goGoProto.MessageName(message)
	parts := strings.Split(msgName, ".")
	if len(parts) >= 6 {
		servicePath := "/" + parts[3] + "/" + parts[5]
		if tx := module.GetTransactions().GetTransaction(servicePath); tx != nil {
			return tx.HandleMessage(sdkTypes.WrapSDKContext(context), message)
		}
	}
	return nil, fmt.Errorf("no matching transaction for message %s", msgName)
}
