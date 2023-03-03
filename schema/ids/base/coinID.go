package base

import (
	"bytes"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.CoinID = (*CoinID)(nil)

func (coinID *CoinID) ValidateBasic() error {
	return coinID.StringID.ValidateBasic()
}
func (coinID *CoinID) IsCoinID() {
}
func (coinID *CoinID) GetTypeID() ids.StringID {
}
func (coinID *CoinID) FromString(idTypeAndValueString string) (ids.ID, error) {
	idTypeString, idString := splitIDTypeAndValueStrings(idTypeAndValueString)

	if idTypeString != coinID.GetTypeID().AsString() {
		return PrototypeCoinID(), errorConstants.IncorrectFormat.Wrapf("expected id type %s, got %s", coinID.GetTypeID().AsString(), idTypeString)
	}

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
	return joinIDTypeAndValueStrings(coinID.GetTypeID().AsString(), coinID.StringID.AsString())
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
		Impl: &AnyID_AnyOwnableID{
			AnyOwnableID: coinID.ToAnyOwnableID().(*AnyOwnableID),
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
