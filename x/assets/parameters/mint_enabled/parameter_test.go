// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint_enabled

import (
	"fmt"
	"github.com/AssetMantle/schema/parameters"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseParameters "github.com/AssetMantle/schema/parameters/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
)

func Test_validator(t *testing.T) {
	type args struct {
		i interface{} // accepts Parameter or invalid types for negative testing
	}
	tests := []struct {
		name      string
		args      args
		wantError bool
	}{
		{"incorrect format", args{baseIDs.NewStringID("")}, true},
		{"valid", args{Parameter}, false},
		{"invalid parameter", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID(""), baseData.NewStringData("")))}, true},
		{"raw boolean data", args{baseData.NewBooleanData(false)}, true},
		{"wrong data type", args{baseData.NewStringData("stringData")}, true},
		{"true boolean parameter", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("mintEnabled"), baseData.NewBooleanData(true)))}, false},
		{"false boolean parameter", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("mintEnabled"), baseData.NewBooleanData(false)))}, false},
		{"incorrect parameter ID", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewBooleanData(false)))}, true},
		{"nil input", args{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, ok := tt.args.i.(parameters.Parameter)
			var err error
			if ok {
				err = validator(p)
			} else {
				err = fmt.Errorf("invalid parameter type %T", tt.args.i)
			}
			if (err != nil) != tt.wantError {
				t.Errorf("validator() error = %v, wantErr %v", err, tt.wantError)
			}
		})
	}
}
