/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulation

import (
	"math/rand"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

func RandomBool(r *rand.Rand) bool {
	return r.Intn(2) == 0
}

func GenerateRandomAddresses(r *rand.Rand) []sdkTypes.AccAddress {
	randomAccounts := simulation.RandomAccounts(r, r.Intn(99))
	addresses := make([]sdkTypes.AccAddress, len(randomAccounts))

	for i, account := range randomAccounts {
		addresses[i] = account.Address
	}

	return addresses
}
