/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import "github.com/cosmos/cosmos-sdk/x/params/subspace"

type Parameters interface {
	String() string
	Validate() error
	Equal(Parameters) bool
	KeyTable() subspace.KeyTable
	subspace.ParamSet
}
