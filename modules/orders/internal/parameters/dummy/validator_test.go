/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package dummy

import (
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Validator(t *testing.T) {

	require.Equal(t, errors.IncorrectFormat, validator(base.NewID("")))
	require.Equal(t, nil, validator(Parameter))
	require.Equal(t, errors.InvalidParameter, validator(base.NewParameter(base.NewID(""), base.NewStringData(""), validator)))
}
