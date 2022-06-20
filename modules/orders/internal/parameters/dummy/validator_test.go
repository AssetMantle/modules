// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
)

func Test_Validator(t *testing.T) {
	require.Equal(t, constants.IncorrectFormat, validator(baseIDs.NewID("")))
	require.Equal(t, nil, validator(Parameter))
	require.Equal(t, constants.InvalidParameter, validator(baseTypes.NewParameter(baseIDs.NewID(""), baseData.NewStringData(""), validator)))
}
