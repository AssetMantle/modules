package base

import (
	"bytes"
	"strconv"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

type orderID struct {
	ids.ClassificationID
	MakerOwnableID ids.ID
	TakerOwnableID ids.ID
	RateID         ids.ID
	CreationID     ids.ID
	MakerID        ids.ID
	ids.HashID
}

var _ ids.OrderID = (*orderID)(nil)

func (orderID orderID) Bytes() []byte {
	var Bytes []byte

	rateIDBytes, err := orderID.getRateIDBytes()
	if err != nil {
		return Bytes
	}

	creationIDBytes, err := orderID.getCreationHeightBytes()
	if err != nil {
		return Bytes
	}

	Bytes = append(Bytes, orderID.ClassificationID.Bytes()...)
	Bytes = append(Bytes, orderID.MakerOwnableID.Bytes()...)
	Bytes = append(Bytes, orderID.TakerOwnableID.Bytes()...)
	Bytes = append(Bytes, rateIDBytes...)
	Bytes = append(Bytes, creationIDBytes...)
	Bytes = append(Bytes, orderID.MakerID.Bytes()...)
	Bytes = append(Bytes, orderID.HashID.Bytes()...)

	return Bytes
}
func (orderID orderID) String() string {
	return stringUtilities.JoinIDStrings(
		orderID.ClassificationID.String(),
		orderID.MakerOwnableID.String(),
		orderID.TakerOwnableID.String(),
		orderID.RateID.String(),
		orderID.CreationID.String(),
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
func (orderID orderID) getRateIDBytes() ([]byte, error) {
	var Bytes []byte

	if orderID.RateID.String() == "" {
		return Bytes, nil
	}

	exchangeRate, err := sdkTypes.NewDecFromStr(orderID.RateID.String())
	if err != nil {
		return Bytes, err
	}

	Bytes = append(Bytes, uint8(len(strings.Split(exchangeRate.String(), ".")[0])))
	Bytes = append(Bytes, []byte(exchangeRate.String())...)

	return Bytes, err
}

func (orderID orderID) getCreationHeightBytes() ([]byte, error) {
	var Bytes []byte

	if orderID.CreationID.String() == "" {
		return Bytes, nil
	}

	height, err := strconv.ParseInt(orderID.CreationID.String(), 10, 64)
	if err != nil {
		return Bytes, err
	}

	Bytes = append(Bytes, uint8(len(orderID.CreationID.String())))
	Bytes = append(Bytes, []byte(strconv.FormatInt(height, 10))...)

	return Bytes, err
}
func orderIDFromInterface(i interface{}) orderID {
	switch value := i.(type) {
	case orderID:
		return value
	default:
		panic(constants.MetaDataError)
	}
}

func NewOrderID(classificationID ids.ClassificationID, makerOwnableID ids.ID, takerOwnableID ids.ID, rateID ids.ID, creationID ids.ID, makerID ids.ID, immutables qualified.Immutables) ids.OrderID {
	return orderID{
		ClassificationID: classificationID,
		MakerOwnableID:   makerOwnableID,
		TakerOwnableID:   takerOwnableID,
		RateID:           rateID,
		CreationID:       creationID,
		MakerID:          makerID,
		HashID:           immutables.GenerateHashID(),
	}
}

func ReadOrderID(orderIDString string) ids.OrderID {

}
