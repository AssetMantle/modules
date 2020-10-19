/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func readMetaID(metaIDString string) types.ID {
	return NewMetaID(base.NewID(metaIDString))
}
func metaIDFromInterface(id types.ID) metaID {
	switch value := id.(type) {
	case metaID:
		return value
	default:
		return metaIDFromInterface(readMetaID(id.String()))
	}
}
