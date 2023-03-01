// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ids

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	legacyAmino.RegisterInterface((*AssetID)(nil), nil)
	legacyAmino.RegisterInterface((*ClassificationID)(nil), nil)
	legacyAmino.RegisterInterface((*CoinID)(nil), nil)
	legacyAmino.RegisterInterface((*DataID)(nil), nil)
	legacyAmino.RegisterInterface((*HashID)(nil), nil)
	legacyAmino.RegisterInterface((*ID)(nil), nil)
	legacyAmino.RegisterInterface((*IdentityID)(nil), nil)
	legacyAmino.RegisterInterface((*MaintainerID)(nil), nil)
	legacyAmino.RegisterInterface((*OrderID)(nil), nil)
	legacyAmino.RegisterInterface((*OwnableID)(nil), nil)
	legacyAmino.RegisterInterface((*PropertyID)(nil), nil)
	legacyAmino.RegisterInterface((*SplitID)(nil), nil)
	legacyAmino.RegisterInterface((*StringID)(nil), nil)
}
