package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/types"
)

type split struct {
	OwnerID   ids.IdentityID
	OwnableID ids.OwnableID
	Value     sdkTypes.Dec
}

var _ types.Split = (*split)(nil)

func (split split) GetOwnerID() ids.IdentityID {
	return split.OwnerID
}
func (split split) GetOwnableID() ids.OwnableID {
	return split.OwnableID
}
func (split split) GetValue() sdkTypes.Dec {
	return split.Value
}
func (split split) Send(outValue sdkTypes.Dec) types.Split {
	split.Value = split.Value.Sub(outValue)
	return split
}
func (split split) Receive(inValue sdkTypes.Dec) types.Split {
	split.Value = split.Value.Add(inValue)
	return split
}
func (split split) CanSend(outValue sdkTypes.Dec) bool {
	return split.Value.GTE(outValue)
}

func NewSplit(ownerID ids.IdentityID, ownableID ids.OwnableID, value sdkTypes.Dec) types.Split {
	return split{
		OwnerID:   ownerID,
		OwnableID: ownableID,
		Value:     value,
	}
}
