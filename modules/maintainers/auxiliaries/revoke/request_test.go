// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

func Test_Revoke_Request(t *testing.T) {
	identityID := baseIDs.NewID("identityID")
	classificationID := baseIDs.NewID("classificationID")
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
