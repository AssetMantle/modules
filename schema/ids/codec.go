// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ids

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*AssetID)(nil), nil)
	codec.RegisterInterface((*ClassificationID)(nil), nil)
	codec.RegisterInterface((*DataID)(nil), nil)
	codec.RegisterInterface((*HashID)(nil), nil)
	codec.RegisterInterface((*ID)(nil), nil)
	codec.RegisterInterface((*IdentityID)(nil), nil)
	codec.RegisterInterface((*MaintainerID)(nil), nil)
	codec.RegisterInterface((*MetaID)(nil), nil)
	codec.RegisterInterface((*OrderID)(nil), nil)
	codec.RegisterInterface((*OwnableID)(nil), nil)
	codec.RegisterInterface((*PropertyID)(nil), nil)
	codec.RegisterInterface((*SplitID)(nil), nil)
	codec.RegisterInterface((*StringID)(nil), nil)
}
