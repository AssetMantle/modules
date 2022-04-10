// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/crypto/openpgp/errors"

	"github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

func validator(interface{}) error {
	return errors.ErrKeyIncorrect
}

func Test_Parameter(t *testing.T) {

	id := baseIDs.NewID("ID")
	data := base.NewStringData("Data")

	testParameter := NewParameter(id, data, validator)
	require.Equal(t, id, testParameter.GetID())
	require.Equal(t, true, testParameter.Equal(testParameter))
	require.Equal(t, errors.ErrKeyIncorrect, testParameter.Validate())
	require.Equal(t, data, testParameter.GetData())

	require.Equal(t, `{"id":{"idString":"ID"},"data":{"value":"Data"}}`, testParameter.String())
	require.Equal(t, "Data2", testParameter.Mutate(base.NewStringData("Data2")).GetData().String())
	require.Equal(t, errors.ErrKeyIncorrect, testParameter.GetValidator()(nil))
}
