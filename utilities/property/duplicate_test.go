// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package property

import (
	"testing"

	baseData "github.com/AssetMantle/modules/schema/data/base"
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
		{"Positive Case, Unique Properties", args{propertyList: []types.Property{base.NewProperty(base.NewID("a"), baseData.NewStringData("factA")),
			base.NewProperty(base.NewID("b"), baseData.NewStringData("factB")),
			base.NewProperty(base.NewID("c"), baseData.NewStringData("factC")),
			base.NewProperty(base.NewID("d"), baseData.NewStringData("factD"))}}, false},
		{"Negative Case, DuplicateExists", args{propertyList: []types.Property{base.NewProperty(base.NewID("a"), baseData.NewStringData("factA")),
			base.NewProperty(base.NewID("b"), baseData.NewStringData("factB")),
			base.NewProperty(base.NewID("c"), baseData.NewStringData("factC")),
			base.NewProperty(base.NewID("a"), baseData.NewStringData("factD"))}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Duplicate(tt.args.propertyList); got != tt.want {
				t.Errorf("Duplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
