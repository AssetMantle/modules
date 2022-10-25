// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"testing"

	base3 "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
	properties2 "github.com/AssetMantle/modules/schema/properties"
	base2 "github.com/AssetMantle/modules/schema/properties/base"
)

func TestReadMetaPropertyList(t *testing.T) {
	type args struct {
		metaPropertiesString string
	}
	tests := []struct {
		name    string
		args    args
		want    lists.MetaPropertyList
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve with empty string", args{""}, base.NewMetaPropertyList([]properties2.MetaProperty{}...), false},
		{"+ve", args{"ID:S|Data,ID1:S|Data1,ID2:S|Data2"}, base.NewMetaPropertyList([]properties2.MetaProperty{base2.NewMetaProperty(base3.NewStringID("ID"), base2.NewStringData("Data")), base2.NewMetaProperty(base3.NewStringID("ID1"), base2.NewStringData("Data1")), base2.NewMetaProperty(base3.NewStringID("ID2"), base2.NewStringData("Data2"))}...), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadMetaPropertyList(tt.args.metaPropertiesString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadMetaPropertyList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i := range got.GetList() {
				if got.GetList()[i].Compare(tt.want.GetList()[i]) != 0 {
					t.Errorf("ReadMetaPropertyList() got = %v, want %v", got.GetList()[i], tt.want.GetList()[i])
				}
			}
		})
	}
}
