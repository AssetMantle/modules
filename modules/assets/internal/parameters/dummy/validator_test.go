// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

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
		{"-ve InvalidParameter", args{baseTypes.NewParameter(baseIDs.NewStringID(""), baseData.NewStringData(""), validator)}, errorConstants.InvalidParameter},
		{"-ve nil", args{}, errorConstants.IncorrectFormat},
		{"+ve with decData", args{baseData.NewDecData(sdkTypes.NewDec(-1))}, errorConstants.InvalidParameter},
		{"-ve with different type of Data", args{baseData.NewStringData("stringData")}, errorConstants.IncorrectFormat},
		{"-ve InvalidParameter", args{baseTypes.NewParameter(baseIDs.NewStringID(""), baseData.NewStringData(""), validator)}, errorConstants.InvalidParameter},
		{"-ve with -ve decData", args{baseTypes.NewParameter(baseIDs.NewStringID("ID"), baseData.NewDecData(sdkTypes.NewDec(-1)), validator)}, errorConstants.InvalidParameter},
		{"+ve with +ve decData", args{baseTypes.NewParameter(baseIDs.NewStringID("ID"), baseData.NewDecData(sdkTypes.NewDec(1)), validator)}, nil},
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
