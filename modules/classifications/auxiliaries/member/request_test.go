// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package member

import (
	base2 "github.com/AssetMantle/modules/schema/qualified/base"
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

func Test_Conform_Request(t *testing.T) {
	mutableProperties := base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1")))
	immutableProperties := base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2")))
	classificationID := baseIDs.NewClassificationID(base2.NewImmutables(immutableProperties), base2.NewMutables(mutableProperties))

	testAuxiliaryRequest := NewAuxiliaryRequest(classificationID, base2.NewImmutables(immutableProperties), base2.NewMutables(mutableProperties))

	require.Equal(t, auxiliaryRequest{classificationID, base2.NewImmutables(immutableProperties), base2.NewMutables(mutableProperties)}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
