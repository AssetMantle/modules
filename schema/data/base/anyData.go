package base

import (
	"strings"

	"github.com/AssetMantle/modules/schema/data"
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
func (x *AnyData_NumberData) get() data.Data {
	return x.NumberData
}
func (x *AnyData_StringData) get() data.Data {
	return x.StringData
}
func (x *AnyData_ListData) get() data.Data {
	return x.ListData
}

var _ data.AnyData = (*AnyData)(nil)

func (x *AnyData) ValidateBasic() error {
	return x.Impl.(getter).get().ValidateBasic()
}
func (x *AnyData) IsAnyData() {}
func (x *AnyData) AsString() string {
	return joinDataTypeAndValueStrings(x.Impl.(getter).get().GetTypeID().AsString(), x.Impl.(getter).get().AsString())
}
func (x *AnyData) FromString(dataString string) (data.Data, error) {
	dataTypeString, dataValueString := splitDataTypeAndValueStrings(dataString)
	if dataTypeString != "" {
		var Data data.Data

		var err error

		switch baseIDs.NewStringID(dataTypeString).AsString() {
		case PrototypeAccAddressData().GetTypeID().AsString():
			Data, err = PrototypeAccAddressData().FromString(dataValueString)
		case PrototypeBooleanData().GetTypeID().AsString():
			Data, err = PrototypeBooleanData().FromString(dataValueString)
		case PrototypeDecData().GetTypeID().AsString():
			Data, err = PrototypeDecData().FromString(dataValueString)
		case PrototypeHeightData().GetTypeID().AsString():
			Data, err = PrototypeHeightData().FromString(dataValueString)
		case PrototypeIDData().GetTypeID().AsString():
			Data, err = PrototypeIDData().FromString(dataValueString)
		case PrototypeListData().GetTypeID().AsString():
			Data, err = PrototypeListData().FromString(dataValueString)
		case PrototypeNumberData().GetTypeID().AsString():
			Data, err = PrototypeNumberData().FromString(dataValueString)
		case PrototypeStringData().GetTypeID().AsString():
			Data, err = PrototypeStringData().FromString(dataValueString)
		default:
			Data, err = nil, errorConstants.IncorrectFormat.Wrapf("type identifier is not recognized")
		}

		if err != nil {
			return nil, err
		}

		return Data, nil
	}

	return nil, errorConstants.IncorrectFormat.Wrapf("type identifier is missing")
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
func (x *AnyData) GetTypeID() ids.StringID {
	return x.Impl.(getter).get().GetTypeID()
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
func (x *AnyData) GetBondWeight() int64 {
	return x.Impl.(getter).get().GetBondWeight()
}

func dataFromListable(listable traits.Listable) (data.Data, error) {
	switch value := listable.(type) {
	case data.Data:
		return value, nil
	default:
		return nil, errorConstants.MetaDataError.Wrapf("unsupported type")
	}
}

func PrototypeAnyData() data.AnyData {
	return &AnyData{}
}
func joinDataTypeAndValueStrings(dataType, dataValue string) string {
	return strings.TrimSpace(dataType) + dataTypeAndValueSeparator + strings.TrimSpace(dataValue)
}
func splitDataTypeAndValueStrings(dataTypeAndValueString string) (dataType, dataValue string) {
	if dataTypeAndValue := strings.SplitN(dataTypeAndValueString, dataTypeAndValueSeparator, 2); len(dataTypeAndValue) == 1 {
		return strings.TrimSpace(dataTypeAndValue[0]), ""
	} else if len(dataTypeAndValue) == 2 {
		return strings.TrimSpace(dataTypeAndValue[0]), strings.TrimSpace(dataTypeAndValue[1])
	} else {
		return "", ""
	}
}

const dataTypeAndValueSeparator = "|"
