// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/x/assets/module"
	"github.com/AssetMantle/modules/x/assets/parameters/burnEnabled"
	"github.com/AssetMantle/modules/x/assets/parameters/mintEnabled"
	"github.com/AssetMantle/modules/x/assets/parameters/renumerateEnabled"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
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
