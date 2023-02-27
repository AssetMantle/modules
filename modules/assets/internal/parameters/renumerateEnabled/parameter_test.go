// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerateEnabled

import (
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"testing"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
)

func Test_validator(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantError error
	}{
		{"-ve incorrectFormat", args{baseIDs.NewStringID("")}, errorConstants.IncorrectFormat},
		{"+ve", args{Parameter}, nil},
		{"-ve InvalidParameter", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID(""), baseData.NewStringData("")))}, errorConstants.IncorrectFormat},
		{"+ve with booleanData", args{baseData.NewBooleanData(false)}, nil},
		{"-ve with different type of Data", args{baseData.NewStringData("stringData")}, errorConstants.IncorrectFormat},
		{"-ve InvalidParameter", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID(""), baseData.NewStringData("")))}, errorConstants.IncorrectFormat},
		{"+ve with true booleanData", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("renumerateEnabled"), baseData.NewBooleanData(true)))}, nil},
		{"+ve with false booleanData", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("renumerateEnabled"), baseData.NewBooleanData(false)))}, nil},
		{"+ve with incorrect ID", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewBooleanData(false)))}, errorConstants.IncorrectFormat},
		{"-ve nil", args{}, errorConstants.IncorrectFormat},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validator(tt.args.i); err != tt.wantError {
				t.Errorf("validator() error = %v, wantErr %v", err, tt.wantError)
			}
		})
	}
}
