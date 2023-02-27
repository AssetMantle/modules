// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wrapAllowedCoins

import (
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/types/base"
	"testing"

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
		{"-ve InvalidParameter", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID(""), baseData.NewStringData("")))}, errorConstants.IncorrectFormat},
		{"-ve with different type of Data", args{baseData.NewStringData("stringData")}, errorConstants.IncorrectFormat},
		{"+ve valid listData", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("wrapAllowedCoins"), baseData.NewListData(baseData.NewIDData(baseIDs.NewCoinID(baseIDs.NewStringID("stake"))))))}, nil},
		{"+ve with incorrect ID", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewHeightData(base.NewHeight(0))))}, errorConstants.IncorrectFormat},
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
