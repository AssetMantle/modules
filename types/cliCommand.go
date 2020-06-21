package types

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authClient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/spf13/cobra"
	"strings"
)

type CLICommand interface {
	registerFlags(*cobra.Command)

	GetFromAddress() sdkTypes.AccAddress
	ReadInt(CLIFlag) int
	ReadBool(CLIFlag) bool
	ReadString(CLIFlag) string

	CreateTransactionCommand(*codec.Codec, func(CLICommand) sdkTypes.Msg) *cobra.Command
	CreateQueryCommand(*codec.Codec, string, func(CLICommand) []byte, func([]byte) interface{}) *cobra.Command
}

type cliCommand struct {
	Use         string
	Short       string
	Long        string
	FromAddress sdkTypes.AccAddress
	CLIFlagList []CLIFlag
}

var _ CLICommand = (*cliCommand)(nil)

func (cliCommand cliCommand) GetFromAddress() sdkTypes.AccAddress {
	return cliCommand.FromAddress
}

func (cliCommand cliCommand) registerFlags(command *cobra.Command) {
	for _, cliFlag := range cliCommand.CLIFlagList {
		cliFlag.Register(command)
	}
}

func (cliCommand cliCommand) ReadInt(cliFlag CLIFlag) int {
	switch cliFlag.GetValue().(type) {
	case int:
		for _, registeredCliFlag := range cliCommand.CLIFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(int)
			}
		}
	default:
		panic(errors.New(fmt.Sprintf("Flag %v not an int flag, Flag type: %T, ", cliFlag.GetName(), cliFlag.GetValue())))
	}
	panic(errors.New(fmt.Sprintf("Uregistered flag %v type %T", cliFlag.GetName(), cliFlag.GetValue())))
}

func (cliCommand cliCommand) ReadBool(cliFlag CLIFlag) bool {
	switch cliFlag.GetValue().(type) {
	case bool:
		for _, registeredCliFlag := range cliCommand.CLIFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(bool)
			}
		}
	default:
		panic(errors.New(fmt.Sprintf("Falg %v not an bool flag, Flag type: %T, ", cliFlag.GetName(), cliFlag.GetValue())))
	}
	panic(errors.New(fmt.Sprintf("Uregistered flag %v type %T", cliFlag.GetName(), cliFlag.GetValue())))
}

func (cliCommand cliCommand) ReadString(cliFlag CLIFlag) string {
	switch cliFlag.GetValue().(type) {
	case string:
		for _, registeredCliFlag := range cliCommand.CLIFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(string)
			}
		}
	default:
		panic(errors.New(fmt.Sprintf("Falg %v not an string flag, Flag type: %T, ", cliFlag.GetName(), cliFlag.GetValue())))
	}
	panic(errors.New(fmt.Sprintf("Uregistered flag %v type %T", cliFlag.GetName(), cliFlag.GetValue())))
}

func (cliCommand cliCommand) CreateTransactionCommand(codec *codec.Codec, makeMessage func(CLICommand) sdkTypes.Msg) *cobra.Command {
	command := &cobra.Command{
		Use:   cliCommand.Use,
		Short: cliCommand.Short,
		Long:  cliCommand.Long,
		RunE: func(command *cobra.Command, args []string) error {
			bufioReader := bufio.NewReader(command.InOrStdin())
			transactionBuilder := auth.NewTxBuilderFromCLI(bufioReader).WithTxEncoder(authClient.GetTxEncoder(codec))
			cliContext := context.NewCLIContextWithInput(bufioReader).WithCodec(codec)
			cliCommand.FromAddress = cliContext.GetFromAddress()
			message := makeMessage(cliCommand)
			if Error := message.ValidateBasic(); Error != nil {
				return Error
			}

			return authClient.GenerateOrBroadcastMsgs(cliContext, transactionBuilder, []sdkTypes.Msg{message})
		},
	}
	cliCommand.registerFlags(command)
	return flags.PostCommands(command)[0]
}

func (cliCommand cliCommand) CreateQueryCommand(codec *codec.Codec, queryRoute string, makeQueryBytes func(CLICommand) []byte, marshallResponse func([]byte) interface{}) *cobra.Command {
	command := &cobra.Command{
		Use:   cliCommand.Use,
		Short: cliCommand.Short,
		Long:  cliCommand.Long,
		RunE: func(command *cobra.Command, args []string) error {
			cliContext := context.NewCLIContext().WithCodec(codec)

			bytes := makeQueryBytes(cliCommand)
			responseBytes, _, queryWithDataError := cliContext.QueryWithData(strings.Join([]string{"", "custom", queryRoute, cliCommand.Use}, "/"), bytes)
			if queryWithDataError != nil {
				return queryWithDataError
			}
			response := marshallResponse(responseBytes)
			return cliContext.PrintOutput(response)
		},
	}

	cliCommand.registerFlags(command)
	return flags.GetCommands(command)[0]
}

func NewCLICommand(use string, short string, long string, cliFlagList []CLIFlag) CLICommand {
	return &cliCommand{
		Use:         use,
		Short:       short,
		Long:        long,
		CLIFlagList: cliFlagList,
	}
}
