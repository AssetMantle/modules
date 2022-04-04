// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintain

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/types/base"
)

func Test_Maintain_Request(t *testing.T) {
	classificationID := base.NewID("classificationID")
	identityID := base.NewID("identityID")
	mutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewStringData("Data1")))

	testAuxiliaryRequest := NewAuxiliaryRequest(classificationID, identityID, mutableProperties)

	require.Equal(t, auxiliaryRequest{ClassificationID: classificationID, IdentityID: identityID, MaintainedProperties: mutableProperties}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
