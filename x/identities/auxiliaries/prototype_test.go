// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"testing"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/stretchr/testify/assert"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{

		{"valid", baseHelpers.NewAuxiliaries(authenticate.Auxiliary).GetAuxiliary("authenticate").GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Prototype().GetAuxiliary("authenticate").GetName()
			assert.Equal(t, tt.want, got, "Prototype()")
		})
	}
}
