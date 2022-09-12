// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func initialize() (helpers.CLICommand, []helpers.CLIFlag) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	cliContext := context.NewCLIContext().WithCodec(Codec)
	cliContext = cliContext.WithChainID("chainID")

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
		// TODO: Add test cases.
		{"+ve", args{"", "", "", testCliFlagList}, cliCommand{"", "", "", testCliFlagList}},
		{"nil", args{"", "", "", nil}, cliCommand{"", "", "", nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCLICommand(tt.args.use, tt.args.short, tt.args.long, tt.args.cliFlagList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCLICommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cliCommand_CreateCommand(t *testing.T) {
	//_, testCliFlagList := initialize()
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
			if got := cliCommand.CreateCommand(tt.args.runE); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCommand() = %v, want %v", got, tt.want)
			}
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
		cliContext context.CLIContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   rest.BaseReq
	}{
		// TODO: Add test cases.
		{"+ve", fields{"", "", "", testCliFlagList}, args{context.CLIContext{ChainID: "chainID"}}, rest.BaseReq{ChainID: "chainID"}},
		{"-ve for nil", fields{"", "", "", nil}, args{context.CLIContext{ChainID: ""}}, rest.BaseReq{ChainID: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cliCommand := cliCommand{
				use:         tt.fields.use,
				short:       tt.fields.short,
				long:        tt.fields.long,
				cliFlagList: tt.fields.cliFlagList,
			}
			if got := cliCommand.ReadBaseReq(tt.args.cliContext); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadBaseReq() = %v, want %v", got, tt.want)
			}
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
		// TODO: Add test cases.
		{"+ve", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name4", false, ",usage")}, false},
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
		// TODO: Add test cases.
		{"-ve flag name not an int flag", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name", "value", ",usage")}, 0, true},
		{"-ve unregistered flag", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name", 1, ",usage")}, 1, true},
		{"+ve", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name3", 123, ",usage")}, 0, false},
		{"-ve should panic", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name4", struct{}{}, ",usage")}, 0, true},
		//{"-ve should not panic", fields{"", "", "", nil}, args{NewCLIFlag("name4", 123, ",usage")}, 0, false},
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
		// TODO: Add test cases.
		{"-ve flag name not an int64 flag", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name", "value", ",usage")}, 1, true},
		{"-ve unregistered flag", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name2", int64(1), ",usage")}, 1, true},
		{"+ve", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name2", int64(-1), ",usage")}, 0, false},
		{"-ve should panic", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name4", struct{}{}, ",usage")}, 0, true},
		{"-ve should panic for nil", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name2", nil, ",usage")}, 0, true},
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
		// TODO: Add test cases.
		{"+ve", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name", "value", ",usage")}, "", false},
		{"-ve unregistered flag", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name", "value1", ",usage")}, "", true},
		{"-ve should panic", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name3", 1, ",usage")}, "0", true},
		{"-ve flag name not a string flag", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name2", int64(-1), ",usage")}, "0", true},
		{"-ve should panic for nil", fields{"", "", "", testCLiFlagList}, args{NewCLIFlag("name2", nil, ",usage")}, "0", true}}
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
