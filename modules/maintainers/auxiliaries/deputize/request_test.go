// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package deputize

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func Test_Deputize_Request(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	identityID := baseIDs.NewID("identityID")
	maintainedProperties := baseTypes.NewProperties(baseTypes.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("Data1")))

	testAuxiliaryRequest := NewAuxiliaryRequest(identityID, identityID, classificationID, maintainedProperties, false, false, false)

	require.Equal(t, testAuxiliaryRequest, auxiliaryRequest{FromID: identityID, ToID: identityID, ClassificationID: classificationID, MaintainedProperties: maintainedProperties})
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
