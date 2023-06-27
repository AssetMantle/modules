// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package member

import (
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/stretchr/testify/require"
)

func Test_Conform_Request(t *testing.T) {
	mutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")))
	immutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2")))
	classificationID := baseIDs.NewClassificationID(baseQualified.NewImmutables(immutableProperties), baseQualified.NewMutables(mutableProperties))

	testAuxiliaryRequest := NewAuxiliaryRequest(classificationID, baseQualified.NewImmutables(immutableProperties), baseQualified.NewMutables(mutableProperties))

	require.Equal(t, auxiliaryRequest{classificationID, baseQualified.NewImmutables(immutableProperties), baseQualified.NewMutables(mutableProperties)}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Panics(t, func() {
		auxiliaryRequestFromInterface(nil)
	})

}
