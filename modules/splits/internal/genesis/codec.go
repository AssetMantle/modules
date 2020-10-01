/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package genesis

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/schema"
)

var packageCodec = codec.New()

func init() {
	mapper.Mapper.RegisterCodec(packageCodec)
	schema.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
