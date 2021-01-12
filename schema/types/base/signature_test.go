/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/base64"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"testing"
)

func Test_Signature(t *testing.T) {

	privateKey := ed25519.GenPrivKey()
	pubKey := privateKey.PubKey()
	signatureBytes := NewID("Temp").Bytes()

	signedBytes, err := privateKey.Sign(signatureBytes)
	require.Nil(t, err)
	id := NewID("ID")
	validityHeight := NewHeight(123)
	testSignature := NewSignature(id, signedBytes, validityHeight)

	require.Equal(t, signature{ID: id, SignatureBytes: signedBytes, ValidityHeight: validityHeight}, testSignature)
	require.Equal(t, base64.URLEncoding.EncodeToString(signedBytes), testSignature.String())
	require.Equal(t, signedBytes, testSignature.Bytes())
	require.Equal(t, id, testSignature.GetID())

	require.Equal(t, validityHeight, testSignature.GetValidityHeight())
	require.Equal(t, true, testSignature.HasExpired(NewHeight(12)))

	require.Equal(t, false, testSignature.Verify(pubKey, []byte{}))
	require.Equal(t, true, testSignature.Verify(pubKey, signatureBytes))

}
