// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/base"
	"reflect"
	"testing"
)

func TestReadMetaProperty(t *testing.T) {
	type args struct {
		metaPropertyString string
	}
	tests := []struct {
		name    string
		args    args
		want    properties.MetaProperty
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve with empty string", args{""}, nil, true},
		{"+ve", args{"id:S|Data"}, base.NewMetaProperty(baseIDs.NewID("id"), baseData.NewStringData("Data")), false},
		{"-ve incorrectFormat", args{"idS|Data"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadMetaProperty(tt.args.metaPropertyString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadMetaProperty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadMetaProperty() got = %v, want %v", got, tt.want)
			}
		})
	}
}
