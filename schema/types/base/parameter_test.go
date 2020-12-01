package base

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/crypto/openpgp/errors"
	"testing"
)

func validator(interface{}) error {
	return errors.ErrKeyIncorrect
}

func Test_Parameter(t *testing.T) {

	id := NewID("ID")
	data := NewStringData("Data")

	testParameter := NewParameter(id, data, validator)
	require.Equal(t, id, testParameter.GetID())
	require.Equal(t, true, testParameter.Equal(testParameter))
	require.Equal(t, errors.ErrKeyIncorrect, testParameter.Validate())
	require.Equal(t, data, testParameter.GetData())
}
