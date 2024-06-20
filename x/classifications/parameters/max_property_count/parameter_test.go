// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package max_property_count

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseParameters "github.com/AssetMantle/schema/parameters/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
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
		{"+ve with zero NumberData", args{baseData.NewNumberData(sdkTypes.ZeroInt())}, true},
		{"+ve with positive NumberData", args{baseData.NewNumberData(sdkTypes.OneInt())}, false},
		{"+ve with negative NumberData", args{baseData.NewNumberData(sdkTypes.NewInt(-1))}, true},
		{"-ve with different type of Data", args{baseData.NewStringData("stringData")}, true},
		{"+ve with positive NumberDataParam", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("maxPropertyCount"), baseData.NewNumberData(sdkTypes.OneInt())))}, false},
		{"+ve with negative NumberDataParam", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("maxPropertyCount"), baseData.NewNumberData(sdkTypes.NewInt(-1))))}, true},
		{"+ve with zero NumberDataParam", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("maxPropertyCount"), baseData.NewNumberData(sdkTypes.ZeroInt())))}, true},
		{"+ve with incorrect ID", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewNumberData(sdkTypes.ZeroInt())))}, true},
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
