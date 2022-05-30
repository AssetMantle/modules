// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
)

func Test_Signatures(t *testing.T) {
	testSignature := base.NewSignature(baseIDs.NewID("ID"), baseIDs.NewID("signature").Bytes(), base.NewHeight(10))
	testSignature2 := base.NewSignature(baseIDs.NewID("ID2"), baseIDs.NewID("signature2").Bytes(), base.NewHeight(20))
	testSignatures := NewSignatures([]types.Signature{testSignature, testSignature2})
	require.Equal(t, signatures{SignatureList: []types.Signature{testSignature, testSignature2}}, testSignatures)
	require.Equal(t, testSignature, testSignatures.Get(baseIDs.NewID("ID")))
	require.Equal(t, nil, testSignatures.Get(baseIDs.NewID("ID3")))
	require.Equal(t, []types.Signature{testSignature, testSignature2}, testSignatures.GetList())

	newSignature := base.NewSignature(baseIDs.NewID("ID3"), baseIDs.NewID("signature3").Bytes(), base.NewHeight(30))
	require.Equal(t, signatures{SignatureList: []types.Signature{testSignature, testSignature2, newSignature}}, testSignatures.Add(newSignature))
	require.Equal(t, signatures{SignatureList: []types.Signature{testSignature}}, testSignatures.Remove(testSignature2))

	mutatedSignature := base.NewSignature(baseIDs.NewID("ID"), baseIDs.NewID("signatureMutated").Bytes(), base.NewHeight(100))
	require.Equal(t, signatures{SignatureList: []types.Signature{mutatedSignature, testSignature2}}, testSignatures.Mutate(mutatedSignature))
}
