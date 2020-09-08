/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type metaID struct {
	HashID types.ID `json:"hashID" valid:"required~required field hashID missing"`
}

var _ types.ID = (*metaID)(nil)

func (metaID metaID) Bytes() []byte {
	return metaID.HashID.Bytes()
}

func (metaID metaID) String() string {
	return metaID.HashID.String()
}

func (metaID metaID) Equal(id types.ID) bool {
	return bytes.Compare(metaID.Bytes(), id.Bytes()) == 0
}

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
func generateKey(metaID types.ID) []byte {
	return append(StoreKeyPrefix, metaIDFromInterface(metaID).Bytes()...)
}
func NewMetaID(hashID types.ID) types.ID {
	return metaID{
		HashID: hashID,
	}
}
