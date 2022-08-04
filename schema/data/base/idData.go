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

type idData struct {
	Value ids.ID `json:"value"`
}

var _ data.IDData = (*idData)(nil)

func (idData idData) GetID() ids.DataID {
	return baseIDs.NewDataID(idData)
}

func (idData idData) Compare(listable traits.Listable) int {
	compareIDData, err := idDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(idData.Value.Bytes(), compareIDData.Value.Bytes())
}
func (idData idData) String() string {
	return idData.Value.String()
}
func (idData idData) Bytes() []byte {
	return idData.Value.Bytes()
}
func (idData idData) GetType() ids.StringID {
	return dataConstants.IDDataID
}
func (idData idData) ZeroValue() data.Data {
	return NewIDData(baseIDs.NewStringID(""))
}
func (idData idData) GenerateHashID() ids.HashID {
	return baseIDs.GenerateHashID(idData.Bytes())
}
func (idData idData) Get() ids.ID {
	return idData.Value
}

func idDataFromInterface(listable traits.Listable) (idData, error) {
	switch value := listable.(type) {
	case idData:
		return value, nil
	default:
		return idData{}, constants.MetaDataError
	}
}

func NewIDData(value ids.ID) data.Data {
	return idData{
		Value: value,
	}
}
