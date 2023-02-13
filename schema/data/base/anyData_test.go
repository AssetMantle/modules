package base

import (
	"reflect"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

var (
	fromAddress1       = "cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"
	fromAccAddress, _  = sdkTypes.AccAddressFromBech32(fromAddress)
	fromAccAddress1, _ = sdkTypes.AccAddressFromBech32(fromAddress1)
	dataList           = []data.Data{NewAccAddressData(fromAccAddress), NewAccAddressData(fromAccAddress1)}
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
		{"String Data", args{"S|newFact"}, NewStringData("newFact"), false},
		{"-ve Unknown Data", args{"SomeRandomData"}, nil, true},
		{"List Data", args{"L|A|cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c,A|cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"}, NewListData(NewListData(dataList...)), false},
		{"List Data empty list", args{"L|"}, NewListData(NewListData()), false},
		{"Id Data", args{"I|data"}, NewIDData(baseIDs.NewStringID("data")), false},
		{"Height Data", args{"H|100"}, NewHeightData(baseTypes.NewHeight(100)), false},
		{"Dec Data", args{"D|100"}, NewDecData(sdkTypes.NewDec(100)), false},
		{"Bool Data", args{"B|true"}, NewBooleanData(true), false},
		{"AccAddress data", args{"A|cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"}, NewAccAddressData(fromAccAddress), false},
		{"-ve String Data", args{"S|S,|newFact"}, NewStringData("S,|newFact"), true},
		{"-ve List Data String", args{"L|S|,TestData,S|,Test"}, NewListData(NewListData([]data.Data{NewStringData("S|,TestData"), NewStringData("S|,Test")}...)), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeAnyData().FromString(tt.args.dataString)
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
		{"-ve", args{"F", "SFw"}, "F|SFw"},
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
	fromAccAddress, nil := sdkTypes.AccAddressFromBech32("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	type args struct {
		dataString string
	}
	tests := []struct {
		name    string
		args    args
		want    data.AccAddressData
		wantErr bool
	}{
		{"+ve nil", args{}, PrototypeAccAddressData().(data.AccAddressData), false},
		{"+ve string", args{"cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"}, NewAccAddressData(fromAccAddress), false},
		{"-ve", args{"testData"}, PrototypeAccAddressData().(data.AccAddressData), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeAccAddressData().FromString(tt.args.dataString)
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
		{"+ve nil", args{}, PrototypeBooleanData(), false},
		{"+ve string", args{"true"}, NewBooleanData(true), false},
		{"+ve string", args{"false"}, NewBooleanData(false), false},
		{"-ve", args{"testData"}, PrototypeBooleanData(), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeBooleanData().FromString(tt.args.dataString)
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
		{"+ve nil", args{}, PrototypeDecData(), false},
		{"+ve string", args{"100"}, NewDecData(sdkTypes.NewDec(100)), false},
		{"+ve with nil", args{"-100"}, NewDecData(sdkTypes.NewDec(-100)), false},
		{"-ve", args{"testData"}, PrototypeDecData(), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeDecData().FromString(tt.args.dataString)
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
		{"+ve nil", args{}, PrototypeHeightData(), false},
		{"+ve string", args{"100"}, NewHeightData(baseTypes.NewHeight(100)), false},
		{"+ve with nil", args{"-100"}, NewHeightData(baseTypes.NewHeight(-100)), false},
		{"-ve", args{"testData"}, PrototypeHeightData(), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeHeightData().FromString(tt.args.dataString)
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
		{"+ve nil", args{}, PrototypeIDData(), false},
		{"+ve", args{"L|A|cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c,A|cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef"}, NewIDData(baseIDs.NewStringID("L|A|cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c,A|cosmos1x53dugvr4xvew442l9v2r5x7j8gfvged2zk5ef")), false},
		{"-ve string with special char", args{"testDataString|,"}, NewIDData(baseIDs.NewStringID("testDataString|,")), false},
		{"-ve", args{"testData"}, NewIDData(baseIDs.NewStringID("testData")), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeIDData().FromString(tt.args.dataString)
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
		{"+ve nil", args{}, PrototypeListData(), false},
		{"+ve string", args{"S|1,S|2,S|3"}, NewListData(NewListData([]data.Data{NewStringData("1"), NewStringData("2"), NewStringData("3")}...)), false},
		{"-ve", args{"testData"}, PrototypeListData(), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeListData().FromString(tt.args.dataString)
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
		{"+ve nil", args{}, PrototypeStringData(), false},
		{"+ve string", args{"testDataString"}, NewStringData("testDataString"), false},
		{"-ve", args{"testData"}, NewStringData("testData"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrototypeStringData().FromString(tt.args.dataString)
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
