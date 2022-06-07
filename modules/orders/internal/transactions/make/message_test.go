// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	xprtErrors "github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	"github.com/AssetMantle/modules/utilities/transaction"
)

func Test_Make_Message(t *testing.T) {

	FromID := baseIDs.NewID("fromID")
	classificationID := baseIDs.NewID("classificationID")
	makerOwnableID := baseIDs.NewID("makerOwnableID")
	takerOwnableID := baseIDs.NewID("takerOwnableID")
	expiresIn := baseTypes.NewHeight(12)
	makerOwnableSplit := sdkTypes.NewDec(2)
	takerOwnableSplit := sdkTypes.NewDec(1)
	zeroTakerOwnableSplit := sdkTypes.ZeroDec()

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	immutableMetaProperties, err := utilities.ReadMetaProperties("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, err)
	immutableProperties, err := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	mutableMetaProperties, err := utilities.ReadMetaProperties("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, err)
	mutableProperties, err := utilities.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, err)

	testMessage := newMessage(fromAccAddress, FromID, classificationID, makerOwnableID, takerOwnableID, expiresIn, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)
	require.Equal(t, message{From: fromAccAddress, FromID: FromID, ClassificationID: classificationID, MakerOwnableID: makerOwnableID, TakerOwnableID: takerOwnableID, ExpiresIn: expiresIn, TakerOwnableSplit: takerOwnableSplit, MakerOwnableSplit: makerOwnableSplit, ImmutableMetaProperties: immutableMetaProperties, ImmutableProperties: immutableProperties, MutableMetaProperties: mutableMetaProperties, MutableProperties: mutableProperties}, testMessage)
	require.Equal(t, module.Name, testMessage.Route())
	require.Equal(t, Transaction.GetName(), testMessage.Type())
	require.Equal(t, nil, testMessage.ValidateBasic())
	require.NotNil(t, message{}.ValidateBasic())
	require.Equal(t, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage)), testMessage.GetSignBytes())
	require.Equal(t, []sdkTypes.AccAddress{fromAccAddress}, testMessage.GetSigners())
	require.Equal(t, testMessage, messageFromInterface(testMessage))
	require.Equal(t, message{}, messageFromInterface(nil))
	require.Equal(t, message{}, messagePrototype())
	require.Error(t, errors.Wrap(xprtErrors.IncorrectMessage, ""), newMessage(fromAccAddress, FromID, classificationID, makerOwnableID, takerOwnableID, expiresIn, makerOwnableSplit, zeroTakerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).ValidateBasic())
	require.Error(t, errors.Wrap(xprtErrors.IncorrectMessage, ""), newMessage(fromAccAddress, FromID, classificationID, makerOwnableID, makerOwnableID, expiresIn, makerOwnableSplit, takerOwnableSplit, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).ValidateBasic())

}
