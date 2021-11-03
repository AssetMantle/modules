/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Keeper interface {
	GetMappable(ctx sdkTypes.Context, id types.ID) Mappable
	SetMappable(ctx sdkTypes.Context, mappable Mappable) error

	GetParameters(ctx sdkTypes.Context) Parameters
	GetParameter(ctx sdkTypes.Context, id types.ID) types.Parameter
	SetParameter(ctx sdkTypes.Context, id types.ID, data types.Data) error
}
