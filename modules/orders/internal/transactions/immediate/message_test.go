// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package immediate

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/errors"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	xprtErrors "github.com/AssetMantle/modules/schema/errors/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"

	"github.com/AssetMantle/modules/modules/orders/internal/module"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	"github.com/AssetMantle/modules/utilities/transaction"
)

func Test_Make_Message(t *testing.T) {
	fromID := baseIDs.NewID("fromID")
	classificationID := baseIDs.NewID("classificationID")
	makerOwnableID := baseIDs.NewID("makerOwnableID")
	takerOwnableID := baseIDs.NewID("takerOwnableID")
	expiresIn := baseTypes.NewHeight(12)
	makerOwnableSplit := sdkTypes.NewDec(2)
	takerOwnableSplit, _ := sdkTypes.NewDecFromStr("2000000000000000000")
	zeroTakerOwnableSplit, _ := sdkTypes.NewDecFromStr("0")

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	immutableMetaProperties, err := utilities.ReadMetaPropertyList("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, err)
	immutableProperties, err := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	mutableMetaProperties, err := utilities.ReadMetaPropertyList("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, err)
	mutableProperties, err := utilities.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, err)

	testMessage := newMessage(fromAccAddress, fromID, classificationID, makerOwnableID, takerOwnableID, expiresIn, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)
	require.Equal(t, message{From: fromAccAddress, FromID: fromID, ClassificationID: classificationID, MakerOwnableID: makerOwnableID, TakerOwnableID: takerOwnableID, ExpiresIn: expiresIn, TakerOwnableSplit: takerOwnableSplit, MakerOwnableSplit: makerOwnableSplit, ImmutableMetaProperties: immutableMetaProperties, ImmutableProperties: immutableProperties, MutableMetaProperties: mutableMetaProperties, MutableProperties: mutableProperties}, testMessage)
	require.Equal(t, module.Name, testMessage.Route())
	require.Equal(t, Transaction.GetName(), testMessage.Type())
	require.Equal(t, nil, testMessage.ValidateBasic())
	require.NotNil(t, message{}.ValidateBasic())
	require.Equal(t, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage)), testMessage.GetSignBytes())
	require.Equal(t, []sdkTypes.AccAddress{fromAccAddress}, testMessage.GetSigners())
	require.Equal(t, testMessage, messageFromInterface(testMessage))
	require.Equal(t, message{}, messageFromInterface(nil))
	require.Equal(t, message{}, messagePrototype())
	require.Error(t, errors.Wrap(xprtErrors.IncorrectMessage, ""), newMessage(fromAccAddress, fromID, classificationID, makerOwnableID, takerOwnableID, expiresIn, makerOwnableSplit, zeroTakerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).ValidateBasic())
	require.Error(t, errors.Wrap(xprtErrors.IncorrectMessage, ""), newMessage(fromAccAddress, fromID, classificationID, makerOwnableID, makerOwnableID, expiresIn, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).ValidateBasic())

}
