// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maxOrderLife

import (
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/types/base"
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
		{"+ve with zero heightData", args{baseData.NewHeightData(base.NewHeight(0))}, nil},
		{"+ve with positive heightData", args{baseData.NewHeightData(base.NewHeight(1))}, nil},
		{"+ve with negative heightData", args{baseData.NewHeightData(base.NewHeight(-1))}, errorConstants.IncorrectFormat},
		{"-ve with different type of Data", args{baseData.NewStringData("stringData")}, errorConstants.IncorrectFormat},
		{"+ve with zero heightDataParam", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("maxOrderLife"), baseData.NewHeightData(base.NewHeight(0))))}, nil},
		{"+ve with positive heightDataParam", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("maxOrderLife"), baseData.NewHeightData(base.NewHeight(1))))}, nil},
		{"+ve with negative heightDataParam", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("maxOrderLife"), baseData.NewHeightData(base.NewHeight(-1))))}, errorConstants.IncorrectFormat},
		{"+ve with incorrect ID", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewHeightData(base.NewHeight(0))))}, errorConstants.IncorrectFormat},
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
