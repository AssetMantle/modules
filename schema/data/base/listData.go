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
)

var _ data.ListData = (*ListData)(nil)

func (listData *ListData) Get() []data.AnyData {
	anyDataList := make([]data.AnyData, listData.Size())
	for i, anyData := range listData.Value.DataList {
		anyDataList[i] = anyData
	}
	return anyDataList
}
func (listData *ListData) Search(data data.AnyData) (int, bool) {
	return listData.Value.Search(data)
}
func (listData *ListData) Add(data ...data.AnyData) data.ListData {
	listData.Value = listData.Value.Add(data...).(*base.AnyDataList)
	return listData
}
func (listData *ListData) Remove(data ...data.AnyData) data.ListData {
	listData.Value = listData.Value.Remove(data...).(*base.AnyDataList)
	return listData
}
func (listData *ListData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(listData)
}
func (listData *ListData) Compare(listable traits.Listable) int {
	compareListData, err := listDataFromInterface(listable)
	if err != nil {
		panic(err)
	}

	// TODO check for optimization
	return bytes.Compare(listData.Bytes(), compareListData.Bytes())
}
func (listData *ListData) Bytes() []byte {
	bytesList := make([][]byte, listData.Value.Size())

	for i, datum := range listData.Value.GetList() {
		if datum != nil {
			bytesList[i] = datum.Bytes()
		}
	}
	// TODO see if separator required
	return bytes.Join(bytesList, nil)
}
func (listData *ListData) GetType() ids.StringID {
	return dataConstants.ListDataID
}
func (listData *ListData) ZeroValue() data.Data {
	return NewListData(base.NewDataList([]data.Data{}...))
}
func (listData *ListData) GenerateHashID() ids.HashID {
	if listData.Compare(listData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(listData.Bytes())
}
func (listData *ListData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_ListData{
			ListData: *listData,
		},
	}
}
func listDataFromInterface(listable traits.Listable) (*ListData, error) {
	switch value := listable.(type) {
	case *ListData:
		return value, nil
	default:
		return &ListData{}, errorConstants.MetaDataError
	}
}

func ListDataPrototype() data.ListData {
	return (&ListData{}).ZeroValue().(data.ListData)
}

// NewListData
// * onus of ensuring all Data are of the same type is on DataList
func NewListData(value lists.AnyDataList) data.ListData {
	return &ListData{Value: value.(*base.AnyDataList)}
}
