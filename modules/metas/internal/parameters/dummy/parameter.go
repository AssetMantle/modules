/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package dummy

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

var _ types.Parameter = (*DummyParameter)(nil)

func (dummyParameter DummyParameter) String() string {
	bytes, Error := json.Marshal(dummyParameter)
	if Error != nil {
		return Error.Error()
	}
	return string(bytes)
}

func (dummyParameter DummyParameter) Equal(compareParameter types.Parameter) bool {
	if compareParameter == nil {
		return false
	}
	return dummyParameter.Data.Compare(compareParameter.GetData()) == 0
}

func (dummyParameter DummyParameter) Validate() error {
	return validator(dummyParameter)
}

func (dummyParameter DummyParameter) GetID() types.ID {
	return &dummyParameter.ID
}

func (dummyParameter DummyParameter) GetData() types.Data {
	return &dummyParameter.Data
}

func (dummyParameter DummyParameter) GetValidator() func(interface{}) error {
	return validator
}

func (dummyParameter DummyParameter) Mutate(data types.Data) types.Parameter {
	decData, _ := data.AsDec()
	dummyParameter.Data = *base.NewDummyDecData(decData)
	return &dummyParameter
}

func NewParameter(id types.ID, data types.Data) *DummyParameter {
	decData, _ := data.AsDec()
	return &DummyParameter{
		ID:   *base.NewID(id.String()),
		Data: *base.NewDummyDecData(decData),
	}
}
