package base

import (
	"encoding/base64"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Signature(t *testing.T) {

	id := NewID("ID")
	signatureBytes := NewID("Temp").Bytes()
	validityHeight := NewHeight(123)
	testSignature := NewSignature(id, signatureBytes, validityHeight)

	require.Equal(t, signature{ID: id, SignatureBytes: signatureBytes, ValidityHeight: validityHeight}, testSignature)
	require.Equal(t, base64.URLEncoding.EncodeToString(signatureBytes), testSignature.String())
	require.Equal(t, signatureBytes, testSignature.Bytes())
	require.Equal(t, id, testSignature.GetID())

	require.Equal(t, validityHeight, testSignature.GetValidityHeight())
	require.Equal(t, true, testSignature.HasExpired(NewHeight(12)))

}
