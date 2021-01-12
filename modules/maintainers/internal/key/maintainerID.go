/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
	"strings"
)

type maintainerID struct {
	ClassificationID types.ID `json:"classificationID" valid:"required~required field classificationID missing"`
	IdentityID       types.ID `json:"identityID" valid:"required~required field identityID missing"`
}

var _ types.ID = (*maintainerID)(nil)
var _ helpers.Key = (*maintainerID)(nil)

func (MaintainerID maintainerID) Bytes() []byte {
	return append(
		MaintainerID.ClassificationID.Bytes(),
		MaintainerID.IdentityID.Bytes()...)
}
func (MaintainerID maintainerID) String() string {
	var values []string
	values = append(values, MaintainerID.ClassificationID.String())
	values = append(values, MaintainerID.IdentityID.String())
	return strings.Join(values, constants.SecondOrderCompositeIDSeparator)
}
func (MaintainerID maintainerID) Equals(id types.ID) bool {
	return bytes.Compare(MaintainerID.Bytes(), id.Bytes()) == 0
}
func (MaintainerID maintainerID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(MaintainerID.Bytes())
}
func (maintainerID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, maintainerID{})
}
func (MaintainerID maintainerID) IsPartial() bool {
	if len(MaintainerID.IdentityID.Bytes()) > 0 {
		return false
	}
	return true
}
func (MaintainerID maintainerID) Matches(key helpers.Key) bool {
	switch value := key.(type) {
	case maintainerID:
		return bytes.Compare(MaintainerID.Bytes(), value.Bytes()) == 0
	default:
		return false
	}
}

func New(id types.ID) helpers.Key {
	return maintainerIDFromInterface(id)
}

func NewMaintainerID(classificationID types.ID, identityID types.ID) types.ID {
	return maintainerID{
		ClassificationID: classificationID,
		IdentityID:       identityID,
	}
}
