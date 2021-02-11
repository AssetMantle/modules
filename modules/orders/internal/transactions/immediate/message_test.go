/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package immediate

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/errors"
	xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/utilities/transaction"
	"github.com/stretchr/testify/require"
)

func Test_Make_Message(t *testing.T) {

	testFromID := base.NewID("fromID")
	testClassificationID := base.NewID("classificationID")
	testMakerOwnableID := base.NewID("makerOwnableID")
	testTakerOwnableID := base.NewID("takerOwnableID")
	testExpiresIn := base.NewHeight(12)
	testMakerOwnableSplit := sdkTypes.NewDec(2)
	testExchangeRate, _ := sdkTypes.NewDecFromStr("2000000000000000000")
	zeroExchangeRate, _ := sdkTypes.NewDecFromStr("0")

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, Error)

	immutableMetaProperties, Error := base.ReadMetaProperties("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, Error)
	immutableProperties, Error := base.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, Error)
	mutableMetaProperties, Error := base.ReadMetaProperties("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, Error)
	mutableProperties, Error := base.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, Error)

	testMessage := newMessage(fromAccAddress, testFromID, testClassificationID, testMakerOwnableID, testTakerOwnableID, testExpiresIn, testMakerOwnableSplit, testExchangeRate, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)
	require.Equal(t, message{From: fromAccAddress, FromID: testFromID, ClassificationID: testClassificationID, MakerOwnableID: testMakerOwnableID, TakerOwnableID: testTakerOwnableID, ExpiresIn: testExpiresIn, ExchangeRate: testExchangeRate, MakerOwnableSplit: testMakerOwnableSplit, ImmutableMetaProperties: immutableMetaProperties, ImmutableProperties: immutableProperties, MutableMetaProperties: mutableMetaProperties, MutableProperties: mutableProperties}, testMessage)
	require.Equal(t, module.Name, testMessage.Route())
	require.Equal(t, Transaction.GetName(), testMessage.Type())
	require.Equal(t, nil, testMessage.ValidateBasic())
	require.NotNil(t, message{}.ValidateBasic())
	require.Equal(t, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage)), testMessage.GetSignBytes())
	require.Equal(t, []sdkTypes.AccAddress{fromAccAddress}, testMessage.GetSigners())
	require.Equal(t, testMessage, messageFromInterface(testMessage))
	require.Equal(t, message{}, messageFromInterface(nil))
	require.Equal(t, message{}, messagePrototype())
	require.Error(t, errors.Wrap(xprtErrors.IncorrectMessage, ""), newMessage(fromAccAddress, testFromID, testClassificationID, testMakerOwnableID, testTakerOwnableID, testExpiresIn, testMakerOwnableSplit, zeroExchangeRate, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).ValidateBasic())
	require.Error(t, errors.Wrap(xprtErrors.IncorrectMessage, ""), newMessage(fromAccAddress, testFromID, testClassificationID, testMakerOwnableID, testMakerOwnableID, testExpiresIn, testMakerOwnableSplit, testExchangeRate, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).ValidateBasic())

}
