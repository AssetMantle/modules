package base

import (
	"fmt"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/go/data"
	base4 "github.com/AssetMantle/schema/go/data/base"
	base3 "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/parameters"
	"github.com/AssetMantle/schema/go/parameters/base"
	base2 "github.com/AssetMantle/schema/go/properties/base"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewValidatableParameter(t *testing.T) {
	tests := []struct {
		name         string
		param        parameters.Parameter
		validateFunc func(i interface{}) error
		wantErr      error
	}{
		{
			name:         "Test invalid input",
			param:        base.NewParameter(base2.NewMetaProperty(base3.NewStringID("testName"), base4.NewStringData("testData"))),
			validateFunc: func(i interface{}) error { return fmt.Errorf("error") },
			wantErr:      fmt.Errorf("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vp := NewValidatableParameter(tt.param, tt.validateFunc)
			if err := vp.Validate(); (err != nil) != (tt.wantErr != nil) {
				t.Errorf("NewValidatableParameter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validatableParameter_GetParameter(t *testing.T) {
	type fields struct {
		parameter parameters.Parameter
		validator func(i interface{}) error
	}
	tests := []struct {
		name   string
		fields fields
		want   parameters.Parameter
	}{
		{
			name: "Test valid input",
			fields: fields{
				parameter: base.NewParameter(base2.NewMetaProperty(base3.NewStringID("testName"), base4.NewStringData("testData"))),
				validator: func(i interface{}) error { return nil },
			},
			want: base.NewParameter(base2.NewMetaProperty(base3.NewStringID("testName"), base4.NewStringData("testData"))),
		},
		{
			name: "Test nil input",
			fields: fields{
				parameter: nil,
				validator: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validatableParameter := validatableParameter{
				parameter: tt.fields.parameter,
				validator: tt.fields.validator,
			}
			assert.Equalf(t, tt.want, validatableParameter.GetParameter(), "GetParameter()")
		})
	}
}

func Test_validatableParameter_GetValidator(t *testing.T) {
	type fields struct {
		parameter parameters.Parameter
		validator func(i interface{}) error
	}
	tests := []struct {
		name   string
		fields fields
		want   func(i interface{}) error
	}{
		{
			name: "Test valid input",
			fields: fields{
				parameter: base.NewParameter(base2.NewMetaProperty(base3.NewStringID("testName"), base4.NewStringData("testData"))),
				validator: func(i interface{}) error {
					if i != base4.NewStringData("valid") {
						return fmt.Errorf("error")
					}
					return nil
				},
			},
			want: func(i interface{}) error {
				if i != base4.NewStringData("valid") {
					return fmt.Errorf("error")
				}
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validatableParameter := validatableParameter{
				parameter: tt.fields.parameter,
				validator: tt.fields.validator,
			}
			assert.Equalf(t, tt.want(base4.NewStringData("aa")), validatableParameter.GetValidator()(base4.NewStringData("aa")), "GetValidator()")
		})
	}
}

func Test_validatableParameter_Mutate(t *testing.T) {
	type fields struct {
		parameter parameters.Parameter
		validator func(i interface{}) error
	}
	type args struct {
		data data.Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.ValidatableParameter
	}{
		{
			name: "Test valid input",
			fields: fields{
				parameter: base.NewParameter(base2.NewMetaProperty(base3.NewStringID("testName"), base4.NewStringData("testData"))),
				validator: func(i interface{}) error { return nil },
			},
			args: args{
				data: base4.NewStringData("mutatedTestData"),
			},
			want: validatableParameter{
				parameter: base.NewParameter(base2.NewMetaProperty(base3.NewStringID("testName"), base4.NewStringData("mutatedTestData"))),
				validator: func(i interface{}) error { return nil },
			},
		},
		{
			name: "Test nil input",
			fields: fields{
				parameter: base.NewParameter(base2.NewMetaProperty(base3.NewStringID("testName"), base4.NewStringData("testData"))),
				validator: func(i interface{}) error { return nil },
			},
			args: args{
				data: nil,
			},
			want: validatableParameter{
				parameter: base.NewParameter(base2.NewMetaProperty(base3.NewStringID("testName"), base4.NewStringData("testData"))),
				validator: func(i interface{}) error { return nil },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validatableParameter := validatableParameter{
				parameter: tt.fields.parameter,
				validator: tt.fields.validator,
			}
			assert.Equalf(t, tt.want.GetParameter().String(), validatableParameter.Mutate(tt.args.data).GetParameter().String(), "Mutate(%v)", tt.args.data)
		})
	}
}

func Test_validatableParameter_Validate(t *testing.T) {
	type fields struct {
		parameter parameters.Parameter
		validator func(i interface{}) error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Test invalid input",
			fields: fields{
				parameter: base.NewParameter(base2.NewMetaProperty(base3.NewStringID("testName"), base4.NewStringData("testData"))),
				validator: func(i interface{}) error { return fmt.Errorf("error") },
			},
			wantErr: assert.Error,
		},
		{
			name: "Test valid input",
			fields: fields{
				parameter: base.NewParameter(base2.NewMetaProperty(base3.NewStringID("testName"), base4.NewStringData("testData"))),
				validator: func(i interface{}) error { return nil },
			},
			wantErr: assert.NoError,
		},
		{
			name: "Test nil input",
			fields: fields{
				parameter: base.NewParameter(base2.NewMetaProperty(base3.NewStringID("testName"), base4.NewStringData("testData"))),
				validator: nil,
			},
			wantErr: assert.NoError,
		},
		{
			name: "Test nil parameter and validator",
			fields: fields{
				parameter: nil,
				validator: nil,
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validatableParameter := validatableParameter{
				parameter: tt.fields.parameter,
				validator: tt.fields.validator,
			}
			tt.wantErr(t, validatableParameter.Validate(), fmt.Sprintf("Validate()"))
		})
	}
}
