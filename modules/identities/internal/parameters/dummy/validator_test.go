// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"testing"
)

func Test_validator(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{"+ve with nil", args{Parameter}, false},
		{"-ve wrong parameter Type", args{baseTypes.NewParameter(baseIDs.NewStringID("newID"), baseData.NewDecData(sdkTypes.NewDec(-1)), validator)}, true},
		{"-ve wrong parameter Type", args{baseTypes.NewParameter(baseIDs.NewID("newID"), baseData.NewStringData("newStringData"), validator)}, true},
		{"+ve empty string", args{baseIDs.NewStringID("")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validator(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("validator() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
