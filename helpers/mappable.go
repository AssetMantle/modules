// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	storeTypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

type Mappable interface {
	codec.ProtoMarshaler
	ValidateBasic() error
}

func ReadMappableFromIterator[T Mappable](iterator storeTypes.Iterator, mappable T) Mappable {
	if err := mappable.Unmarshal(iterator.Value()); err != nil {
		panic(err)
	}

	return mappable
}
