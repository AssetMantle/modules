/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package parameters

import (
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/parameters/dummy"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Prototype(t *testing.T) {
	prototype := Prototype()
	require.Equal(t, baseHelpers.NewParameters(dummy.Parameter).String(), prototype.String())
	require.Equal(t, nil, prototype.Validate())
	require.Equal(t, dummy.Parameter.String(), prototype.Get(base.NewID("dummy")).String())
}
