// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/modules/classifications/internal/module"
	"github.com/AssetMantle/modules/modules/classifications/internal/parameters/bondRate"
	"github.com/AssetMantle/modules/modules/classifications/internal/parameters/maxPropertyCount"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name      string
		want      helpers.ParameterManager
		wantError error
	}{
		{"+ve", baseHelpers.NewParameterManager(module.Name, bondRate.ValidatableParameter, maxPropertyCount.ValidatableParameter), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
