// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

type Parameter interface {
	String() string

	Equal(Parameter) bool
	Validate() error

	GetID() ID
	GetData() Data
	GetValidator() func(interface{}) error

	Mutate(Data) Parameter
}
