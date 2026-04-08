// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"testing"

	"github.com/AssetMantle/modules/x/assets/queries/asset"
	"github.com/stretchr/testify/assert"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name      string
		want      string
		getString string
	}{

		{"valid", asset.Query.GetServicePath(), asset.Query.GetServicePath()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Prototype()
			assert.Equal(t, tt.want, got.GetQuery(tt.getString).GetServicePath())
		})
	}
}
