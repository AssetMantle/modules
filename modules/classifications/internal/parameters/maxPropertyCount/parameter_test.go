// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maxPropertyCount

import (
	"testing"

	baseProperties "github.com/AssetMantle/modules/schema/properties/base"

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
		{"+ve with zero NumberData", args{baseData.NewNumberData(0)}, nil},
		{"+ve with positive NumberData", args{baseData.NewNumberData(1)}, nil},
		{"+ve with negative NumberData", args{baseData.NewNumberData(-1)}, nil},
		{"-ve with different type of Data", args{baseData.NewStringData("stringData")}, errorConstants.IncorrectFormat},
		{"+ve with positive NumberDataParam", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("maxPropertyCount"), baseData.NewNumberData(1)))}, nil},
		{"+ve with negative NumberDataParam", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("maxPropertyCount"), baseData.NewNumberData(-1)))}, nil},
		{"+ve with zero NumberDataParam", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("maxPropertyCount"), baseData.NewNumberData(0)))}, nil},
		{"+ve with incorrect ID", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewNumberData(0)))}, errorConstants.IncorrectFormat},
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
