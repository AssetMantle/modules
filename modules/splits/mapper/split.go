/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type split struct {
	ID    types.ID     `json:"id" valid:"required field id missing"`
	Split sdkTypes.Dec `json:"split" valid:"required~required field split missing matches(^[0-9]$)~invalid field split"`
}

var _ mappables.Split = (*split)(nil)

func (split split) GetID() types.ID { return split.ID }
func (split split) GetOwnerID() types.ID {
	return splitIDFromInterface(split.ID).OwnerID
}
func (split split) GetOwnableID() types.ID {
	return splitIDFromInterface(split.ID).OwnableID
}
func (split split) GetSplit() sdkTypes.Dec {
	return split.Split
}
func (split split) Send(Split sdkTypes.Dec) traits.Transactional {
	split.Split = split.Split.Sub(Split)
	return split
}
func (split split) Receive(Split sdkTypes.Dec) traits.Transactional {
	split.Split = split.Split.Add(Split)
	return split
}
func (split split) CanSend(Split sdkTypes.Dec) bool {
	return split.Split.GTE(Split)
}
func (split split) Encode() []byte {
	return packageCodec.MustMarshalBinaryBare(split)
}
func (split split) Decode(bytes []byte) traits.Mappable {
	packageCodec.MustUnmarshalBinaryBare(bytes, &split)
	return split
}
func splitPrototype() traits.Mappable {
	return split{}
}
func NewSplit(splitID types.ID, Split sdkTypes.Dec) mappables.Split {
	return split{
		ID:    splitID,
		Split: Split,
	}
}
