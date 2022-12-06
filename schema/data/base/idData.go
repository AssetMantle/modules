// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data/base"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

type idData base.IDData

var _ data.IDData = (*idData)(nil)

func (idData *idData) GetID() ids.DataID {
	return baseIDs.NewDataID(idData)
}
func (idData *idData) Compare(listable traits.Listable) int {
	compareIDData, err := idDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(idData.Bytes(), compareIDData.Bytes())
}
func (idData *idData) String() string {
	return idData.Value.String()
}
func (idData *idData) Bytes() []byte {
	return idData.Bytes()
}
func (idData *idData) GetType() ids.StringID {
	return dataConstants.IDDataID
}
func (idData *idData) ZeroValue() data.Data {
	return NewIDData(baseIDs.NewStringID(""))
}
func (idData *idData) GenerateHashID() ids.HashID {
	return baseIDs.GenerateHashID(idData.Bytes())
}
func (idData *idData) Get() ids.ID {
	return idData.(data.IDData)
}

func idDataFromInterface(listable traits.Listable) (*idData, error) {
	switch value := listable.(type) {
	case *idData:
		return value, nil
	default:
		return &idData{}, constants.MetaDataError
	}
}

func IDDataPrototype() data.IDData {
	return (&idDataI{}).ZeroValue().(data.IDData)
}

func NewIDData(value ids.ID) data.IDData {
	return &idData{
		Value: value,
	}
}
