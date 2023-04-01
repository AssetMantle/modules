// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maxOrderLife

import (
	"testing"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseParameters "github.com/AssetMantle/modules/schema/parameters/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/types/base"
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
		{"+ve with zero heightData", args{baseData.NewHeightData(base.NewHeight(0))}, false},
		{"+ve with positive heightData", args{baseData.NewHeightData(base.NewHeight(1))}, false},
		{"+ve with negative heightData", args{baseData.NewHeightData(base.NewHeight(-1))}, true},
		{"-ve with different type of Data", args{baseData.NewStringData("stringData")}, true},
		{"+ve with zero heightDataParam", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("maxOrderLife"), baseData.NewHeightData(base.NewHeight(0))))}, false},
		{"+ve with positive heightDataParam", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("maxOrderLife"), baseData.NewHeightData(base.NewHeight(1))))}, false},
		{"+ve with negative heightDataParam", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("maxOrderLife"), baseData.NewHeightData(base.NewHeight(-1))))}, true},
		{"+ve with incorrect ID", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewHeightData(base.NewHeight(0))))}, true},
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
