// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	"testing"

	xprtErrors "github.com/AssetMantle/modules/schema/errors/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	"github.com/AssetMantle/modules/utilities/transaction"

	"github.com/AssetMantle/modules/modules/orders/internal/key"

	"github.com/cosmos/cosmos-sdk/types/errors"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/orders/internal/module"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func Test_Make_Message(t *testing.T) {

	fromID := baseIDs.NewID("fromID")
	classificationID := baseIDs.NewID("classificationID")
	makerOwnableID := baseIDs.NewID("makerOwnableID")
	takerOwnableID := baseIDs.NewID("takerOwnableID")
	expiresIn := baseTypes.NewHeight(12)
	makerOwnableSplit := sdkTypes.NewDec(2)
	makerID := baseIDs.NewID("makerID")
	rateID := baseIDs.NewID("0.11")
	creationId := baseIDs.NewID("100")
	immutableProperties, err := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	mutableMetaProperties, err := utilities.ReadMetaPropertyList("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, err)
	mutableProperties, err := utilities.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, err)
	orderID := baseIDs.NewID(key.NewOrderID(classificationID, makerOwnableID, takerOwnableID, rateID, creationId, makerID, immutableProperties).String())

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	testMessage := newMessage(fromAccAddress, fromID, orderID, sdkTypes.OneDec(), makerOwnableSplit, expiresIn, mutableMetaProperties, mutableProperties)
	require.Equal(t, message{From: fromAccAddress, FromID: fromID, OrderID: orderID, TakerOwnableSplit: sdkTypes.OneDec(), MakerOwnableSplit: makerOwnableSplit, ExpiresIn: expiresIn, MutableMetaProperties: mutableMetaProperties, MutableProperties: mutableProperties}, testMessage)
	require.Equal(t, module.Name, testMessage.Route())
	require.Equal(t, Transaction.GetName(), testMessage.Type())
	require.Equal(t, nil, testMessage.ValidateBasic())
	require.NotNil(t, message{}.ValidateBasic())
	require.Equal(t, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage)), testMessage.GetSignBytes())
	require.Equal(t, []sdkTypes.AccAddress{fromAccAddress}, testMessage.GetSigners())
	require.Equal(t, testMessage, messageFromInterface(testMessage))
	require.Equal(t, message{}, messageFromInterface(nil))
	require.Equal(t, message{}, messagePrototype())
	require.Error(t, errors.Wrap(xprtErrors.IncorrectMessage, ""), newMessage(fromAccAddress, fromID, orderID, sdkTypes.OneDec().Neg(), makerOwnableSplit, expiresIn, mutableMetaProperties, mutableProperties).ValidateBasic())
	require.Error(t, errors.Wrap(xprtErrors.IncorrectMessage, ""), newMessage(fromAccAddress, fromID, orderID, sdkTypes.OneDec(), makerOwnableSplit, baseTypes.NewHeight(-12), mutableMetaProperties, mutableProperties).ValidateBasic())

}
