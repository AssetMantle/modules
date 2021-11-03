/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Meta interface {
	GetData() types.Data
	helpers.Mappable
	codecTypes.UnpackInterfacesMessage
}
