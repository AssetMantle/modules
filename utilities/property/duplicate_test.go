// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package property

import (
	"testing"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
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
		{"Positive Case, Unique PropertyList", args{propertyList: []types.Property{baseTypes.NewProperty(baseIDs.NewID("a"), baseData.NewStringData("factA")),
			baseTypes.NewProperty(baseIDs.NewID("b"), baseData.NewStringData("factB")),
			baseTypes.NewProperty(baseIDs.NewID("c"), baseData.NewStringData("factC")),
			baseTypes.NewProperty(baseIDs.NewID("d"), baseData.NewStringData("factD"))}}, false},
		{"Negative Case, DuplicateExists", args{propertyList: []types.Property{baseTypes.NewProperty(baseIDs.NewID("a"), baseData.NewStringData("factA")),
			baseTypes.NewProperty(baseIDs.NewID("b"), baseData.NewStringData("factB")),
			baseTypes.NewProperty(baseIDs.NewID("c"), baseData.NewStringData("factC")),
			baseTypes.NewProperty(baseIDs.NewID("a"), baseData.NewStringData("factD"))}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Duplicate(tt.args.propertyList); got != tt.want {
				t.Errorf("Duplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
