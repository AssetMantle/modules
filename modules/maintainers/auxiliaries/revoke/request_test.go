/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package revoke

import (
	"testing"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_Revoke_Request(t *testing.T) {

	classificationID := base.NewID("classificationID")
	identityID := base.NewID("identityID")
	mutables := base.NewMutables(base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("Data1")))))

	testAuxiliaryRequest := NewAuxiliaryRequest(classificationID, identityID, mutables)

	require.Equal(t, auxiliaryRequest{MaintainerID: classificationID, IdentityID: identityID, MutableProperties: mutables}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
