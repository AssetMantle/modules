package base

import (
	"github.com/AssetMantle/modules/schema/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/types"
)

//type split struct {
//	OwnerID   ids.IdentityID
//	OwnableID ids.OwnableID
//	Value     sdkTypes.Dec
//}

var _ types.Split = (*Split)(nil)

func (split *Split) GetOwnerID() ids.IdentityID {
	return split.OwnerID
}
func (split *Split) GetOwnableID() ids.OwnableID {
	return split.OwnableID
}
func (split *Split) GetValue() sdkTypes.Dec {
	return split.Value
}
func (split *Split) Send(outValue sdkTypes.Dec) types.Split {
	split.Value = split.Value.Sub(outValue)
	return split
}
func (split *Split) Receive(inValue sdkTypes.Dec) types.Split {
	split.Value = split.Value.Add(inValue)
	return split
}
func (split *Split) CanSend(outValue sdkTypes.Dec) bool {
	return split.Value.GTE(outValue)
}

func NewSplit(ownerID ids.IdentityID, ownableID ids.OwnableID, value sdkTypes.Dec) types.Split {
	return &Split{
		OwnerID:   ownerID.(*base.IdentityID),
		OwnableID: ownableID.(*base.OwnableID),
		Value:     value,
	}
}
