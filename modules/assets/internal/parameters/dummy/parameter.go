/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package dummy

import (
	"encoding/json"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

var _ types.Parameter = (*DummyParameter)(nil)

func (dummyParameter DummyParameter) String() string {
	bytes, Error := json.Marshal(dummyParameter.BaseParameter)
	if Error != nil {
		return Error.Error()
	}
	return string(bytes)
}

func (dummyParameter DummyParameter) Equal(compareParameter types.Parameter) bool {
	if compareParameter == nil {
		return false
	}
	return dummyParameter.BaseParameter.Data.Compare(compareParameter.GetData()) == 0
}

func (dummyParameter DummyParameter) Validate() error {
	return validator(dummyParameter)
}

func (dummyParameter DummyParameter) GetID() types.ID {
	return dummyParameter.BaseParameter.ID
}

func (dummyParameter DummyParameter) GetData() types.Data {
	return dummyParameter.BaseParameter.Data
}

func (dummyParameter DummyParameter) GetValidator() func(interface{}) error {
	return validator
}

func (dummyParameter DummyParameter) Mutate(data types.Data) types.Parameter {
	dummyParameter.BaseParameter.Data = data
	return &dummyParameter
}

func (DummyParameter) RegisterImplementation(interfaceRegistry codecTypes.InterfaceRegistry){
	interfaceRegistry.RegisterImplementations((*types.Parameter)(nil), &DummyParameter{})
}

func NewParameter(id types.ID, data types.Data) types.Parameter {
	return &DummyParameter{
		BaseParameter: base.Parameter{
			ID:   id,
			Data: data,
		},
	}
}
