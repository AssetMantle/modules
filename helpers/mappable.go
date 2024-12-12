// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Mappable interface {
	codec.ProtoMarshaler
	ValidateBasic() error
}

func ReadMappableFromIterator[T Mappable](iterator sdkTypes.Iterator, mappable T) Mappable {
	if err := mappable.Unmarshal(iterator.Value()); err != nil {
		panic(err)
	}

	return mappable
}
