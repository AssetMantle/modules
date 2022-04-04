// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/types"
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
