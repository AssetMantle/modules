// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/types/base"
)

func Test_Validator(t *testing.T) {
	require.Equal(t, errors.IncorrectFormat, validator(base.NewID("")))
	require.Equal(t, nil, validator(Parameter))
	require.Equal(t, errors.InvalidParameter, validator(base.NewParameter(base.NewID(""), base.NewStringData(""), validator)))
}
