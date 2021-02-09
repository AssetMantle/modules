/*
Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	moduleConstants "github.com/persistenceOne/persistenceSDK/constants/modules"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type decID struct {
	Value sdkTypes.Dec `json:"value"`
}

var _ types.ID = (*decID)(nil)

func (d decID) String() string {
	return d.Value.String()
}

// Sorts negative numbers in descending order and positive numbers in ascending order
func (d decID) Bytes() []byte {
	var Bytes []byte

	if d.Value.IsNegative() {
		Bytes = append(Bytes, moduleConstants.NegativeExchangeRate)
	} else {
		Bytes = append(Bytes, moduleConstants.PositiveExchangeRate)
	}

	Bytes = append(Bytes, uint8(len(strings.Split(d.Value.Abs().String(), ".")[0])))
	Bytes = append(Bytes, []byte(d.Value.Abs().String())...)

	return Bytes
}

func (d decID) Equals(i types.ID) bool {
	switch v := i.(type) {
	case decID:
		return d.Value.Equal(v.Value)
	default:
		return false
	}
}

func NewDecID(d sdkTypes.Dec) types.ID {
	return decID{
		Value: d,
	}
}
