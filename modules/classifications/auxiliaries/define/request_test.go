/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_Define_Request(t *testing.T) {

	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID2"), base.NewStringData("Data2")))
	mutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewStringData("Data1")))

	testAuxiliaryRequest := NewAuxiliaryRequest(immutableProperties, mutableProperties)

	require.Equal(t, auxiliaryRequest{ImmutableProperties: immutableProperties, MutableProperties: mutableProperties}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
