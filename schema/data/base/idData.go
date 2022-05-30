// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data"
	idsConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
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
func (idData idData) GetType() ids.ID {
	return idsConstants.IDDataID
}
func (idData idData) ZeroValue() data.Data {
	return NewIDData(baseIDs.NewID(""))
}
func (idData idData) GenerateHash() ids.ID {
	return baseIDs.NewID(stringUtilities.Hash(idData.Value.String()))
}
func (idData idData) Get() ids.ID {
	return idData.Value
}

func idDataFromInterface(listable traits.Listable) (idData, error) {
	switch value := listable.(type) {
	case idData:
		return value, nil
	default:
		return idData{}, errors.MetaDataError
	}
}

func NewIDData(value ids.ID) data.Data {
	return idData{
		Value: value,
	}
}

func ReadIDData(idData string) (data.Data, error) {
	return NewIDData(baseIDs.NewID(idData)), nil
}
