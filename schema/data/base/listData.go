// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"sort"
	"strings"

	Data "github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

var _ Data.ListData = (*ListData)(nil)

func (listData *ListData) ValidateBasic() error {
	for _, data := range listData.DataList {
		if data.GetTypeID().Compare(listData.GetTypeID()) == 0 {
			return errorConstants.IncorrectFormat.Wrapf("ListData cannot contain ListData")
		}

		if err := data.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}
func (listData *ListData) Get() []Data.AnyData {
	listData = listData.Sort()
	anyDataList := make([]Data.AnyData, len(listData.DataList))
	for i, anyData := range listData.DataList {
		anyDataList[i] = anyData
	}

	return anyDataList
}
func (listData *ListData) GetBondWeight() int64 {
	return dataConstants.ListDataWeight
}
func (listData *ListData) AsString() string {
	listData = listData.Sort()
	dataStrings := make([]string, len(listData.DataList))

	for i, datum := range listData.DataList {
		dataStrings[i] = datum.AsString()
	}

	return stringUtilities.JoinListStrings(dataStrings...)
}
func (listData *ListData) FromString(dataString string) (Data.Data, error) {
	dataString = strings.TrimSpace(dataString)
	if dataString == "" {
		return PrototypeListData(), nil
	}

	dataStringList := stringUtilities.SplitListString(dataString)
	dataList := make([]Data.Data, len(dataStringList))

	for i, datumString := range dataStringList {
		// TODO: check if all data are same type,[T]
		data, err := PrototypeAnyData().FromString(datumString)
		if err != nil {
			return PrototypeListData(), err
		}

		dataList[i] = data
	}

	return NewListData(dataList...), nil
}
func (listData *ListData) Search(data Data.Data) (int, bool) {
	listData = listData.Sort()
	index := sort.Search(
		len(listData.DataList),
		func(i int) bool {
			return listData.DataList[i].Compare(data) >= 0
		},
	)

	if index < len(listData.DataList) && listData.DataList[index].Compare(data) == 0 {
		return index, true
	}

	return index, false
}
func (listData *ListData) Add(data ...Data.Data) Data.ListData {
	updatedListData := listData.Sort()
	for _, listable := range data {
		if index, found := updatedListData.Search(listable); !found {
			updatedListData.DataList = append(updatedListData.DataList, listable.ToAnyData().(*AnyData))
			copy(updatedListData.DataList[index+1:], updatedListData.DataList[index:])
			updatedListData.DataList[index] = listable.ToAnyData().(*AnyData)
		}
	}

	return updatedListData
}
func (listData *ListData) Remove(data ...Data.Data) Data.ListData {
	updatedListData := listData.Sort()

	for _, listable := range data {
		if index, found := updatedListData.Search(listable); found {
			updatedListData.DataList = append(updatedListData.DataList[:index], updatedListData.DataList[index+1:]...)
		}
	}

	return updatedListData
}
func (listData *ListData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(listData.Sort())
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
	bytesList := make([][]byte, len(listData.Sort().DataList))

	for i, datum := range listData.DataList {
		if datum != nil {
			bytesList[i] = datum.Bytes()
		}
	}
	// TODO see if separator required
	return bytes.Join(bytesList, nil)
}
func (listData *ListData) GetTypeID() ids.StringID {
	return dataConstants.ListDataTypeID
}
func (listData *ListData) ZeroValue() Data.Data {
	return NewListData([]Data.Data{}...)
}
func (listData *ListData) GenerateHashID() ids.HashID {
	if listData.Compare(listData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(listData.Sort().Bytes())
}
func (listData *ListData) Sort() *ListData {
	sort.Slice(listData.DataList, func(i, j int) bool {
		return listData.DataList[i].Compare(listData.DataList[j]) < 0
	})

	return listData
}
func (listData *ListData) ToAnyData() Data.AnyData {
	return &AnyData{
		Impl: &AnyData_ListData{
			ListData: listData,
		},
	}
}
func listDataFromInterface(listable traits.Listable) (*ListData, error) {
	switch value := listable.(type) {
	case *ListData:
		return value.Sort(), nil
	default:
		return &ListData{}, errorConstants.IncorrectFormat.Wrapf("unsupported type")
	}
}

func PrototypeListData() Data.ListData {
	return (&ListData{}).ZeroValue().(Data.ListData)
}

// NewListData
// * onus of ensuring all Data are of the same type is on DataList
func NewListData(data ...Data.Data) Data.ListData {
	return (&ListData{}).Add(data...)
}
