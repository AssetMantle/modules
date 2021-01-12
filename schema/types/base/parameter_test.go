/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

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

	require.Equal(t, `{"id":{"idString":"ID"},"data":{"value":"Data"}}`, testParameter.String())
	require.Equal(t, "Data2", testParameter.Mutate(NewStringData("Data2")).GetData().String())
	require.Equal(t, errors.ErrKeyIncorrect, testParameter.GetValidator()(nil))
}
