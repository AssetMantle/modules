// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	"github.com/AssetMantle/modules/constants/errors"
	idsConstants "github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/utilities/meta"
)

type idData struct {
	Value types.ID `json:"value"`
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
func (idData idData) GetType() types.ID {
	return idsConstants.IDDataID
}
func (idData idData) ZeroValue() types.Data {
	return NewIDData(baseIDs.NewID(""))
}
func (idData idData) GenerateHash() types.ID {
	return baseIDs.NewID(meta.Hash(idData.Value.String()))
}
func (idData idData) Get() types.ID {
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

func NewIDData(value types.ID) types.Data {
	return idData{
		Value: value,
	}
}

func ReadIDData(idData string) (types.Data, error) {
	return NewIDData(baseIDs.NewID(idData)), nil
}
