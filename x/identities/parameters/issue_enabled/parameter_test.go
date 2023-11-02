// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue_enabled

import (
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseParameters "github.com/AssetMantle/schema/go/parameters/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
)

func Test_validator(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantError bool
	}{
		{"-ve incorrectFormat", args{baseIDs.NewStringID("")}, true},
		{"+ve", args{Parameter}, false},
		{"-ve InvalidParameter", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID(""), baseData.NewStringData("")))}, true},
		{"+ve with booleanData", args{baseData.NewBooleanData(false)}, false},
		{"-ve with different type of Data", args{baseData.NewStringData("stringData")}, true},
		{"+ve with true booleanData", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("issueEnabled"), baseData.NewBooleanData(true)))}, false},
		{"+ve with false booleanData", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("issueEnabled"), baseData.NewBooleanData(false)))}, false},
		{"+ve with incorrect ID", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewBooleanData(false)))}, true},
		{"-ve nil", args{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validator(tt.args.i); (err != nil) != tt.wantError {
				t.Errorf("validator() error = %v, wantErr %v", err, tt.wantError)
			}
		})
	}
}
