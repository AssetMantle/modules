// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/types/base"
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
