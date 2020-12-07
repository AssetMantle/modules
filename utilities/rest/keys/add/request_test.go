package add

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Add_Request(t *testing.T) {
	require.Equal(t, nil, request{Name: "name", Mnemonic: "mnemonic"}.Validate())
}
