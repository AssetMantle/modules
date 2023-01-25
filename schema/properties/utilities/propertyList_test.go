// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"testing"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

func TestDuplicate(t *testing.T) {
	type args struct {
		propertyList []properties.AnyProperty
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Positive Case, Unique PropertyList", args{propertyList: []properties.AnyProperty{baseProperties.NewMesaProperty(baseIDs.NewStringID("a"), baseData.NewStringData("factA")).ToAnyProperty(),
			baseProperties.NewMesaProperty(baseIDs.NewStringID("b"), baseData.NewStringData("factB")).ToAnyProperty(),
			baseProperties.NewMesaProperty(baseIDs.NewStringID("c"), baseData.NewStringData("factC")).ToAnyProperty(),
			baseProperties.NewMesaProperty(baseIDs.NewStringID("d"), baseData.NewStringData("factD")).ToAnyProperty()}}, false},
		{"Negative Case, DuplicateExists", args{propertyList: []properties.AnyProperty{baseProperties.NewMesaProperty(baseIDs.NewStringID("a"), baseData.NewStringData("factA")).ToAnyProperty(),
			baseProperties.NewMesaProperty(baseIDs.NewStringID("b"), baseData.NewStringData("factB")).ToAnyProperty(),
			baseProperties.NewMesaProperty(baseIDs.NewStringID("c"), baseData.NewStringData("factC")).ToAnyProperty(),
			baseProperties.NewMesaProperty(baseIDs.NewStringID("a"), baseData.NewStringData("factD")).ToAnyProperty()}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDuplicate(tt.args.propertyList); got != tt.want {
				t.Errorf("IsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
