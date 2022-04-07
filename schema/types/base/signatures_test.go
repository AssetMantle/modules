// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types"
)

func Test_Signatures(t *testing.T) {
	testSignature := NewSignature(baseIDs.NewID("ID"), baseIDs.NewID("signature").Bytes(), NewHeight(10))
	testSignature2 := NewSignature(baseIDs.NewID("ID2"), baseIDs.NewID("signature2").Bytes(), NewHeight(20))
	testSignatures := NewSignatures([]types.Signature{testSignature, testSignature2})
	require.Equal(t, signatures{SignatureList: []types.Signature{testSignature, testSignature2}}, testSignatures)
	require.Equal(t, testSignature, testSignatures.Get(baseIDs.NewID("ID")))
	require.Equal(t, nil, testSignatures.Get(baseIDs.NewID("ID3")))
	require.Equal(t, []types.Signature{testSignature, testSignature2}, testSignatures.GetList())

	newSignature := NewSignature(baseIDs.NewID("ID3"), baseIDs.NewID("signature3").Bytes(), NewHeight(30))
	require.Equal(t, signatures{SignatureList: []types.Signature{testSignature, testSignature2, newSignature}}, testSignatures.Add(newSignature))
	require.Equal(t, signatures{SignatureList: []types.Signature{testSignature}}, testSignatures.Remove(testSignature2))

	mutatedSignature := NewSignature(baseIDs.NewID("ID"), baseIDs.NewID("signatureMutated").Bytes(), NewHeight(100))
	require.Equal(t, signatures{SignatureList: []types.Signature{mutatedSignature, testSignature2}}, testSignatures.Mutate(mutatedSignature))
}
