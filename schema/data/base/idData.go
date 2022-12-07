// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.IDData = (*IdDataI_IdData)(nil)

func (idData *IdDataI_IdData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(idData)
}
func (idData *IdDataI_IdData) Compare(listable traits.Listable) int {
	compareIDData, err := idDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(idData.Bytes(), compareIDData.Bytes())
}
func (idData *IdDataI_IdData) String() string {
	return idData.IdData.String()
}
func (idData *IdDataI_IdData) Bytes() []byte {
	return idData.Bytes()
}
func (idData *IdDataI_IdData) GetType() ids.StringID {
	return dataConstants.IDDataID
}
func (idData *IdDataI_IdData) ZeroValue() data.Data {
	return IDDataPrototype()
}
func (idData *IdDataI_IdData) GenerateHashID() ids.HashID {
	return baseIDs.GenerateHashID(idData.Bytes())
}
func (idData *IdDataI_IdData) Get() ids.ID {
	return idData.IdData.Value
}

func idDataFromInterface(listable traits.Listable) (*IdDataI, error) {
	switch value := listable.(type) {
	case *IdDataI:
		return value, nil
	default:
		return nil, constants.MetaDataError
	}
}

func IDDataPrototype() data.IDData {
	return NewIDData(baseIDs.NewStringID(""))
}

func NewIDData(value ids.ID) data.IDData {
	return &IdDataI{
		Impl: &IdDataI_IdData{
			IdData: &IDData{
				Value: value.(*baseIDs.IDI),
			},
		},
	}
}
