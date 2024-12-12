// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mapper

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/assets/record"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Mapper
	}{
		// TODO: it should pass, but possibly due to a bug in the code, it fails
		{"+ve", baseHelpers.NewMapper(record.Prototype)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
