// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package property

import (
	"testing"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

func TestDuplicate(t *testing.T) {
	type args struct {
		propertyList []properties.Property
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Positive Case, Unique PropertyList", args{propertyList: []properties.Property{baseProperties.NewMesaProperty(baseIDs.NewStringID("a"), baseData.NewStringData("factA")),
			baseProperties.NewMesaProperty(baseIDs.NewStringID("b"), baseData.NewStringData("factB")),
			baseProperties.NewMesaProperty(baseIDs.NewStringID("c"), baseData.NewStringData("factC")),
			baseProperties.NewMesaProperty(baseIDs.NewStringID("d"), baseData.NewStringData("factD"))}}, false},
		{"Negative Case, DuplicateExists", args{propertyList: []properties.Property{baseProperties.NewMesaProperty(baseIDs.NewStringID("a"), baseData.NewStringData("factA")),
			baseProperties.NewMesaProperty(baseIDs.NewStringID("b"), baseData.NewStringData("factB")),
			baseProperties.NewMesaProperty(baseIDs.NewStringID("c"), baseData.NewStringData("factC")),
			baseProperties.NewMesaProperty(baseIDs.NewStringID("a"), baseData.NewStringData("factD"))}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Duplicate(tt.args.propertyList); got != tt.want {
				t.Errorf("Duplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
