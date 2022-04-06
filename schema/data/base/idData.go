// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
	"github.com/AssetMantle/modules/utilities/meta"
)

type idData struct {
	Value types.ID `json:"value"`
}

var _ data.IDData = (*idData)(nil)

func (idData idData) GetID() types.ID {
	return base.NewDataID(idData)
}
func (idData idData) Compare(data types.Data) int {
	compareIDData, err := idDataFromInterface(data)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(idData.Value.Bytes(), compareIDData.Value.Bytes())
}
func (idData idData) String() string {
	return idData.Value.String()
}
func (idData idData) ZeroValue() types.Data {
	return NewIDData(base.NewID(""))
}
func (idData idData) GetTypeID() types.ID {
	return IDDataID
}
func (idData idData) GenerateHashID() types.ID {
	return base.NewID(meta.Hash(idData.Value.String()))
}
func (idData idData) Get() types.ID {
	return idData.Value
}

func idDataFromInterface(data types.Data) (idData, error) {
	switch value := data.(type) {
	case idData:
		return value, nil
	default:
		return idData{}, errors.MetaDataError
	}
}

func NewIDData(value types.ID) types.Data {
	return idData{
		Value: value,
	}
}

func ReadIDData(idData string) (types.Data, error) {
	return NewIDData(base.NewID(idData)), nil
}
