package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Split struct {
	ID    types.ID     `json:"id" valid:"required field id missing"`
	Split sdkTypes.Dec `json:"split" valid:"required~required field split missing matches(^[0-9]$)~invalid field split"`
}

var _ mappables.Split = (*Split)(nil)

func (split Split) GetID() types.ID { return split.ID }
func (split Split) GetOwnerID() types.ID {
	return splitIDFromInterface(split.ID).OwnerID
}
func (split Split) GetOwnableID() types.ID {
	return splitIDFromInterface(split.ID).OwnableID
}
func (split Split) GetSplit() sdkTypes.Dec {
	return split.Split
}
func (split Split) Send(Split sdkTypes.Dec) traits.Transactional {
	split.Split = split.Split.Sub(Split)
	return split
}
func (split Split) Receive(Split sdkTypes.Dec) traits.Transactional {
	split.Split = split.Split.Add(Split)
	return split
}
func (split Split) CanSend(Split sdkTypes.Dec) bool {
	return split.Split.GTE(Split)
}
func (split Split) Encode() []byte {
	return packageCodec.MustMarshalBinaryBare(split)
}
func (split Split) Decode(bytes []byte) traits.Mappable {
	packageCodec.MustUnmarshalBinaryBare(bytes, &split)
	return split
}
func splitPrototype() traits.Mappable {
	return Split{}
}
func NewSplit(splitID types.ID, spl sdkTypes.Dec) mappables.Split {
	return Split{
		ID:    splitID,
		Split: spl,
	}
}
