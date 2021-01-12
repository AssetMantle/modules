/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package conform

import (
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Conform_Request(t *testing.T) {

	classificationID := base.NewID("classificationID")
	mutables := base.NewMutables(base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("Data1")))))
	immutables := base.NewImmutables(base.NewProperties(base.NewProperty(base.NewID("ID2"), base.NewFact(base.NewStringData("Data2")))))

	testAuxiliaryRequest := NewAuxiliaryRequest(classificationID, immutables, mutables)

	require.Equal(t, auxiliaryRequest{ClassificationID: classificationID, Immutables: immutables, Mutables: mutables}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
