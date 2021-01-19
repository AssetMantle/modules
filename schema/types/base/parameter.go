/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/json"

	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type parameter struct {
	ID        types.ID   `json:"id"`
	Data      types.Data `json:"data"`
	validator func(interface{}) error
}

var _ types.Parameter = (*parameter)(nil)

func (parameter parameter) String() string {
	bytes, Error := json.Marshal(parameter)
	if Error != nil {
		return Error.Error()
	}

	return string(bytes)
}
func (parameter parameter) Equal(compareParameter types.Parameter) bool {
	if compareParameter == nil {
		return false
	}

	return parameter.Data.Equal(compareParameter.GetData())
}
func (parameter parameter) Validate() error {
	return parameter.validator(parameter)
}
func (parameter parameter) GetID() types.ID {
	return parameter.ID
}
func (parameter parameter) GetData() types.Data {
	return parameter.Data
}
func (parameter parameter) GetValidator() func(interface{}) error {
	return parameter.validator
}
func (parameter parameter) Mutate(data types.Data) types.Parameter {
	parameter.Data = data
	return parameter
}

func NewParameter(id types.ID, data types.Data, validator func(interface{}) error) types.Parameter {
	return parameter{
		ID:        id,
		Data:      data,
		validator: validator,
	}
}
