/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/qualified"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	baseTypes "github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
	"strconv"
	"strings"
)

var _ types.ID = (*OrderID)(nil)
var _ helpers.Key = (*OrderID)(nil)

func (orderID OrderID) GetStructReference() codec.ProtoMarshaler {
	return &orderID
}
func (orderID OrderID) Bytes() []byte {
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
func (orderID OrderID) String() string {
	var values []string
	values = append(values, orderID.ClassificationID.String())
	values = append(values, orderID.MakerOwnableID.String())
	values = append(values, orderID.TakerOwnableID.String())
	values = append(values, orderID.RateID.String())
	values = append(values, orderID.CreationID.String())
	values = append(values, orderID.MakerID.String())
	values = append(values, orderID.HashID.String())

	return strings.Join(values, constants.SecondOrderCompositeIDSeparator)
}
func (orderID OrderID) Compare(id types.ID) int {
	return bytes.Compare(orderID.Bytes(), id.Bytes())
}
func (orderID OrderID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(orderID.Bytes())
}
func (OrderID) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterLegacyAminoXPRTConcrete(codec, module.Name, OrderID{})
}
func (orderID OrderID) IsPartial() bool {
	return len(orderID.HashID.Bytes()) == 0
}
func (orderID OrderID) Equals(key helpers.Key) bool {
	id := orderIDFromInterface(key)
	return orderID.Compare(&id) == 0
}

func (orderID OrderID) getRateIDBytes() ([]byte, error) {
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

func (orderID OrderID) getCreationHeightBytes() ([]byte, error) {
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

func NewOrderID(classificationID types.ID, makerOwnableID types.ID, takerOwnableID types.ID, rateID types.ID, creationID types.ID, makerID types.ID, immutableProperties types.Properties) types.ID {
	return &OrderID{
		ClassificationID: classificationID,
		MakerOwnableID:   makerOwnableID,
		TakerOwnableID:   takerOwnableID,
		RateID:           rateID,
		CreationID:       creationID,
		MakerID:          makerID,
		HashID:           baseTraits.HasImmutables{Properties: *baseTypes.NewProperties(immutableProperties.GetList()...)}.GenerateHashID(),
	}
}
