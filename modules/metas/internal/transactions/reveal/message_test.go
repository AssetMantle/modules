// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/metas/internal/module"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/data/utilities"
	"github.com/AssetMantle/modules/utilities/transaction"
)

func Test_Reveal_Message(t *testing.T) {
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	data := "S|newData"
	newData, err := utilities.ReadData(data)
	require.Equal(t, nil, err)

	testMessage := newMessage(fromAccAddress, newData).(*Message)
	require.Equal(t, &Message{From: fromAccAddress.String(), Data: newData.ToAnyData().(*baseData.AnyData)}, testMessage)
	require.Equal(t, module.Name, testMessage.Route())
	require.Equal(t, Transaction.GetName(), testMessage.Type())
	require.Equal(t, nil, testMessage.ValidateBasic())
	require.NotNil(t, (&Message{}).ValidateBasic())
	require.Equal(t, sdkTypes.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(testMessage)), testMessage.GetSignBytes())
	require.Equal(t, []sdkTypes.AccAddress{fromAccAddress}, testMessage.GetSigners())
	require.Equal(t, testMessage, messageFromInterface(testMessage))
	require.Equal(t, &Message{}, messageFromInterface(nil))
	require.Equal(t, &Message{}, messagePrototype())

}
