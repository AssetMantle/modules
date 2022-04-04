// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package property

import (
	"testing"

	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
)

func TestDuplicate(t *testing.T) {
	type args struct {
		propertyList []types.Property
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Positive Case, Unique Properties", args{propertyList: []types.Property{base.NewProperty(base.NewID("a"), base.NewStringData("factA")),
			base.NewProperty(base.NewID("b"), base.NewStringData("factB")),
			base.NewProperty(base.NewID("c"), base.NewStringData("factC")),
			base.NewProperty(base.NewID("d"), base.NewStringData("factD"))}}, false},
		{"Negative Case, DuplicateExists", args{propertyList: []types.Property{base.NewProperty(base.NewID("a"), base.NewStringData("factA")),
			base.NewProperty(base.NewID("b"), base.NewStringData("factB")),
			base.NewProperty(base.NewID("c"), base.NewStringData("factC")),
			base.NewProperty(base.NewID("a"), base.NewStringData("factD"))}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Duplicate(tt.args.propertyList); got != tt.want {
				t.Errorf("Duplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
