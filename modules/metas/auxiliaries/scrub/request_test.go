/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package scrub

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_Scrub_Request(t *testing.T) {

	metaProperty := base.NewMetaProperty(base.NewID("id"), base.NewStringData("Data"))
	testAuxiliaryRequest := NewAuxiliaryRequest(metaProperty)

	require.Equal(t, auxiliaryRequest{MetaPropertyList: []types.MetaProperty{metaProperty}}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
