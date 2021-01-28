/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mint

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_Mint_Request(t *testing.T) {

	ownerID := base.NewID("ownerID")
	ownableID := base.NewID("ownableID")
	splits := sdkTypes.NewDec(10)
	testAuxiliaryRequest := NewAuxiliaryRequest(ownerID, ownableID, splits)

	require.Equal(t, auxiliaryRequest{OwnerID: ownerID, OwnableID: ownableID, Value: splits}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
