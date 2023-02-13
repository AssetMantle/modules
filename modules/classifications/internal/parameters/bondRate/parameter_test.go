// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package bondRate

import (
	"testing"

	baseProperties "github.com/AssetMantle/modules/schema/properties/base"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/errors/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
)

func Test_Validator(t *testing.T) {
	require.Equal(t, constants.IncorrectFormat, validator(baseIDs.NewStringID("")))
	require.Equal(t, nil, validator(Parameter))
	require.Equal(t, constants.InvalidParameter, validator(baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID(""), baseData.NewStringData("")))))
}
