// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

func Test_Mint_Request(t *testing.T) {
	ownerID := baseIDs.NewID("ownerID")
	ownableID := baseIDs.NewID("ownableID")
	splits := sdkTypes.NewDec(10)
	testAuxiliaryRequest := NewAuxiliaryRequest(ownerID, ownableID, splits)

	require.Equal(t, auxiliaryRequest{OwnerID: ownerID, OwnableID: ownableID, Value: splits}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
