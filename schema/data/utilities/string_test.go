// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func TestReadData(t *testing.T) {
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAddress1 := "cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"
	fromAccAddress, _ := types.AccAddressFromBech32(fromAddress)
	fromAccAddress1, _ := types.AccAddressFromBech32(fromAddress1)
	dataList := make([]data.Data, 2)
	dataList[0] = base.NewAccAddressData(fromAccAddress)
	dataList[1] = base.NewAccAddressData(fromAccAddress1)
	type args struct {
		dataString string
	}
	tests := []struct {
		name    string
		args    args
		want    data.Data
		wantErr bool
	}{
		// TODO:  fix after issue #50 fix, https://github.com/AssetMantle/modules/issues/50
		{"String Data", args{"S|newFact"}, base.NewStringData("newFact"), false},
		{"Unknown Data", args{"SomeRandomData"}, nil, true},
		{"List Data", args{"L|cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c,cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"}, base.NewListData(baseLists.NewDataList(dataList...)), false}, // TODO: stack overflow
		{"List Data", args{"L|"}, base.NewListData(nil), false}, // TODO: stack overflow
		{"Id Data", args{"I|data"}, base.NewIDData(baseIDs.NewStringID("data")), false},
		{"Height Data", args{"H|100"}, base.NewHeightData(baseTypes.NewHeight(100)), false},
		{"Dec Data", args{"D|100"}, base.NewDecData(types.NewDec(100)), false},
		{"Bool Data", args{"B|true"}, base.NewBooleanData(true), false},
		{"AccAddress data", args{"A|cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"}, base.NewAccAddressData(fromAccAddress), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadData(tt.args.dataString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadData() got = %v, want %v", got, tt.want)
				t.Errorf("ReadData() got = %T, want %T", got, tt.want)
			}
		})
	}
}
