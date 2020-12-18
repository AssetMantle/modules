/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
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

type splitID struct {
	OwnerID   types.ID `json:"ownerID" valid:"required~required field ownerID missing"`
	OwnableID types.ID `json:"ownableID" valid:"required~required field ownableID missing"`
}

var _ types.ID = (*splitID)(nil)
var _ helpers.Key = (*splitID)(nil)

func (SplitID splitID) Bytes() []byte {
	return append(
		SplitID.OwnerID.Bytes(),
		SplitID.OwnableID.Bytes()...)

}
func (SplitID splitID) String() string {
	var values []string
	values = append(values, SplitID.OwnerID.String())
	values = append(values, SplitID.OwnableID.String())
	return strings.Join(values, constants.SecondOrderCompositeIDSeparator)
}
func (SplitID splitID) Equals(id types.ID) bool {
	return bytes.Compare(SplitID.Bytes(), id.Bytes()) == 0
}
func (SplitID splitID) GenerateStoreKeyBytes() []byte {
	return append([]byte{0x13}, SplitID.Bytes()...)
}
func (splitID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, splitID{})
}
func (SplitID splitID) IsPartial() bool {
	if len(SplitID.OwnableID.Bytes()) > 0 {
		return false
	}
	return true
}
func (SplitID splitID) Matches(key helpers.Key) bool {
	switch value := key.(type) {
	case splitID:
		return bytes.Compare(SplitID.Bytes(), value.Bytes()) == 0
	default:
		return false
	}
}

func New(id types.ID) helpers.Key {
	return splitIDFromInterface(id)
}

func NewSplitID(ownerID types.ID, ownableID types.ID) types.ID {
	return splitID{
		OwnerID:   ownerID,
		OwnableID: ownableID,
	}
}
