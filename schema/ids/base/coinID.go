package base

import (
	"bytes"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/constants"
	"github.com/AssetMantle/modules/schema/traits"
	"strings"
)

var _ ids.CoinID = (*CoinID)(nil)

func (coinID *CoinID) ValidateBasic() error {
	return coinID.StringID.ValidateBasic()
}
func (coinID *CoinID) IsCoinID() {
}
func (coinID *CoinID) GetTypeID() ids.StringID {
	return NewStringID(constants.CoinIDType)
}
func (coinID *CoinID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeCoinID(), nil
	}

	if stringID, err := PrototypeStringID().FromString(idString); err != nil {
		return PrototypeCoinID(), err
	} else {
		return &CoinID{
			StringID: stringID.(*StringID),
		}, nil
	}
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
	compareID, err := idFromListable(listable)
	if err != nil {
		panic(err)
	}
	return bytes.Compare(coinID.Bytes(), compareID.Bytes())
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
