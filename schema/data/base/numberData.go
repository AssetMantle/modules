package base

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"strings"

	"github.com/AssetMantle/modules/schema/data"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.NumberData = (*NumberData)(nil)

func (numberData *NumberData) ValidateBasic() error {
	return nil
}
func (numberData *NumberData) GetID() ids.DataID {
	return baseIDs.GenerateDataID(numberData)
}
func (numberData *NumberData) GetBondWeight() int64 {
	return dataConstants.NumberDataWeight
}
func (numberData *NumberData) AsString() string {
	return strconv.FormatInt(numberData.Value, 10)
}
func (numberData *NumberData) FromString(dataString string) (data.Data, error) {
	dataString = strings.TrimSpace(dataString)
	if dataString == "" {
		return PrototypeNumberData(), nil
	}

	value, err := strconv.ParseInt(dataString, 10, 64)
	if err != nil {
		return PrototypeNumberData(), errorConstants.IncorrectFormat.Wrapf(err.Error())
	}

	return NewNumberData(value), nil
}
func (numberData *NumberData) Bytes() []byte {
	Bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(Bytes, uint64(numberData.Get()))

	return Bytes
}
func (numberData *NumberData) GetTypeID() ids.StringID {
	return dataConstants.NumberDataTypeID
}
func (numberData *NumberData) ZeroValue() data.Data {
	return NewNumberData(0)
}
func (numberData *NumberData) GenerateHashID() ids.HashID {
	if numberData.Compare(numberData.ZeroValue()) == 0 {
		return baseIDs.GenerateHashID()
	}

	return baseIDs.GenerateHashID(numberData.Bytes())
}
func (numberData *NumberData) ToAnyData() data.AnyData {
	return &AnyData{
		Impl: &AnyData_NumberData{
			NumberData: numberData,
		},
	}
}
func (numberData *NumberData) Compare(listable traits.Listable) int {
	compareNumberData, err := dataFromListable(listable)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(numberData.Bytes(), compareNumberData.Bytes())
}
func (numberData *NumberData) Get() int64 {
	return numberData.Value
}

func PrototypeNumberData() data.NumberData {
	return NewNumberData(0).ZeroValue().(*NumberData)
}

func NewNumberData(value int64) data.NumberData {
	return &NumberData{
		Value: value,
	}
}
