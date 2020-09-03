/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package parameters

import (
	"bytes"
	"encoding/json"
	"fmt"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type parameters struct {
	Dummy sdkTypes.Dec `json:"dummy" valid:"required~required field fromID missing"`
}

var _ types.Parameters = (*parameters)(nil)

func (parameters parameters) String() string {
	Bytes, Error := json.Marshal(parameters)
	if Error != nil {
		panic(Error)
	} else {
		return string(Bytes)
	}
}

func (parameters parameters) Validate() error {
	if err := validateDummy(parameters.Dummy); err != nil {
		return err
	}
	return nil
}

func (parameters parameters) Equal(Parameters types.Parameters) bool {
	Bytes, Error := json.Marshal(parameters)
	if Error != nil {
		panic(Error)
	}
	CompareBytes, Error := json.Marshal(Parameters)
	if Error != nil {
		panic(Error)
	}
	return bytes.Compare(Bytes, CompareBytes) == 0
}

func (parameters parameters) ParamSetPairs() subspace.ParamSetPairs {
	return subspace.ParamSetPairs{
		params.NewParamSetPair(KeyDummy, &parameters.Dummy, validateDummy),
	}
}

func (parameters parameters) KeyTable() subspace.KeyTable {
	return subspace.NewKeyTable().RegisterParamSet(Prototype())
}

func newParameters(dummy sdkTypes.Dec) types.Parameters {
	return parameters{
		Dummy: dummy,
	}
}
func Prototype() types.Parameters {
	return parameters{}
}
func validateDummy(i interface{}) error {
	v, ok := i.(sdkTypes.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsZero() {
		return fmt.Errorf("invalid dummy %d", v)
	}
	return nil
}
