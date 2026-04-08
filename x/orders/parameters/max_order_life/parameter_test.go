// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package max_order_life

import (
	"fmt"
	"github.com/AssetMantle/schema/parameters"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseParameters "github.com/AssetMantle/schema/parameters/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"github.com/AssetMantle/schema/types/base"
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
		{"raw zero height data", args{baseData.NewHeightData(base.NewHeight(0))}, true},
		{"raw positive height data", args{baseData.NewHeightData(base.NewHeight(1))}, true},
		{"raw negative height data", args{baseData.NewHeightData(base.NewHeight(-1))}, true},
		{"wrong data type", args{baseData.NewStringData("stringData")}, true},
		{"zero height parameter", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("maxOrderLife"), baseData.NewHeightData(base.NewHeight(0))))}, false},
		{"valid positive height parameter", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("maxOrderLife"), baseData.NewHeightData(base.NewHeight(1))))}, false},
		{"negative height parameter", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("maxOrderLife"), baseData.NewHeightData(base.NewHeight(-1))))}, false},
		{"incorrect parameter ID", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewHeightData(base.NewHeight(0))))}, true},
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
