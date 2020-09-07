/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

type Parameter interface {
	String() string
	Equal(Parameter) bool
	Validate() error

	GetKey() string
	GetData() Data
	GetValidator() func(interface{}) error

	Mutate(Data) Parameter
}
