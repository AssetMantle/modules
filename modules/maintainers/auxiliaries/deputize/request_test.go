// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_Deputize_Request(t *testing.T) {
	classificationID := base.NewID("classificationID")
	identityID := base.NewID("identityID")
	maintainedProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewStringData("Data1")))

	testAuxiliaryRequest := NewAuxiliaryRequest(identityID, identityID, classificationID, maintainedProperties, false, false, false)

	require.Equal(t, testAuxiliaryRequest, auxiliaryRequest{FromID: identityID, ToID: identityID, ClassificationID: classificationID, MaintainedProperties: maintainedProperties})
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
