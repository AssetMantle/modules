/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

var _ types.ID = (*MaintainerID)(nil)
var _ helpers.Key = (*MaintainerID)(nil)

func (maintainerID MaintainerID) Bytes() []byte {
	return append(
		maintainerID.ClassificationID.Bytes(),
		maintainerID.IdentityID.Bytes()...)
}
func (maintainerID MaintainerID) String() string {
	var values []string
	values = append(values, maintainerID.ClassificationID.String())
	values = append(values, maintainerID.IdentityID.String())

	return strings.Join(values, constants.SecondOrderCompositeIDSeparator)
}
func (maintainerID MaintainerID) Compare(id types.ID) int {
	return bytes.Compare(maintainerID.Bytes(), id.Bytes())
}
func (maintainerID MaintainerID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(maintainerID.Bytes())
}
func (MaintainerID) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, MaintainerID{})
}
func (maintainerID MaintainerID) IsPartial() bool {
	return len(maintainerID.IdentityID.Bytes()) == 0
}
func (maintainerID MaintainerID) Equals(key helpers.Key) bool {
	return maintainerID.Compare(maintainerIDFromInterface(key)) == 0
}

func NewMaintainerID(classificationID types.ID, identityID types.ID) types.ID {
	return &MaintainerID{
		ClassificationID: classificationID,
		IdentityID:       identityID,
	}
}
