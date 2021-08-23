/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.Parameter = (*Parameter)(nil)

func (parameter Parameter) String() string {
	panic("failed to fetch module level implementation")
}
func (parameter Parameter) Equal(compareParameter types.Parameter) bool {
	panic("failed to fetch module level implementation")
}
func (parameter Parameter) Validate() error {
	panic("failed to fetch module level implementation")
}
func (parameter Parameter) GetID() types.ID {
	panic("failed to fetch module level implementation")
}
func (parameter Parameter) GetData() types.Data {
	panic("failed to fetch module level implementation")
}
func (parameter Parameter) GetValidator() func(interface{}) error {
	panic("failed to fetch module level implementation")
}
func (parameter Parameter) Mutate(data types.Data) types.Parameter {
	panic("failed to fetch module level implementation")
}

func NewParameter(id types.ID, data types.Data) types.Parameter {
	return &Parameter{
		ID:   id,
		Data: data,
	}
}
