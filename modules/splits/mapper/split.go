package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/entities"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type split struct {
	ID    types.ID
	Split sdkTypes.Dec
}

var _ entities.Split = (*split)(nil)

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
func NewSplit(splitID types.ID, Split sdkTypes.Dec) entities.Split {
	return split{
		ID:    splitID,
		Split: Split,
	}
}
