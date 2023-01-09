package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

// TODO rename to something more appropriate
var _ ids.CoinID = (*CoinID)(nil)

func (coinID *CoinID) IsCoinID() {
}
func (coinID *CoinID) AsString() string {
	return coinID.StringID.AsString()
}

// TODO: Verify
func (coinID *CoinID) Bytes() []byte {
	return []byte(coinID.StringID.IDString)
}
func (coinID *CoinID) IsOwnableID() {}
func (coinID *CoinID) Compare(listable traits.Listable) int {
	// TODO devise a better strategy to compare coinID and coinID
	return bytes.Compare(coinID.Bytes(), coinIDFromInterface(listable).Bytes())
}
func (coinID *CoinID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_OwnableID{
			OwnableID: coinID.ToAnyOwnableID().(*AnyOwnableID),
		},
	}
}
func (coinID *CoinID) ToAnyOwnableID() ids.AnyOwnableID {
	return &AnyOwnableID{
		Impl: &AnyOwnableID_CoinID{
			CoinID: coinID,
		},
	}
}

func coinIDFromInterface(i interface{}) ids.CoinID {
	switch value := i.(type) {
	case ids.CoinID:
		return value
	default:
		panic(constants.MetaDataError)
	}
}
func NewCoinID(stringID ids.StringID) ids.CoinID {
	return &CoinID{
		StringID: stringID.(*StringID),
	}
}

func PrototypeCoinID() ids.OwnableID {
	return &CoinID{
		StringID: PrototypeStringID().(*StringID),
	}
}

func ReadCoinID(idString string) ids.CoinID {
	// TODO ***** do not allow hashes
	return &CoinID{
		StringID: NewStringID(idString).(*StringID),
	}
}
