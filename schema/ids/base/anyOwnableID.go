package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type ownableIDGetter interface {
	get() ids.OwnableID
}

var _ ownableIDGetter = (*AnyOwnableID_AssetID)(nil)

func (m *AnyOwnableID_AssetID) get() ids.OwnableID {
	return m.AssetID
}

var _ ownableIDGetter = (*AnyOwnableID_CoinID)(nil)

func (m *AnyOwnableID_CoinID) get() ids.OwnableID {
	return m.CoinID
}

var _ ids.AnyOwnableID = (*AnyOwnableID)(nil)

func (m *AnyOwnableID) Get() ids.OwnableID {
	return m.Impl.(ownableIDGetter).get()
}
func (m *AnyOwnableID) Compare(listable traits.Listable) int {
	return m.Impl.(ownableIDGetter).get().Compare(listable)
}
func (m *AnyOwnableID) AsString() string {
	return m.Impl.(ownableIDGetter).get().AsString()
}
func (m *AnyOwnableID) Bytes() []byte {
	return m.Impl.(ownableIDGetter).get().Bytes()
}
func (m *AnyOwnableID) ToAnyID() ids.AnyID {
	return m.Impl.(ownableIDGetter).get().ToAnyID()
}
func (m *AnyOwnableID) ToAnyOwnableID() ids.AnyOwnableID {
	return m.Impl.(ownableIDGetter).get().ToAnyOwnableID()
}
func (m *AnyOwnableID) IsAnyOwnableID() {
}

func (m *AnyOwnableID) ValidateBasic() error {
	return m.Impl.(ownableIDGetter).get().ValidateBasic()
}

func PrototypeOwnableID() ids.AnyOwnableID {
	return PrototypeAssetID().ToAnyOwnableID()
}

func ReadOwnableID(ownableIDString string) (ids.AnyOwnableID, error) {
	if assetID, err := ReadAssetID(ownableIDString); err == nil {
		return assetID.ToAnyOwnableID(), nil
	} else {
		return ReadCoinID(ownableIDString).ToAnyOwnableID(), nil
	}
}
