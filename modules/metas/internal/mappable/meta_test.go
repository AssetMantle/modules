// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/metas/internal/key"
	"github.com/AssetMantle/modules/schema/data/base"
)

func Test_Meta_Methods(t *testing.T) {
	data := base.NewStringData("Data")

	testMeta := NewMeta(data).(meta)
	require.Equal(t, meta{ID: key.NewMetaID(data.GetType(), data.GenerateHash()), Data: data}, testMeta)
	require.Equal(t, data, testMeta.GetData())
	require.Equal(t, key.NewMetaID(data.GetType(), data.GenerateHash()), testMeta.GetKey())
	require.Equal(t, key.GenerateMetaID(data), testMeta.GetID())

}
