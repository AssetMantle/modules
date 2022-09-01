package base

import (
	"bytes"
	"strconv"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
)

type orderID struct {
	ids.ClassificationID
	MakerOwnableID ids.OwnableID
	TakerOwnableID ids.OwnableID
	ExchangeRate   sdkTypes.Dec
	CreationHeight types.Height
	MakerID        ids.IdentityID
	ids.HashID
}

func (orderID orderID) IsOrderID() {
	// TODO implement me
	panic("implement me")
}

var _ ids.OrderID = (*orderID)(nil)

func (orderID orderID) Bytes() []byte {
	var Bytes []byte

	Bytes = append(Bytes, orderID.ClassificationID.Bytes()...)
	Bytes = append(Bytes, orderID.MakerOwnableID.Bytes()...)
	Bytes = append(Bytes, orderID.TakerOwnableID.Bytes()...)
	Bytes = append(Bytes, orderID.ExchangeRate.Bytes()...)
	Bytes = append(Bytes, orderID.CreationHeight.Bytes()...)
	Bytes = append(Bytes, orderID.MakerID.Bytes()...)
	Bytes = append(Bytes, orderID.HashID.Bytes()...)

	return Bytes
}
func (orderID orderID) String() string {
	return stringUtilities.JoinIDStrings(
		orderID.ClassificationID.String(),
		orderID.MakerOwnableID.String(),
		orderID.TakerOwnableID.String(),
		orderID.ExchangeRate.String(),
		orderID.CreationHeight.String(),
		orderID.MakerID.String(),
		orderID.HashID.String(),
	)
}
func (orderID orderID) Compare(listable traits.Listable) int {
	return bytes.Compare(orderID.Bytes(), orderIDFromInterface(listable).Bytes())
}
func (orderID orderID) GetHashID() ids.HashID {
	return orderID.HashID
}
func orderIDFromInterface(i interface{}) orderID {
	switch value := i.(type) {
	case orderID:
		return value
	default:
		panic(constants.MetaDataError)
	}
}

func NewOrderID(classificationID ids.ClassificationID, makerOwnableID ids.OwnableID, takerOwnableID ids.OwnableID, rate sdkTypes.Dec, creationHeight types.Height, makerID ids.IdentityID, immutables qualified.Immutables) ids.OrderID {
	return orderID{
		ClassificationID: classificationID,
		MakerOwnableID:   makerOwnableID,
		TakerOwnableID:   takerOwnableID,
		ExchangeRate:     rate,
		CreationHeight:   creationHeight,
		MakerID:          makerID,
		HashID:           immutables.GenerateHashID(),
	}
}

func PrototypeOrderID() ids.OrderID {
	return orderID{
		ClassificationID: PrototypeClassificationID(),
		MakerOwnableID:   PrototypeOwnableID(),
		TakerOwnableID:   PrototypeOwnableID(),
		ExchangeRate:     sdkTypes.ZeroDec(),
		CreationHeight:   base.NewHeight(0),
		MakerID:          PrototypeIdentityID(),
		HashID:           PrototypeHashID(),
	}
}

func ReadOrderID(orderIDString string) (ids.OrderID, error) {
	if orderIDStringSplit := stringUtilities.SplitCompositeIDString(orderIDString); len(orderIDStringSplit) == 7 {
		if classificationID, err := ReadClassificationID(orderIDStringSplit[0]); err == nil {
			if makerOwnableID, err := ReadOwnableID(orderIDStringSplit[1]); err == nil {
				if takerOwnableID, err := ReadOwnableID(orderIDStringSplit[2]); err == nil {
					if exchangeRate, err := sdkTypes.NewDecFromStr(orderIDStringSplit[3]); err == nil {
						if creationHeightDec, err := strconv.ParseInt(orderIDStringSplit[4], 10, 64); err == nil {
							if makerID, err := ReadIdentityID(orderIDStringSplit[5]); err == nil {
								if hashID, err := ReadHashID(orderIDStringSplit[6]); err == nil {
									return orderID{
										ClassificationID: classificationID,
										MakerOwnableID:   makerOwnableID,
										TakerOwnableID:   takerOwnableID,
										ExchangeRate:     exchangeRate,
										CreationHeight:   base.NewHeight(creationHeightDec),
										MakerID:          makerID,
										HashID:           hashID,
									}, nil
								}
							}
						}
					}
				}
			}
		}
	}

	if orderIDString == "" {
		return PrototypeOrderID(), nil
	}

	return orderID{}, nil
}
