// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants/errors"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func Test_Validator(t *testing.T) {
	require.Equal(t, errors.IncorrectFormat, validator(baseIDs.NewID("")))
	require.Equal(t, nil, validator(Parameter))
	require.Equal(t, errors.InvalidParameter, validator(baseTypes.NewParameter(baseIDs.NewID(""), baseData.NewStringData(""), validator)))
}
