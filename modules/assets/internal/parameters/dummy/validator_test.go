// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/errors"
	baseData "github.com/AssetMantle/modules/schema/data/base"
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

		{"-ve incorrectFormat", args{baseIDs.NewID("")}, errors.IncorrectFormat},
		{"+ve", args{Parameter}, nil},
		{"+ve with decData", args{baseData.NewDecData(sdkTypes.NewDec(-1))}, errors.InvalidParameter},
		{"-ve with different type of Data", args{baseData.NewStringData("stringData")}, errors.IncorrectFormat},
		{"-ve InvalidParameter", args{baseTypes.NewParameter(baseIDs.NewID(""), baseData.NewStringData(""), validator)}, errors.InvalidParameter},
		{"-ve with -ve decData", args{baseTypes.NewParameter(baseIDs.NewID("ID"), baseData.NewDecData(sdkTypes.NewDec(-1)), validator)}, errors.InvalidParameter},
		{"+ve with +ve decData", args{baseTypes.NewParameter(baseIDs.NewID("ID"), baseData.NewDecData(sdkTypes.NewDec(1)), validator)}, nil},
		{"-ve nil", args{}, errors.IncorrectFormat},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validator(tt.args.i); err != tt.wantError {
				t.Errorf("validator() error = %v, wantErr %v", err, tt.wantError)
			}
		})
	}
}
