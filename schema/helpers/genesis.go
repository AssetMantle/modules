/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Genesis interface {
	Default() Genesis
	Validate() error
	Import(sdkTypes.Context, Mapper, Parameters)
	Export(sdkTypes.Context, Mapper, Parameters) Genesis

	Encode() []byte
	Decode([]byte) Genesis

	Initialize([]Mappable, []types.Parameter) Genesis

	GetParameterList() []types.Parameter
	GetMappableList() []Mappable
}
