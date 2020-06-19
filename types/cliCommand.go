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
	Use() string
	Short() string
	Long() string
	RegisterFlags(*cobra.Command)

	FromAddress() sdkTypes.AccAddress
	CreateTransactionCommand(*codec.Codec, func(CLICommand) sdkTypes.Msg) *cobra.Command
	CreateQueryCommand(*codec.Codec, string, func(CLICommand) []byte, func([]byte) interface{}) *cobra.Command
	GetInt(CLIFlag) int
	GetBool(CLIFlag) bool
	GetString(CLIFlag) string
}

type BaseCLICommand struct {
	use         string
	short       string
	long        string
	fromAddress sdkTypes.AccAddress
	cliFlagList []CLIFlag
}

var _ CLICommand = (*BaseCLICommand)(nil)

func (baseCLICommand BaseCLICommand) Use() string   { return baseCLICommand.use }
func (baseCLICommand BaseCLICommand) Short() string { return baseCLICommand.short }
func (baseCLICommand BaseCLICommand) Long() string  { return baseCLICommand.long }
func (baseCLICommand BaseCLICommand) FromAddress() sdkTypes.AccAddress {
	return baseCLICommand.fromAddress
}

func (baseCLICommand BaseCLICommand) CreateTransactionCommand(codec *codec.Codec, makeMessage func(CLICommand) sdkTypes.Msg) *cobra.Command {
	command := &cobra.Command{
		Use:   baseCLICommand.Use(),
		Short: baseCLICommand.Short(),
		Long:  baseCLICommand.Long(),
		RunE: func(command *cobra.Command, args []string) error {
			bufioReader := bufio.NewReader(command.InOrStdin())
			transactionBuilder := auth.NewTxBuilderFromCLI(bufioReader).WithTxEncoder(authClient.GetTxEncoder(codec))
			cliContext := context.NewCLIContextWithInput(bufioReader).WithCodec(codec)
			baseCLICommand.fromAddress = cliContext.GetFromAddress()
			message := makeMessage(baseCLICommand)
			if Error := message.ValidateBasic(); Error != nil {
				return Error
			}

			return authClient.GenerateOrBroadcastMsgs(cliContext, transactionBuilder, []sdkTypes.Msg{message})
		},
	}
	baseCLICommand.RegisterFlags(command)
	return flags.PostCommands(command)[0]
}

func (baseCLICommand BaseCLICommand) CreateQueryCommand(codec *codec.Codec, queryRoute string, makeQueryBytes func(CLICommand) []byte, marshallResponse func([]byte) interface{}) *cobra.Command {
	command := &cobra.Command{
		Use:   baseCLICommand.Use(),
		Short: baseCLICommand.Short(),
		Long:  baseCLICommand.Long(),
		RunE: func(command *cobra.Command, args []string) error {
			cliContext := context.NewCLIContext().WithCodec(codec)

			bytes := makeQueryBytes(baseCLICommand)
			responseBytes, _, queryWithDataError := cliContext.QueryWithData(strings.Join([]string{"", "custom", queryRoute, baseCLICommand.Use()}, "/"), bytes)
			if queryWithDataError != nil {
				return queryWithDataError
			}
			response := marshallResponse(responseBytes)
			return cliContext.PrintOutput(response)
		},
	}

	baseCLICommand.RegisterFlags(command)
	return flags.GetCommands(command)[0]
}
func (baseCLICommand BaseCLICommand) RegisterFlags(command *cobra.Command) {
	for _, cliFlag := range baseCLICommand.cliFlagList {
		switch value := cliFlag.Value().(type) {
		case int:
			command.Flags().Int(cliFlag.Name(), cliFlag.Value().(int), cliFlag.Usage())
		case bool:
			command.Flags().Bool(cliFlag.Name(), cliFlag.Value().(bool), cliFlag.Usage())
		case string:
			command.Flags().String(cliFlag.Name(), cliFlag.Value().(string), cliFlag.Usage())
		default:
			panic(value)
		}
	}
}

func (baseCLICommand BaseCLICommand) GetInt(cliFlag CLIFlag) int {
	switch cliFlag.Value().(type) {
	case int:
		for _, registeredCliFlag := range baseCLICommand.cliFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(int)
			}
		}
	default:
		panic(errors.New(fmt.Sprintf("Falg %v not an int flag, Flag type: %T, ", cliFlag.Name(), cliFlag.Value())))
	}
	panic(errors.New(fmt.Sprintf("Uregistered flag %v type %T", cliFlag.Name(), cliFlag.Value())))
}

func (baseCLICommand BaseCLICommand) GetBool(cliFlag CLIFlag) bool {
	switch cliFlag.Value().(type) {
	case bool:
		for _, registeredCliFlag := range baseCLICommand.cliFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(bool)
			}
		}
	default:
		panic(errors.New(fmt.Sprintf("Falg %v not an bool flag, Flag type: %T, ", cliFlag.Name(), cliFlag.Value())))
	}
	panic(errors.New(fmt.Sprintf("Uregistered flag %v type %T", cliFlag.Name(), cliFlag.Value())))
}

func (baseCLICommand BaseCLICommand) GetString(cliFlag CLIFlag) string {
	switch cliFlag.Value().(type) {
	case string:
		for _, registeredCliFlag := range baseCLICommand.cliFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(string)
			}
		}
	default:
		panic(errors.New(fmt.Sprintf("Falg %v not an string flag, Flag type: %T, ", cliFlag.Name(), cliFlag.Value())))
	}
	panic(errors.New(fmt.Sprintf("Uregistered flag %v type %T", cliFlag.Name(), cliFlag.Value())))
}

func NewCLICommand(use string, short string, long string, cliFlagList []CLIFlag) CLICommand {
	return &BaseCLICommand{
		use:         use,
		short:       short,
		long:        long,
		cliFlagList: cliFlagList,
	}
}
