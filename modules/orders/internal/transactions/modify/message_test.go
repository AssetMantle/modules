/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package modify

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/utilities/transaction"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Make_Message(t *testing.T) {

	fromID := base.NewID("fromID")
	classificationID := base.NewID("classificationID")
	makerOwnableID := base.NewID("makerOwnableID")
	takerOwnableID := base.NewID("takerOwnableID")
	expiresIn := base.NewHeight(12)
	makerOwnableSplit := sdkTypes.NewDec(2)
	makerID := base.NewID("makerID")
	rateID := base.NewID("0.11")
	creationId := base.NewID("100")
	immutableProperties, Error := base.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, Error)
	mutableMetaProperties, Error := base.ReadMetaProperties("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, Error)
	mutableProperties, Error := base.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, Error)
	orderID := base.NewID(key.NewOrderID(classificationID, makerOwnableID, takerOwnableID, rateID, creationId, makerID, immutableProperties).String())

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, Error)

	testMessage := newMessage(fromAccAddress, fromID, orderID, sdkTypes.OneDec(), makerOwnableSplit, expiresIn, mutableMetaProperties, mutableProperties)
	require.Equal(t, Message{From: fromAccAddress, FromID: fromID, OrderID: orderID, TakerOwnableSplit: sdkTypes.OneDec(), MakerOwnableSplit: makerOwnableSplit, ExpiresIn: expiresIn, MutableMetaProperties: mutableMetaProperties, MutableProperties: mutableProperties}, testMessage)
	require.Equal(t, module.Name, testMessage.Route())
	require.Equal(t, Transaction.GetName(), testMessage.Type())
	require.Equal(t, nil, testMessage.ValidateBasic())
	require.NotNil(t, Message{}.ValidateBasic())
	require.Equal(t, sdkTypes.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(testMessage)), testMessage.GetSignBytes())
	require.Equal(t, []sdkTypes.AccAddress{fromAccAddress}, testMessage.GetSigners())
	require.Equal(t, testMessage, messageFromInterface(testMessage))
	require.Equal(t, Message{}, messageFromInterface(nil))
	require.Equal(t, Message{}, messagePrototype())
	require.Error(t, errors.Wrap(xprtErrors.IncorrectMessage, ""), newMessage(fromAccAddress, fromID, orderID, sdkTypes.OneDec().Neg(), makerOwnableSplit, expiresIn, mutableMetaProperties, mutableProperties).ValidateBasic())
	require.Error(t, errors.Wrap(xprtErrors.IncorrectMessage, ""), newMessage(fromAccAddress, fromID, orderID, sdkTypes.OneDec(), makerOwnableSplit, base.NewHeight(-12), mutableMetaProperties, mutableProperties).ValidateBasic())

}
