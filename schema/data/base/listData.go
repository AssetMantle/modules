// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/traits"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

type listData struct {
	Value lists.DataList `json:"value"`
}

var _ data.ListData = (*listData)(nil)

func (listData listData) Search(data data.Data) (int, bool) {
	return listData.Search(data)
}
func (listData listData) Add(data ...data.Data) data.ListData {
	listData.Value = listData.Value.Add(data...)
	return listData
}
func (listData listData) Remove(data ...data.Data) data.ListData {
	listData.Value = listData.Value.Remove(data...)
	return listData
}
func (listData listData) GetID() ids.DataID {
	return baseIDs.NewDataID(listData)
}
func (listData listData) Compare(listable traits.Listable) int {
	compareListData, err := listDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	// TODO check for optimization
	return bytes.Compare(listData.Bytes(), compareListData.Bytes())
}
func (listData listData) String() string {
	dataStrings := make([]string, listData.Value.Size())

	for i, datum := range listData.Value.GetList() {
		dataStrings[i] = datum.String()
	}

	return stringUtilities.JoinListStrings(dataStrings...)
}
func (listData listData) Bytes() []byte {
	hashBytesList := make([][]byte, listData.Value.Size())

	for i, datum := range listData.Value.GetList() {
		hashBytesList[i] = datum.Bytes()
	}
	// TODO see if separator required
	return bytes.Join(hashBytesList, nil)
}
func (listData listData) GetType() ids.StringID {
	return dataConstants.ListDataID
}
func (listData listData) ZeroValue() data.Data {
	return NewListData(base.NewDataList([]data.Data{}...))
}

// TODO test
func (listData listData) GenerateHashID() ids.HashID {
	if listData.Compare(listData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(listData.Bytes())
}
func listDataFromInterface(listable traits.Listable) (listData, error) {
	switch value := listable.(type) {
	case listData:
		return value, nil
	default:
		return listData{}, errorConstants.MetaDataError
	}
}

// NewListData
// * onus of ensuring all Data are of the same type is on DataList
func NewListData(value lists.DataList) data.Data {
	return listData{Value: value}
}
