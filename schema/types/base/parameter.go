/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"fmt"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type parameter struct {
	key       string
	data      types.Data
	validator func(interface{}) error
}

var _ types.Parameter = (*parameter)(nil)

func (parameter parameter) String() string {
	return fmt.Sprintf("%v:%v", parameter.key, parameter.data.String())
}

func (parameter parameter) Equal(Parameter types.Parameter) bool {
	return parameter.data.Equal(Parameter.GetData())
}

func (parameter parameter) Validate() error {
	return parameter.validator(parameter)
}

func (parameter parameter) GetKey() string {
	return parameter.key
}

func (parameter parameter) GetData() types.Data {
	return parameter.data
}

func (parameter parameter) GetValidator() func(interface{}) error {
	return parameter.validator
}

func (parameter parameter) Mutate(data types.Data) types.Parameter {
	parameter.data = data
	return parameter
}

func NewParameter(key string, data types.Data, validator func(interface{}) error) types.Parameter {
	return parameter{
		key:       key,
		data:      data,
		validator: validator,
	}
}
