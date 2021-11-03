package types

import (
	"encoding/json"
	"fmt"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

var (
	_ types.Parameter    = (*Parameter)(nil)
	_ helpers.Parameters = (*Parameters)(nil)
)

func (parameter Parameter) String() string {
	bytes, err := json.Marshal(parameter)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
func (parameter Parameter) Equal(compareParameter types.Parameter) bool {
	if compareParameter == nil {
		return false
	}

	return parameter.Data.Compare(compareParameter.GetData()) == 0
}
func (parameter Parameter) GetID() types.ID {
	return &parameter.ID
}
func (parameter Parameter) GetData() types.Data {
	data, ok := parameter.Data.GetCachedValue().(types.Data)
	if !ok {
		return nil
	}
	return data
}
func (parameter *Parameter) SetData(data types.Data) error {
	if data == nil {
		parameter.Data = codecTypes.Any{}
		return nil
	}
	any, err := codecTypes.NewAnyWithValue(data)
	if err == nil {
		parameter.Data = *any
	}
	return err
}

func (parameter Parameter) UnpackInterfaces(unpacker codecTypes.AnyUnpacker) error {
	var data types.Data
	return unpacker.UnpackAny(&parameter.Data, &data)
}

func NewParameter(id base.ID, data types.Data) Parameter {
	p := Parameter{ID: id}
	err := p.SetData(data)
	if err != nil {
		panic(err)
	}
	return p
}

var (
	MaxStringLengthID          = base.NewID("MaxStringLength")
	DefaultMaxStringLengthData = base.NewHeightData(base.NewHeight(512))
	DefaultMaxStringLength     = NewParameter(MaxStringLengthID, DefaultMaxStringLengthData)
)

func NewParameters(parameters ...Parameter) Parameters {
	return Parameters{Value: parameters}
}

func (Parameters) Validate() error {
	return nil
}

func ParamsKeyTable() paramTypes.KeyTable {
	return paramTypes.NewKeyTable().RegisterParamSet(&Parameters{})
}

func (Parameters) ParamSetPairs() paramTypes.ParamSetPairs {
	return paramTypes.ParamSetPairs{
		paramTypes.NewParamSetPair(MaxStringLengthID.Bytes(), &DefaultMaxStringLengthData, validateMaxStringLength),
	}
}

func validateMaxStringLength(i interface{}) error {
	v, ok := i.(*base.HeightData)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.Value.Value == 0 {
		return fmt.Errorf("invalid MaxStringLength: %d", v.Value.Value)
	}

	return nil
}

var (
	DefaultParameters = NewParameters(DefaultMaxStringLength)
)
