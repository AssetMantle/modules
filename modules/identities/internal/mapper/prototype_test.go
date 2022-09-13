// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mapper

import (
	"github.com/AssetMantle/modules/schema/helpers"
	"reflect"
	"testing"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Mapper
	}{
		// TODO: Add test cases.
		// TODO: Pointer based unit tests required
		//{"+ve", baseHelpers.NewMapper(key.Prototype, mappable.Prototype)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
