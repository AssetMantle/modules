// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"github.com/AssetMantle/modules/schema/lists"
	"reflect"
	"testing"
)

func TestReadMetaProperties(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadMetaProperties(tt.args.metaPropertiesString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadMetaProperties() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadMetaProperties() got = %v, want %v", got, tt.want)
			}
		})
	}
}
