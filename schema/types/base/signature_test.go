package base

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Signature(t *testing.T) {

	id := NewID("ID")
	signatureBytes := []byte{}
	validityHeight := NewHeight(123)
	testSignature := NewSignature(id, signatureBytes, validityHeight)

	require.Equal(t, signature{ID: id, SignatureBytes: signatureBytes, ValidityHeight: validityHeight}, testSignature)

}
