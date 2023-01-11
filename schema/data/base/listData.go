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
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.ListData = (*ListData)(nil)

func (listData *ListData) Get() []data.AnyData {
	anyDataList := make([]data.AnyData, len(listData.DataList))
	for i, anyData := range listData.DataList {
		anyDataList[i] = anyData
	}
	return anyDataList
}
func (listData *ListData) Search(data data.Data) (int, bool) {
	dataList := NewDataList()
	for _, val := range listData.DataList {
		dataList = dataList.Add(val.Get())
	}
	return dataList.Search(data)
}
func (listData *ListData) Add(data ...data.Data) data.ListData {
	for _, i := range data {
		listData.DataList = append(listData.DataList, i.ToAnyData().(*AnyData))
	}
	return listData
}
func (listData *ListData) Remove(data ...data.Data) data.ListData {
	dataList := NewDataList()
	for _, val := range listData.DataList {
		dataList = dataList.Add(val.Get())
	}
	dataList = dataList.Remove(data...)
	listData.DataList = make([]*AnyData, 0)
	for _, datum := range dataList.GetList() {
		listData.DataList = append(listData.DataList, datum.ToAnyData().(*AnyData))
	}
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
	bytesList := make([][]byte, len(listData.DataList))

	for i, datum := range listData.DataList {
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
	return NewListData([]data.Data{}...)
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
			ListData: listData,
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
func NewListData(data ...data.Data) data.ListData {
	dataList := make([]*AnyData, 0)
	for _, datum := range data {
		dataList = append(dataList, datum.ToAnyData().(*AnyData))
	}
	return &ListData{DataList: dataList}
}
