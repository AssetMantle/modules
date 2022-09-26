// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/errors/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
	"testing"
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

		{"-ve incorrectFormat", args{baseIDs.NewStringID("")}, constants.IncorrectFormat},
		{"+ve", args{Parameter}, nil},
		{"-ve InvalidParameter", args{baseTypes.NewParameter(baseIDs.NewStringID(""), baseData.NewStringData(""), validator)}, constants.InvalidParameter},
		{"-ve nil", args{}, constants.IncorrectFormat},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validator(tt.args.i); err != tt.wantError {
				t.Errorf("validator() error = %v, wantErr %v", err, tt.wantError)
			}
		})
	}
}
