/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package common

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/mappable"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

var Codec *codec.Codec

func init() {
	Codec = codecUtilities.RegisterModuleCodec(key.Prototype, mappable.Prototype)
}
