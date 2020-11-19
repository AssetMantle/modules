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
	//require.Equal(t,parameter{ID: id,Data: data,validator: validator},testParameter)
	require.Equal(t, id, testParameter.GetID())
	//require.Equal(t,true,testParameter.Equal(testParameter))
	require.Equal(t, errors.ErrKeyIncorrect, testParameter.Validate())
	require.Equal(t, data, testParameter.GetData())

	// cannot take func as paramter
	//require.Equal(t, validator, testParameter.GetValidator())

	//Fix Bug
	//require.Equal(t,parameter{ID: id,Data: NewDecData(sdkTypes.NewDec(123)),validator: validator},testParameter.Mutate(NewDecData(sdkTypes.NewDec(123))))
}
