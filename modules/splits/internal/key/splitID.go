// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"bytes"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/modules/splits/internal/module"
	"github.com/AssetMantle/modules/schema/capabilities"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type splitID struct {
	OwnerID   ids.ID `json:"ownerID" valid:"required~required field ownerID missing"`
	OwnableID ids.ID `json:"ownableID" valid:"required~required field ownableID missing"`
}

var _ ids.ID = (*splitID)(nil)
var _ helpers.Key = (*splitID)(nil)

func (splitID splitID) Bytes() []byte {
	return append(
		splitID.OwnerID.Bytes(),
		splitID.OwnableID.Bytes()...)
}
func (splitID splitID) String() string {
	var values []string
	values = append(values, splitID.OwnerID.String())
	values = append(values, splitID.OwnableID.String())

	return strings.Join(values, constants.SecondOrderCompositeIDSeparator)
}
func (splitID splitID) Compare(listable capabilities.Listable) int {
	return bytes.Compare(splitID.Bytes(), splitIDFromInterface(listable).Bytes())
}
func (splitID splitID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(splitID.Bytes())
}
func (splitID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, splitID{})
}
func (splitID splitID) IsPartial() bool {
	return len(splitID.OwnableID.Bytes()) == 0
}
func (splitID splitID) Equals(key helpers.Key) bool {
	return splitID.Compare(splitIDFromInterface(key)) == 0
}

func NewSplitID(ownerID ids.ID, ownableID ids.ID) ids.ID {
	return splitID{
		OwnerID:   ownerID,
		OwnableID: ownableID,
	}
}
