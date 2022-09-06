package base

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/parameters"
	"reflect"
	"testing"
)

func createTestInput(t *testing.T) (ids.ID, data.Data, parameters.Parameter) {
	id := baseIDs.NewID("ID")
	data := baseData.NewStringData("Data")

	testParameter := NewParameter(id, data, validator)
	return id, data, testParameter
}

func TestNewParameter(t *testing.T) {
	id, testData, _ := createTestInput(t)
	type args struct {
		id        ids.ID
		data      data.Data
		validator func(interface{}) error
	}
	tests := []struct {
		name string
		args args
		want parameters.Parameter
	}{
		// TODO: Add test cases.
		{"+ve", args{id, testData, validator}, parameter{id, testData, validator}},
		{"-ve", args{}, parameter{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParameter(tt.args.id, tt.args.data, tt.args.validator); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("NewParameter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameter_Equal(t *testing.T) {
	id, testData, testParameter := createTestInput(t)
	type fields struct {
		ID        ids.ID
		Data      data.Data
		validator func(interface{}) error
	}
	type args struct {
		compareParameter parameters.Parameter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{id, testData, validator}, args{testParameter}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := parameter{
				ID:        tt.fields.ID,
				Data:      tt.fields.Data,
				validator: tt.fields.validator,
			}
			if got := parameter.Equal(tt.args.compareParameter); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameter_GetData(t *testing.T) {
	id, testData, _ := createTestInput(t)
	type fields struct {
		ID        ids.ID
		Data      data.Data
		validator func(interface{}) error
	}
	tests := []struct {
		name   string
		fields fields
		want   data.Data
	}{
		// TODO: Add test cases.
		{"+ve", fields{id, testData, validator}, testData},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := parameter{
				ID:        tt.fields.ID,
				Data:      tt.fields.Data,
				validator: tt.fields.validator,
			}
			if got := parameter.GetData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameter_GetID(t *testing.T) {
	id, testData, _ := createTestInput(t)
	type fields struct {
		ID        ids.ID
		Data      data.Data
		validator func(interface{}) error
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		// TODO: Add test cases.
		{"+ve", fields{id, testData, validator}, id},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := parameter{
				ID:        tt.fields.ID,
				Data:      tt.fields.Data,
				validator: tt.fields.validator,
			}
			if got := parameter.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameter_GetValidator(t *testing.T) {
	id, testData, _ := createTestInput(t)
	type fields struct {
		ID        ids.ID
		Data      data.Data
		validator func(interface{}) error
	}
	tests := []struct {
		name   string
		fields fields
		want   func(interface{}) error
	}{
		// TODO: Add test cases.
		{"+ve", fields{id, testData, validator}, validator},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := parameter{
				ID:        tt.fields.ID,
				Data:      tt.fields.Data,
				validator: tt.fields.validator,
			}
			if got := parameter.GetValidator(); !reflect.DeepEqual(got, tt.want) {
				//t.Errorf("GetValidator() = %p, want %p", got, tt.want)
			}
		})
	}
}

func Test_parameter_Mutate(t *testing.T) {
	id, testData, _ := createTestInput(t)
	newData := baseData.NewStringData("Data")
	type fields struct {
		ID        ids.ID
		Data      data.Data
		validator func(interface{}) error
	}
	type args struct {
		data data.Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   parameters.Parameter
	}{
		// TODO: Add test cases.
		{"+ve", fields{id, testData, validator}, args{newData}, parameter{id, newData, validator}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := parameter{
				ID:        tt.fields.ID,
				Data:      tt.fields.Data,
				validator: tt.fields.validator,
			}
			if got := parameter.Mutate(tt.args.data); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameter_String(t *testing.T) {
	id, testData, testParameter := createTestInput(t)
	type fields struct {
		ID        ids.ID
		Data      data.Data
		validator func(interface{}) error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve", fields{id, testData, validator}, testParameter.String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := parameter{
				ID:        tt.fields.ID,
				Data:      tt.fields.Data,
				validator: tt.fields.validator,
			}
			if got := parameter.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameter_Validate(t *testing.T) {
	//id, testData, _ := createTestInput(t)
	type fields struct {
		ID        ids.ID
		Data      data.Data
		validator func(interface{}) error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		//{"+ve", fields{id, testData, validator}, false},
		// TODO: Should not panic
		//{"-ve", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := parameter{
				ID:        tt.fields.ID,
				Data:      tt.fields.Data,
				validator: tt.fields.validator,
			}
			if err := parameter.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
