// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
)

func initialize() (helpers.CLICommand, []helpers.CLIFlag) {
	testCliFlag := NewCLIFlag("name", "value", ",usage")
	testCliFlag2 := NewCLIFlag("name2", int64(-1), ",usage")
	testCliFlag3 := NewCLIFlag("name3", 123, ",usage")
	testCliFlag4 := NewCLIFlag("name4", false, ",usage")
	testCliFlagList := []helpers.CLIFlag{testCliFlag, testCliFlag2, testCliFlag3, testCliFlag4}
	testCliCommand := NewCLICommand("", "", "", testCliFlagList).(cliCommand)
	return testCliCommand, testCliFlagList
}

func TestNewCLICommand(t *testing.T) {
	_, testCliFlagList := initialize()
	type args struct {
		use         string
		short       string
		long        string
		cliFlagList []helpers.CLIFlag
	}
	tests := []struct {
		name string
		args args
		want helpers.CLICommand
	}{

		{"valid", args{"", "", "", testCliFlagList}, cliCommand{"", "", "", testCliFlagList}},
		{"nil", args{"", "", "", nil}, cliCommand{"", "", "", nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCLICommand(tt.args.use, tt.args.short, tt.args.long, tt.args.cliFlagList)
			assert.Equal(t, tt.want, got, "NewCLICommand()")
		})
	}
}

func Test_cliCommand_CreateCommand(t *testing.T) {
	// _, testCliFlagList := initialize()
	type fields struct {
		use         string
		short       string
		long        string
		cliFlagList []helpers.CLIFlag
	}
	type args struct {
		runE func(command *cobra.Command, args []string) error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *cobra.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cliCommand := cliCommand{
				use:         tt.fields.use,
				short:       tt.fields.short,
				long:        tt.fields.long,
				cliFlagList: tt.fields.cliFlagList,
			}
			got := cliCommand.CreateCommand(tt.args.runE)
			assert.Equal(t, tt.want, got, "CreateCommand()")
		})
	}
}

func Test_cliCommand_ReadBaseReq(t *testing.T) {
	_, testCliFlagList := initialize()

	type fields struct {
		use         string
		short       string
		long        string
		cliFlagList []helpers.CLIFlag
	}
	type args struct {
		context client.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.CommonTransactionRequest
	}{

		{"valid", fields{"", "", "", testCliFlagList}, args{client.Context{ChainID: "chainID"}}, helpers.CommonTransactionRequest{ChainID: "chainID"}},
		{"nil input", fields{"", "", "", nil}, args{client.Context{ChainID: ""}}, helpers.PrototypeCommonTransactionRequest()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cliCommand := cliCommand{
				use:         tt.fields.use,
				short:       tt.fields.short,
				long:        tt.fields.long,
				cliFlagList: tt.fields.cliFlagList,
			}
			got := cliCommand.ReadCommonTransactionRequest(tt.args.context)
			assert.Equal(t, tt.want, got, "ReadCommonTransactionRequest()")
		})
	}
}

func Test_cliCommand_ReadBool(t *testing.T) {
	testCliCommand, testCLiFlagList := initialize()

	type fields struct {
		use         string
		short       string
		long        string
		cliFlagList []helpers.CLIFlag
	}
	type args struct {
		cliFlag helpers.CLIFlag
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{

		{"valid", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name4", false, ",usage")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cliCommand := cliCommand{
				use:         tt.fields.use,
				short:       tt.fields.short,
				long:        tt.fields.long,
				cliFlagList: tt.fields.cliFlagList,
			}
			if got := cliCommand.ReadBool(tt.args.cliFlag); got != tt.want {
				t.Errorf("ReadBool() = %v, want %v", got, tt.want)
			}
		})
	}
	require.Panics(t, func() {
		testCliCommand.ReadString(NewCLIFlag("name", 1, ",usage"))
	})
}

func Test_cliCommand_ReadInt(t *testing.T) {
	_, testCLiFlagList := initialize()
	type fields struct {
		use         string
		short       string
		long        string
		cliFlagList []helpers.CLIFlag
	}
	type args struct {
		cliFlag helpers.CLIFlag
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        int
		shouldPanic bool
	}{

		{"flag name not int", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name", "value", ",usage")}, 0, true},
		{"unregistered flag", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name", 1, ",usage")}, 1, true},
		{"valid", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name3", 123, ",usage")}, 0, false},
		{"should panic", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name4", struct{}{}, ",usage")}, 0, true},
		// {"should not panic", fields{"", "", "", nil}, args{NewCLIFlag("name4", 123, ",usage")}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cliCommand := cliCommand{
				use:         tt.fields.use,
				short:       tt.fields.short,
				long:        tt.fields.long,
				cliFlagList: tt.fields.cliFlagList,
			}
			if tt.shouldPanic {
				assert.Panics(t, func() { cliCommand.ReadInt(tt.args.cliFlag) }, "The code did not panic, but it should panic")
			} else {
				if got := cliCommand.ReadInt(tt.args.cliFlag); got != tt.want {
					t.Errorf("ReadInt() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_cliCommand_ReadInt64(t *testing.T) {
	_, testCLiFlagList := initialize()
	type fields struct {
		use         string
		short       string
		long        string
		cliFlagList []helpers.CLIFlag
	}
	type args struct {
		cliFlag helpers.CLIFlag
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        int64
		shouldPanic bool
	}{

		{"flag name not int64", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name", "value", ",usage")}, 1, true},
		{"unregistered flag", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name2", int64(1), ",usage")}, 1, true},
		{"valid", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name2", int64(-1), ",usage")}, 0, false},
		{"should panic", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name4", struct{}{}, ",usage")}, 0, true},
		{"panic on nil", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name2", nil, ",usage")}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cliCommand := cliCommand{
				use:         tt.fields.use,
				short:       tt.fields.short,
				long:        tt.fields.long,
				cliFlagList: tt.fields.cliFlagList,
			}
			if tt.shouldPanic {
				assert.Panics(t, func() { cliCommand.ReadInt(tt.args.cliFlag) }, "The code did not panic, but it should panic")
			} else {
				if got := cliCommand.ReadInt64(tt.args.cliFlag); got != tt.want {
					t.Errorf("ReadInt64() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_cliCommand_ReadString(t *testing.T) {
	_, testCLiFlagList := initialize()
	type fields struct {
		use         string
		short       string
		long        string
		cliFlagList []helpers.CLIFlag
	}
	type args struct {
		cliFlag helpers.CLIFlag
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        string
		shouldPanic bool
	}{

		{"valid", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name", "value", ",usage")}, "", false},
		{"unregistered flag", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name", "value1", ",usage")}, "", true},
		{"should panic", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name3", 1, ",usage")}, "0", true},
		{"flag name not string", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name2", int64(-1), ",usage")}, "0", true},
		{"panic on nil", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name2", nil, ",usage")}, "0", true}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cliCommand := cliCommand{
				use:         tt.fields.use,
				short:       tt.fields.short,
				long:        tt.fields.long,
				cliFlagList: tt.fields.cliFlagList,
			}
			if tt.shouldPanic {
				assert.Panics(t, func() { cliCommand.ReadInt(tt.args.cliFlag) }, "The code did not panic, but it should panic")
			} else {
				if got := cliCommand.ReadString(tt.args.cliFlag); got != tt.want {
					t.Errorf("ReadString() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_cliCommand_registerFlags(t *testing.T) {
	_, testCLiFlagList := initialize()
	testCliCommand := NewCLICommand("", "", "", testCLiFlagList).(cliCommand)

	require.NotPanics(t, func() {
		testCliCommand.registerFlags(&cobra.Command{})
	})
}
