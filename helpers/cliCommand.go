// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/spf13/cobra"
)

type CLICommand interface {
	ReadInt64(CLIFlag) int64
	ReadInt(CLIFlag) int
	ReadBool(CLIFlag) bool
	ReadString(CLIFlag) string
	ReadBaseReq(client.Context) rest.BaseReq

	CreateCommand(func(command *cobra.Command, args []string) error) *cobra.Command
}
