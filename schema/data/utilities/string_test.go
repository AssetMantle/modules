// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"reflect"
	"testing"

	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"

	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/data/base"
)

var (
	fromAddress        = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAddress1       = "cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"
	fromAccAddress, _  = types.AccAddressFromBech32(fromAddress)
	fromAccAddress1, _ = types.AccAddressFromBech32(fromAddress1)
	dataList           = []data.Data{base.NewAccAddressData(fromAccAddress), base.NewAccAddressData(fromAccAddress1)}
)

func TestReadData(t *testing.T) {
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
		{"List Data", args{"L|A|cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c,A|cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"}, base.NewListData(baseLists.NewDataList(dataList...)), false}, // TODO: stack overflow
		{"List Data empty list", args{"L|"}, base.NewListData(baseLists.NewDataList()), false},                                                                                                // TODO: stack overflow
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
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadData() got = %v, want %v", got, tt.want)
				t.Errorf("ReadData() got = %T, want %T", got, tt.want)
			}
		})
	}
}

func Test_joinDataTypeAndValueStrings(t *testing.T) {
	type args struct {
		dataType  string
		dataValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"+ve string", args{"S", "Data"}, "S|Data"},
		{"+ve Id Data", args{"I", "Data"}, "I|Data"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := joinDataTypeAndValueStrings(tt.args.dataType, tt.args.dataValue); got != tt.want {
				t.Errorf("joinDataTypeAndValueStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readAccAddressData(t *testing.T) {
	fromAccAddress, nil := types.AccAddressFromBech32("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	type args struct {
		dataString string
	}
	tests := []struct {
		name    string
		args    args
		want    data.AccAddressData
		wantErr bool
	}{
		{"+ve nil", args{}, base.AccAddressDataPrototype(), false},
		{"+ve string", args{"cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"}, base.NewAccAddressData(fromAccAddress), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readAccAddressData(tt.args.dataString)
			if (err != nil) != tt.wantErr {
				t.Errorf("readAccAddressData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readAccAddressData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readBooleanData(t *testing.T) {
	type args struct {
		dataString string
	}
	tests := []struct {
		name    string
		args    args
		want    data.BooleanData
		wantErr bool
	}{
		{"+ve nil", args{}, base.BooleanDataPrototype(), false},
		{"+ve string", args{"true"}, base.NewBooleanData(true), false},
		{"+ve string", args{"false"}, base.NewBooleanData(false), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readBooleanData(tt.args.dataString)
			if (err != nil) != tt.wantErr {
				t.Errorf("readBooleanData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readBooleanData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readDecData(t *testing.T) {
	type args struct {
		dataString string
	}
	tests := []struct {
		name    string
		args    args
		want    data.DecData
		wantErr bool
	}{
		{"+ve nil", args{}, base.DecDataPrototype(), false},
		{"+ve string", args{"100"}, base.NewDecData(types.NewDec(100)), false},
		{"+ve with nil", args{"-100"}, base.NewDecData(types.NewDec(-100)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readDecData(tt.args.dataString)
			if (err != nil) != tt.wantErr {
				t.Errorf("readDecData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readDecData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readHeightData(t *testing.T) {
	type args struct {
		dataString string
	}
	tests := []struct {
		name    string
		args    args
		want    data.HeightData
		wantErr bool
	}{
		{"+ve nil", args{}, base.HeightDataPrototype(), false},
		{"+ve string", args{"100"}, base.NewHeightData(baseTypes.NewHeight(100)), false},
		{"+ve with nil", args{"-100"}, base.NewHeightData(baseTypes.NewHeight(-100)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readHeightData(tt.args.dataString)
			if (err != nil) != tt.wantErr {
				t.Errorf("readHeightData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readHeightData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readIDData(t *testing.T) {
	type args struct {
		dataString string
	}
	tests := []struct {
		name    string
		args    args
		want    data.IDData
		wantErr bool
	}{
		{"+ve nil", args{}, base.IDDataPrototype(), false},
		{"+", args{"L|A|cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c,A|cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"}, base.NewIDData(baseIDs.NewStringID("L|A|cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c,A|cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef")), false},
		{"-ve string with special char", args{"testDataString|,"}, base.NewIDData(baseIDs.NewStringID("testDataString|,")), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readIDData(tt.args.dataString)
			if (err != nil) != tt.wantErr {
				t.Errorf("readIDData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readIDData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readListData(t *testing.T) {
	type args struct {
		dataString string
	}
	tests := []struct {
		name    string
		args    args
		want    data.ListData
		wantErr bool
	}{
		{"+ve nil", args{}, base.ListDataPrototype(), false},
		{"+ve string", args{"S|1,S|2,S|3"}, base.NewListData(baseLists.NewDataList([]data.Data{base.NewStringData("1"), base.NewStringData("2"), base.NewStringData("3")}...)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readListData(tt.args.dataString)
			if (err != nil) != tt.wantErr {
				t.Errorf("readListData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readListData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readStringData(t *testing.T) {
	type args struct {
		dataString string
	}
	tests := []struct {
		name    string
		args    args
		want    data.StringData
		wantErr bool
	}{
		{"+ve nil", args{}, base.StringDataPrototype(), false},
		{"+ve string", args{"testDataString"}, base.NewStringData("testDataString"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readStringData(tt.args.dataString)
			if (err != nil) != tt.wantErr {
				t.Errorf("readStringData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readStringData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitDataTypeAndValueStrings(t *testing.T) {
	type args struct {
		dataTypeAndValueString string
	}
	tests := []struct {
		name          string
		args          args
		wantDataType  string
		wantDataValue string
	}{
		{"+ve String", args{"S|data"}, "S", "data"},
		{"+ve Bool", args{"B|true"}, "B", "true"},
		{"+ve Int", args{"I|100"}, "I", "100"},
		{"+ve Dec", args{"D|100.00"}, "D", "100.00"},
		{"+ve Height", args{"H|100"}, "H", "100"},
		{"+ve ID", args{"ID|100"}, "ID", "100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDataType, gotDataValue := splitDataTypeAndValueStrings(tt.args.dataTypeAndValueString)
			if gotDataType != tt.wantDataType {
				t.Errorf("splitDataTypeAndValueStrings() gotDataType = %v, want %v", gotDataType, tt.wantDataType)
			}
			if gotDataValue != tt.wantDataValue {
				t.Errorf("splitDataTypeAndValueStrings() gotDataValue = %v, want %v", gotDataValue, tt.wantDataValue)
			}
		})
	}
}
