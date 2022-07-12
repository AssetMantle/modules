// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

func Test_Burn_Request(t *testing.T) {
	ownerID := baseIDs.NewStringID("ownerID")
	ownableID := baseIDs.NewStringID("ownableID")
	testValue := sdkTypes.NewDec(10)
	testAuxiliaryRequest := NewAuxiliaryRequest(ownerID, ownableID, testValue)

	require.Equal(t, auxiliaryRequest{OwnerID: ownerID, OwnableID: ownableID, Value: testValue}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
