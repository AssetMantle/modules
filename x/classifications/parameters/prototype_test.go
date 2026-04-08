// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"testing"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/parameters/bond_rate"
	"github.com/AssetMantle/modules/x/classifications/parameters/define_enabled"
	"github.com/AssetMantle/modules/x/classifications/parameters/max_property_count"
	"github.com/stretchr/testify/assert"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name      string
		want      helpers.ParameterManager
		wantError error
	}{
		{"valid", baseHelpers.NewParameterManager(bond_rate.ValidatableParameter, define_enabled.ValidatableParameter, max_property_count.ValidatableParameter), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Prototype()
			assert.Equal(t, tt.want, got, "Prototype()")
		})
	}
}
