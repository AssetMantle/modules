// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants/errors"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

func Test_Define_Response(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")

	testAuxiliaryResponse := newAuxiliaryResponse(classificationID, nil)
	require.Equal(t, auxiliaryResponse{Success: true, Error: nil, ClassificationID: classificationID}, testAuxiliaryResponse)
	require.Equal(t, true, testAuxiliaryResponse.IsSuccessful())
	require.Equal(t, nil, testAuxiliaryResponse.GetError())

	testAuxiliaryResponse2 := newAuxiliaryResponse(classificationID, errors.IncorrectFormat)
	require.Equal(t, auxiliaryResponse{Success: false, Error: errors.IncorrectFormat, ClassificationID: classificationID}, testAuxiliaryResponse2)
	require.Equal(t, false, testAuxiliaryResponse2.IsSuccessful())
	require.Equal(t, errors.IncorrectFormat, testAuxiliaryResponse2.GetError())

	classificationIDFromResponse, err := GetClassificationIDFromResponse(testAuxiliaryResponse)
	require.Equal(t, classificationID, classificationIDFromResponse)
	require.Equal(t, nil, err)

	classificationIDFromResponse2, err := GetClassificationIDFromResponse(testAuxiliaryResponse2)
	require.Equal(t, classificationID, classificationIDFromResponse2)
	require.Equal(t, errors.IncorrectFormat, err)

	_, err = GetClassificationIDFromResponse(nil)
	require.Equal(t, errors.InvalidRequest, err)

}
