// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package dummy

import (
	"github.com/AssetMantle/modules/constants/errors"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
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

		{"-ve incorrectFormat", args{baseIDs.NewID("")}, errors.IncorrectFormat},
		{"+ve", args{Parameter}, nil},
		//{"-ve InvalidParameter", args{baseTypes.NewParameter(baseIDs.NewID(""), baseData.NewStringData(""), validator)}, errors.InvalidParameter},
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
