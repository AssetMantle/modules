// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"reflect"
	"testing"

	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"

	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{

		{"+ve", baseHelpers.NewAuxiliaries(authenticate.Auxiliary).Get("authenticate").GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype().Get("authenticate").GetName(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
