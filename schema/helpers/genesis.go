/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Genesis interface {
	Default() Genesis
	Validate() error
	Import(sdkTypes.Context, Mapper, Parameters)
	Export(sdkTypes.Context, Mapper, Parameters) Genesis

	Marshall() []byte
	Unmarshall([]byte) Genesis

	Initialize([]traits.Mappable, []types.Parameter) Genesis
}
