// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"strings"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
)

type dataID struct {
	Type types.ID
	Hash types.ID
}

var _ ids.DataID = (*dataID)(nil)

func (dataID dataID) String() string {
	var values []string
	values = append(values, dataID.Type.String())
	values = append(values, dataID.Hash.String())

	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}
func (dataID dataID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, dataID.Type.Bytes()...)
	Bytes = append(Bytes, dataID.Hash.Bytes()...)

	return Bytes
}
func (dataID dataID) Compare(listable traits.Listable) int {
	return bytes.Compare(dataID.Bytes(), dataIDFromInterface(listable).Bytes())
}
func (dataID dataID) GetHash() types.ID {
	return dataID.Hash
}
func dataIDFromInterface(i interface{}) dataID {
	// TODO impl
	panic("impl")
}
func NewDataID(data types.Data) ids.DataID {
	return dataID{
		Type: data.GetType(),
		Hash: data.GenerateHash(),
	}
}
