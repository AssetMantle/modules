// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"context"
	"encoding/binary"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/types"
)

var _ types.Height = (*Height)(nil)

func (height *Height) ValidateBasic() error {
	if height.Value < -1 {
		return errorConstants.IncorrectFormat
	}
	return nil
}
func (height *Height) Bytes() []byte {
	Bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(Bytes, uint64(height.Get()))
	return Bytes
}
func (height *Height) Compare(compareHeight types.Height) int {
	if height.Get() > compareHeight.Get() {
		return 1
	} else if height.Get() < compareHeight.Get() {
		return -1
	}

	return 0
}
func (height *Height) Get() int64 { return height.Value }

func NewHeight(value int64) types.Height {
	if value < 0 {
		value = -1
	}

	return &Height{Value: value}
}

func CurrentHeight(context context.Context) types.Height {
	return NewHeight(sdkTypes.UnwrapSDKContext(context).BlockHeight())
}
