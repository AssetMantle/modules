/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Signatures(t *testing.T) {

	testSignature := NewSignature(NewID("ID"), NewID("signature").Bytes(), NewHeight(10))
	testSignature2 := NewSignature(NewID("ID2"), NewID("signature2").Bytes(), NewHeight(20))
	testSignatures := NewSignatures([]types.Signature{testSignature, testSignature2})
	require.Equal(t, signatures{SignatureList: []types.Signature{testSignature, testSignature2}}, testSignatures)
	require.Equal(t, testSignature, testSignatures.Get(NewID("ID")))
	require.Equal(t, nil, testSignatures.Get(NewID("ID3")))
	require.Equal(t, []types.Signature{testSignature, testSignature2}, testSignatures.GetList())

	newSignature := NewSignature(NewID("ID3"), NewID("signature3").Bytes(), NewHeight(30))
	require.Equal(t, signatures{SignatureList: []types.Signature{testSignature, testSignature2, newSignature}}, testSignatures.Add(newSignature))
	require.Equal(t, signatures{SignatureList: []types.Signature{testSignature}}, testSignatures.Remove(testSignature2))

	mutatedSignature := NewSignature(NewID("ID"), NewID("signatureMutated").Bytes(), NewHeight(100))
	require.Equal(t, signatures{SignatureList: []types.Signature{mutatedSignature, testSignature2}}, testSignatures.Mutate(mutatedSignature))
}
