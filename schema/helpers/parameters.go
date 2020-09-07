/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Parameters interface {
	String() string

	Validate() error
	Equal(Parameters) bool

	GetList() []types.Parameter
	GetKeyTable() subspace.KeyTable
	subspace.ParamSet

	Initialize(params.Subspace)
}
