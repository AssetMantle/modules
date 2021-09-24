/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package revoke

import (
	"testing"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_Revoke_Request(t *testing.T) {
	identityID := base.NewID("identityID")
	classificationID := base.NewID("classificationID")
	testAuxiliaryRequest := NewAuxiliaryRequest(identityID, identityID, classificationID)

	require.Equal(t, auxiliaryRequest{
		FromID:           identityID,
		ToID:             identityID,
		ClassificationID: classificationID,
	}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))
}
