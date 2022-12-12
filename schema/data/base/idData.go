// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.IDData = (*Data_IdData)(nil)

func (idData *Data_IdData) Unmarshal(bytes []byte) error {
	// TODO implement me
	panic("implement me")
}
func (idData *Data_IdData) GetID() ids.ID {
	return baseIDs.GenerateDataID(idData)
}
func (idData *Data_IdData) Compare(listable traits.Listable) int {
	compareIDData, err := dataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(idData.Bytes(), compareIDData.Bytes())
}
func (idData *Data_IdData) String() string {
	return idData.IdData.String()
}
func (idData *Data_IdData) Bytes() []byte {
	return idData.Bytes()
}
func (idData *Data_IdData) GetType() ids.ID {
	return dataConstants.IDDataID
}
func (idData *Data_IdData) ZeroValue() data.Data {
	return IDDataPrototype()
}
func (idData *Data_IdData) GenerateHashID() ids.ID {
	return baseIDs.GenerateHashID(idData.Bytes())
}
func (idData *Data_IdData) Get() ids.ID {
	return idData.IdData.Value
}

func IDDataPrototype() data.Data {
	return NewIDData(baseIDs.NewStringID(""))
}

func NewIDData(value ids.ID) data.Data {
	return &Data{
		Impl: &Data_IdData{
			IdData: &IDData{
				Value: value.(*baseIDs.ID),
			},
		},
	}
}
