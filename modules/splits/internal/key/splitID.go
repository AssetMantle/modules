/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
	"strings"
)

var _ types.ID = (*SplitID)(nil)
var _ helpers.Key = (*SplitID)(nil)

func (splitID SplitID) GetStructReference() codec.ProtoMarshaler {
	return &splitID
}
func (splitID SplitID) Bytes() []byte {
	return append(
		splitID.OwnerID.Bytes(),
		splitID.OwnableID.Bytes()...)
}
func (splitID SplitID) String() string {
	var values []string
	values = append(values, splitID.OwnerID.String())
	values = append(values, splitID.OwnableID.String())

	return strings.Join(values, constants.SecondOrderCompositeIDSeparator)
}
func (splitID SplitID) Compare(id types.ID) int {
	return bytes.Compare(splitID.Bytes(), id.Bytes())
}
func (splitID SplitID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(splitID.Bytes())
}
func (SplitID) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, module.Name, SplitID{})
}
func (splitID SplitID) IsPartial() bool {
	return len(splitID.OwnableID.Bytes()) == 0
}
func (splitID SplitID) Equals(key helpers.Key) bool {
	id := splitIDFromInterface(key)
	return splitID.Compare(&id) == 0
}

func NewSplitID(ownerID types.ID, ownableID types.ID) types.ID {
	return &SplitID{
		OwnerID:   ownerID,
		OwnableID: ownableID,
	}
}
