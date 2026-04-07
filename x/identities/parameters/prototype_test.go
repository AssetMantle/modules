// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/identities/parameters/issue_enabled"
	"github.com/AssetMantle/modules/x/identities/parameters/max_provision_address_count"
	"github.com/AssetMantle/modules/x/identities/parameters/quash_enabled"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.ParameterManager
	}{

		{"+ve", baseHelpers.NewParameterManager(issue_enabled.ValidatableParameter, max_provision_address_count.ValidatableParameter, quash_enabled.ValidatableParameter)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
