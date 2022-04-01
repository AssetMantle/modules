// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_Verify_Request(t *testing.T) {
	identityID := base.NewID("identityID")
	testAddress := sdkTypes.AccAddress("addr")
	testAuxiliaryRequest := NewAuxiliaryRequest(testAddress, identityID)

	require.Equal(t, auxiliaryRequest{Address: testAddress, IdentityID: identityID}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
