package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"strings"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

type getter interface {
	get() data.Data
}

func (x *AnyData_AccAddressData) get() data.Data {
	return x.AccAddressData
}
func (x *AnyData_BooleanData) get() data.Data {
	return x.BooleanData
}
func (x *AnyData_DecData) get() data.Data {
	return x.DecData
}
func (x *AnyData_HeightData) get() data.Data {
	return x.HeightData
}
func (x *AnyData_IDData) get() data.Data {
	return x.IDData
}
func (x *AnyData_StringData) get() data.Data {
	return x.StringData
}
func (x *AnyData_ListData) get() data.Data {
	return x.ListData
}

var _ data.AnyData = (*AnyData)(nil)

func (x *AnyData) IsAnyData() {}
func (x *AnyData) AsString() string {
	return x.Impl.(getter).get().AsString()
}
func (x *AnyData) FromString(dataString string) (data.Data, error) {
	dataTypeString, _ := splitDataTypeAndValueStrings(dataString)
	if dataTypeString != "" {
		var Data data.Data

		var err error

		switch baseIDs.NewStringID(dataTypeString).AsString() {
		case dataConstants.AccAddressDataID.AsString():
			Data, err = PrototypeAccAddressData().FromString(dataString)
		case dataConstants.BooleanDataID.AsString():
			Data, err = PrototypeBooleanData().FromString(dataString)
		case dataConstants.DecDataID.AsString():
			Data, err = PrototypeDecData().FromString(dataString)
		case dataConstants.HeightDataID.AsString():
			Data, err = PrototypeHeightData().FromString(dataString)
		case dataConstants.IDDataID.AsString():
			Data, err = PrototypeIDData().FromString(dataString)
		case dataConstants.ListDataID.AsString():
			Data, err = PrototypeListData().FromString(dataString)
		case dataConstants.StringDataID.AsString():
			Data, err = PrototypeStringData().FromString(dataString)
		default:
			Data, err = nil, errorConstants.UnsupportedParameter
		}

		if err != nil {
			return nil, err
		}

		return Data, nil
	}

	return nil, errorConstants.IncorrectFormat
}
func (x *AnyData) Get() data.Data {
	return x.Impl.(getter).get()
}
func (x *AnyData) GetID() ids.DataID {
	return x.Impl.(getter).get().GetID()
}
func (x *AnyData) Bytes() []byte {
	return x.Impl.(getter).get().Bytes()
}
func (x *AnyData) GetType() ids.StringID {
	return x.Impl.(getter).get().GetType()
}
func (x *AnyData) ZeroValue() data.Data {
	return x.Impl.(getter).get().ZeroValue()
}
func (x *AnyData) GenerateHashID() ids.HashID {
	return x.Impl.(getter).get().GenerateHashID()
}
func (x *AnyData) ToAnyData() data.AnyData {
	return x.Impl.(getter).get().ToAnyData()
}
func (x *AnyData) Compare(listable traits.Listable) int {
	return x.Impl.(getter).get().Compare(listable)
}
func (x *AnyData) GetBondWeight() sdkTypes.Dec {
	return x.Impl.(getter).get().GetBondWeight()
}

func dataFromListable(listable traits.Listable) (data.Data, error) {
	switch value := listable.(type) {
	case data.Data:
		return value, nil
	default:
		return nil, errorConstants.MetaDataError
	}
}

func PrototypeAnyData() data.AnyData {
	return &AnyData{}
}
func joinDataTypeAndValueStrings(dataType, dataValue string) string {
	return strings.Join([]string{dataType, dataValue}, dataConstants.DataTypeAndValueSeparator)
}
func splitDataTypeAndValueStrings(dataTypeAndValueString string) (dataType, dataValue string) {
	if dataTypeAndValue := strings.SplitN(dataTypeAndValueString, dataConstants.DataTypeAndValueSeparator, 2); len(dataTypeAndValue) == 1 {
		return dataTypeAndValue[0], ""
	} else if len(dataTypeAndValue) == 2 {
		return dataTypeAndValue[0], dataTypeAndValue[1]
	} else {
		return "", ""
	}
}
