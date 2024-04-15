// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCLIFlag(t *testing.T) {
	type args struct {
		name  string
		value interface{}
		usage string
	}
	tests := []struct {
		name string
		args args
		want helpers.CLIFlag
	}{
		{"+ve", args{"name", "value", ",usage"}, cliFlag{"name", "value", ",usage"}},
		{"nil", args{"", nil, ""}, cliFlag{"", nil, ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewCLIFlag(tt.args.name, tt.args.value, tt.args.usage), "NewCLIFlag(%v, %v, %v)", tt.args.name, tt.args.value, tt.args.usage)
		})
	}
}

func Test_cliFlag_GetName(t *testing.T) {
	type fields struct {
		name  string
		value interface{}
		usage string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test GetName",
			fields: fields{
				name:  "test",
				value: "test",
				usage: "test",
			},
			want: "test",
		},
		{
			name: "Test GetName",
			fields: fields{
				name:  "",
				value: nil,
				usage: "",
			},
			want: "",
		},
		{
			name: "Test GetName",
			fields: fields{
				name:  "test",
				value: nil,
				usage: "",
			},
			want: "test",
		},
		{
			name: "Test GetName",
			fields: fields{
				name:  "",
				value: "test",
				usage: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cliFlag := cliFlag{
				name:  tt.fields.name,
				value: tt.fields.value,
				usage: tt.fields.usage,
			}
			assert.Equalf(t, tt.want, cliFlag.GetName(), "GetName()")
		})
	}
}

func Test_cliFlag_GetValue(t *testing.T) {
	type fields struct {
		name  string
		value interface{}
		usage string
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "Test GetValue",
			fields: fields{
				name:  "test",
				value: "test",
				usage: "test",
			},
			want: "test",
		},
		{
			name: "Test GetValue",
			fields: fields{
				name:  "",
				value: nil,
				usage: "",
			},
			want: nil,
		},
		{
			name: "Test GetValue",
			fields: fields{
				name:  "test",
				value: nil,
				usage: "",
			},
			want: nil,
		},
		{
			name: "Test GetValue",
			fields: fields{
				name:  "",
				value: "test",
				usage: "",
			},
			want: "test",
		},
		{
			name: "Test GetValue",
			fields: fields{
				name:  "",
				value: 1,
				usage: "",
			},
			want: 1,
		},
		{
			name: "Test GetValue",
			fields: fields{
				name:  "",
				value: int64(1),
				usage: "",
			},
			want: int64(1),
		},
		{
			name: "Test GetValue",
			fields: fields{
				name:  "",
				value: false,
				usage: "",
			},
			want: false,
		},
		{
			name: "Test GetValue",
			fields: fields{
				name:  "",
				value: true,
				usage: "",
			},
			want: true,
		},
		{
			name: "Test GetValue",
			fields: fields{
				name:  "",
				value: struct{}{},
				usage: "",
			},
			want: struct{}{},
		},
		{
			name: "Test GetValue",
			fields: fields{
				name:  "",
				value: nil,
				usage: "",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cliFlag := cliFlag{
				name:  tt.fields.name,
				value: tt.fields.value,
				usage: tt.fields.usage,
			}
			assert.Equalf(t, tt.want, cliFlag.GetValue(), "GetValue()")
		})
	}
}

func Test_cliFlag_ReadCLIValue(t *testing.T) {
	type fields struct {
		name  string
		value interface{}
		usage string
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
		panics bool
	}{
		{
			name: "Test ReadCLIValue",
			fields: fields{
				name:  "test",
				value: "test",
				usage: "test",
			},
			want: "test",
		},
		{
			name: "Test ReadCLIValue",
			fields: fields{
				name:  "",
				value: nil,
				usage: "",
			},
			panics: true,
		}, {
			name: "Test ReadCLIValue",
			fields: fields{
				name:  "test",
				value: nil,
				usage: "",
			},
			panics: true,
		},
		{
			name: "Test ReadCLIValue",
			fields: fields{
				name:  "",
				value: "test",
				usage: "",
			},
			want: "test",
		},
		{
			name: "Test ReadCLIValue",
			fields: fields{
				name:  "",
				value: 1,
				usage: "",
			},
			want: 1,
		},
		{
			name: "Test ReadCLIValue",
			fields: fields{
				name:  "",
				value: int64(1),
				usage: "",
			},
			want: int64(1),
		},
		{
			name: "Test ReadCLIValue",
			fields: fields{
				name:  "",
				value: false,
				usage: "",
			},
			want: false,
		},
		{
			name: "Test ReadCLIValue",
			fields: fields{
				name:  "",
				value: true,
				usage: "",
			},
			want: true,
		},
		{
			name: "Test ReadCLIValue",
			fields: fields{
				name:  "",
				value: struct{}{},
				usage: "",
			},
			panics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cliFlag := cliFlag{
				name:  tt.fields.name,
				value: tt.fields.value,
				usage: tt.fields.usage,
			}
			viper.Set(tt.fields.name, tt.fields.value)
			if tt.panics {
				assert.Panics(t, func() { cliFlag.ReadCLIValue() }, "ReadCLIValue()")
				return
			} else {
				assert.Equalf(t, tt.want, cliFlag.ReadCLIValue(), "ReadCLIValue()")
			}
		})
	}
}

func Test_cliFlag_Register(t *testing.T) {
	type fields struct {
		name  string
		value interface{}
		usage string
	}
	type args struct {
		command *cobra.Command
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		panics bool
	}{
		{
			name: "Test Register",
			fields: fields{
				name:  "test",
				value: "test",
				usage: "test",
			},
			args: args{
				command: &cobra.Command{},
			},
		},
		{
			name: "Test Register",
			fields: fields{
				name:  "",
				value: nil,
				usage: "",
			},
			args: args{
				command: &cobra.Command{},
			},
			panics: true,
		},
		{
			name: "Test Register",
			fields: fields{
				name:  "test",
				value: nil,
				usage: "",
			},
			args: args{
				command: &cobra.Command{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cliFlag := cliFlag{
				name:  tt.fields.name,
				value: tt.fields.value,
				usage: tt.fields.usage,
			}
			if tt.panics {
				assert.Panics(t, func() { cliFlag.Register(tt.args.command) }, "Register()")
				return
			}
		})
	}
}
