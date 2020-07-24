package helpers

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/spf13/cobra"
)

type CLICommand interface {
	ReadInt64(CLIFlag) int64
	ReadInt(CLIFlag) int
	ReadBool(CLIFlag) bool
	ReadString(CLIFlag) string
	ReadBaseReq(context.CLIContext) rest.BaseReq

	CreateCommand(func(command *cobra.Command, args []string) error) *cobra.Command
}
