package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ data.RoyaltyData = (*RoyaltyData)(nil)

func (m *RoyaltyData) GetID() ids.DataID {
	//TODO implement me
	panic("implement me")
}

func (m *RoyaltyData) GetBondWeight() int64 {
	//TODO implement me
	panic("implement me")
}

func (m *RoyaltyData) AsString() string {
	//TODO implement me
	panic("implement me")
}

func (m *RoyaltyData) FromString(s string) (data.Data, error) {
	//TODO implement me
	panic("implement me")
}

func (m *RoyaltyData) Bytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (m *RoyaltyData) GetType() ids.StringID {
	//TODO implement me
	panic("implement me")
}

func (m *RoyaltyData) ZeroValue() data.Data {
	//TODO implement me
	panic("implement me")
}

func (m *RoyaltyData) GenerateHashID() ids.HashID {
	//TODO implement me
	panic("implement me")
}

func (m *RoyaltyData) ToAnyData() data.AnyData {
	//TODO implement me
	panic("implement me")
}

func (m *RoyaltyData) Compare(listable traits.Listable) int {
	//TODO implement me
	panic("implement me")
}

func (m *RoyaltyData) GetOwnableID() ids.OwnableID {
	//TODO implement me
	panic("implement me")
}

func (m *RoyaltyData) GetIdentityID() ids.IdentityID {
	//TODO implement me
	panic("implement me")
}

func (m *RoyaltyData) GetSplit() []data.DecData {
	//TODO implement me
	panic("implement me")
}
