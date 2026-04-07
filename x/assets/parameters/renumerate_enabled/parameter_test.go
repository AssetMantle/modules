// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate_enabled

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
		{"-ve incorrectFormat", args{baseIDs.NewStringID("")}, true},
		{"+ve", args{Parameter}, false},
		{"-ve InvalidParameter", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID(""), baseData.NewStringData("")))}, true},
		{"+ve with booleanData", args{baseData.NewBooleanData(false)}, true},
		{"-ve with different type of Data", args{baseData.NewStringData("stringData")}, true},
		{"+ve with true booleanData", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("renumerateEnabled"), baseData.NewBooleanData(true)))}, false},
		{"+ve with false booleanData", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("renumerateEnabled"), baseData.NewBooleanData(false)))}, false},
		{"+ve with incorrect ID", args{baseParameters.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewBooleanData(false)))}, true},
		{"-ve nil", args{}, true},
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
