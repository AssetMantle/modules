// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"github.com/AssetMantle/modules/modules/assets/internal/module"
	"github.com/AssetMantle/modules/modules/assets/internal/parameters/burnEnabled"
	"github.com/AssetMantle/modules/modules/assets/internal/parameters/renumerateEnabled"
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/schema/helpers"

	"github.com/AssetMantle/modules/modules/assets/internal/parameters/mintEnabled"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name      string
		want      helpers.ParameterManager
		wantError error
	}{
		{"+ve", baseHelpers.NewParameterManager(module.Name, burnEnabled.ValidatableParameter, mintEnabled.ValidatableParameter, renumerateEnabled.ValidatableParameter), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
