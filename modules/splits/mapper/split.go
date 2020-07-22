package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/trait"
)

type split struct {
	ID    schema.ID
	Split sdkTypes.Dec
}

var _ schema.Split = (*split)(nil)

func (split split) GetID() schema.ID { return split.ID }
func (split split) GetOwnerID() schema.ID {
	return splitIDFromInterface(split.ID).OwnerID
}
func (split split) GetOwnableID() schema.ID {
	return splitIDFromInterface(split.ID).OwnableID
}
func (split split) GetSplit() sdkTypes.Dec {
	return split.Split
}
func (split split) Send(Split sdkTypes.Dec) trait.Transactional {
	split.Split = split.Split.Sub(Split)
	return split
}
func (split split) Receive(Split sdkTypes.Dec) trait.Transactional {
	split.Split = split.Split.Add(Split)
	return split
}
func (split split) CanSend(Split sdkTypes.Dec) bool {
	return split.Split.GTE(Split)
}
func NewSplit(splitID schema.ID, Split sdkTypes.Dec) schema.Split {
	return split{
		ID:    splitID,
		Split: Split,
	}
}
