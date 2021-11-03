/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type Data interface {
	proto.Message
	// Compare returns 1 if this > parameter
	// returns -1 if this < parameter
	// returns 0 if this = parameter
	Compare(Data) int

	String() string

	GetTypeID() ID

	ZeroValue() Data

	GenerateHashID() ID

	AsAccAddress() (sdkTypes.AccAddress, error)
	AsListData() (ListData, error)
	AsString() (string, error)
	AsDec() (sdkTypes.Dec, error)
	AsHeight() (Height, error)
	AsID() (ID, error)
	AsAny() (*codecTypes.Any, error)

	Get() interface{}
}
