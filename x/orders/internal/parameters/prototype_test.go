// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"

	"github.com/AssetMantle/modules/x/orders/internal/module"
	"github.com/AssetMantle/modules/x/orders/internal/parameters/maxOrderLife"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.ParameterManager
	}{
		{"+ve", baseHelpers.NewParameterManager(module.Name, maxOrderLife.ValidatableParameter)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
