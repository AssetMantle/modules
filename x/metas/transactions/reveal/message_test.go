// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func Test_Reveal_Message(t *testing.T) {
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	data := "S|newData"
	newData, err := baseData.PrototypeAnyData().FromString(data)
	require.Equal(t, nil, err)

	testMessage := NewMessage(fromAccAddress, newData).(*Message)
	require.Equal(t, &Message{From: fromAccAddress.String(), Data: newData.ToAnyData().(*baseData.AnyData)}, testMessage)
	require.Equal(t, nil, testMessage.ValidateBasic())
	require.NotNil(t, (&Message{}).ValidateBasic())
	require.Equal(t, []sdkTypes.AccAddress{fromAccAddress}, testMessage.GetSigners())
	require.Equal(t, &Message{}, messagePrototype())

}
