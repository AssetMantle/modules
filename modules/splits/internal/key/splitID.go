// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/splits/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
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
	return stringUtilities.JoinIDStrings(splitID.OwnerID.String(), splitID.OwnableID.String())
}
func (splitID splitID) Compare(listable traits.Listable) int {
	return bytes.Compare(splitID.Bytes(), splitIDFromInterface(listable).Bytes())
}
func (splitID splitID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(splitID.Bytes())
}
func (splitID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, splitID{})
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
