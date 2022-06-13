// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/identities/internal/module"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	"github.com/AssetMantle/modules/utilities/transaction"
)

func Test_Issue_Message(t *testing.T) {
	testFromID := baseIDs.NewID("fromID")
	testClassificationID := baseIDs.NewID("classificationID")

	const fromAddress = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	const toAddress = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	var toAccAddress sdkTypes.AccAddress
	toAccAddress, err = sdkTypes.AccAddressFromBech32(toAddress)
	require.Nil(t, err)

	var immutableMetaProperties lists.MetaPropertyList
	immutableMetaProperties, err = utilities.ReadMetaProperties("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, err)

	var immutableProperties lists.PropertyList
	immutableProperties, err = utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)

	var mutableMetaProperties lists.MetaPropertyList
	mutableMetaProperties, err = utilities.ReadMetaProperties("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, err)

	var mutableProperties lists.PropertyList
	mutableProperties, err = utilities.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, err)

	testMessage := newMessage(fromAccAddress, toAccAddress, testFromID, testClassificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)
	require.Equal(t, message{From: fromAccAddress, To: toAccAddress, FromID: testFromID, ClassificationID: testClassificationID, ImmutableMetaProperties: immutableMetaProperties, ImmutableProperties: immutableProperties, MutableMetaProperties: mutableMetaProperties, MutableProperties: mutableProperties}, testMessage)
	require.Equal(t, module.Name, testMessage.Route())
	require.Equal(t, Transaction.GetName(), testMessage.Type())
	require.Equal(t, nil, testMessage.ValidateBasic())
	require.NotNil(t, message{}.ValidateBasic())
	require.Equal(t, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage)), testMessage.GetSignBytes())
	require.Equal(t, []sdkTypes.AccAddress{fromAccAddress}, testMessage.GetSigners())
	require.Equal(t, testMessage, messageFromInterface(testMessage))
	require.Equal(t, message{}, messageFromInterface(nil))
	require.Equal(t, message{}, messagePrototype())
}
