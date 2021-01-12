/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Cli_Command(t *testing.T) {
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
	require.Equal(t, cliCommand{use: "", short: "", long: "", cliFlagList: testCliFlagList}, testCliCommand)

	require.NotPanics(t, func() {
		testCliCommand.registerFlags(&cobra.Command{})
	})
	require.Equal(t, "", testCliCommand.ReadString(testCliFlag))
	require.Panics(t, func() {
		testCliCommand.ReadString(NewCLIFlag("name", 1, ",usage"))
	})

	require.Equal(t, int64(0), testCliCommand.ReadInt64(testCliFlag2))
	require.Panics(t, func() {
		testCliCommand.ReadInt64(NewCLIFlag("name", "1", ",usage"))
	})
	require.Equal(t, 0, testCliCommand.ReadInt(testCliFlag3))
	require.Panics(t, func() {
		testCliCommand.ReadInt(NewCLIFlag("name", "1", ",usage"))
	})
	require.Equal(t, false, testCliCommand.ReadBool(testCliFlag4))
	require.Panics(t, func() {
		testCliCommand.ReadBool(NewCLIFlag("name", "1", ",usage"))
	})
	require.Equal(t, rest.BaseReq{ChainID: "chainID"}, testCliCommand.ReadBaseReq(cliContext))

}
