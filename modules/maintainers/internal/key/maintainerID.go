// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"bytes"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/modules/maintainers/internal/module"
	"github.com/AssetMantle/modules/schema/capabilities"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type maintainerID struct {
	ClassificationID ids.ID `json:"classificationID" valid:"required~required field classificationID missing"`
	IdentityID       ids.ID `json:"identityID" valid:"required~required field identityID missing"`
}

var _ ids.ID = (*maintainerID)(nil)
var _ helpers.Key = (*maintainerID)(nil)

func (maintainerID maintainerID) Bytes() []byte {
	return append(
		maintainerID.ClassificationID.Bytes(),
		maintainerID.IdentityID.Bytes()...)
}
func (maintainerID maintainerID) String() string {
	var values []string
	values = append(values, maintainerID.ClassificationID.String())
	values = append(values, maintainerID.IdentityID.String())

	return strings.Join(values, constants.SecondOrderCompositeIDSeparator)
}
func (maintainerID maintainerID) Compare(listable capabilities.Listable) int {
	return bytes.Compare(maintainerID.Bytes(), maintainerIDFromInterface(listable).Bytes())
}
func (maintainerID maintainerID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(maintainerID.Bytes())
}
func (maintainerID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, maintainerID{})
}
func (maintainerID maintainerID) IsPartial() bool {
	return len(maintainerID.IdentityID.Bytes()) == 0
}
func (maintainerID maintainerID) Equals(key helpers.Key) bool {
	return maintainerID.Compare(maintainerIDFromInterface(key)) == 0
}

func NewMaintainerID(classificationID ids.ID, identityID ids.ID) ids.ID {
	return maintainerID{
		ClassificationID: classificationID,
		IdentityID:       identityID,
	}
}
