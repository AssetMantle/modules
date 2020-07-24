package base

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/spf13/cobra"
)

type cliCommand struct {
	use         string
	short       string
	long        string
	cliFlagList []helpers.CLIFlag
}

var _ helpers.CLICommand = (*cliCommand)(nil)

func (cliCommand cliCommand) registerFlags(command *cobra.Command) {
	for _, cliFlag := range cliCommand.cliFlagList {
		cliFlag.Register(command)
	}
}

func (cliCommand cliCommand) ReadInt64(cliFlag helpers.CLIFlag) int64 {
	switch cliFlag.GetValue().(type) {
	case int64:
		for _, registeredCliFlag := range cliCommand.cliFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(int64)
			}
		}
	default:
		panic(errors.New(fmt.Sprintf("Flag %v not an int64 flag, Flag type: %T, ", cliFlag.GetName(), cliFlag.GetValue())))
	}
	panic(errors.New(fmt.Sprintf("Uregistered flag %v type %T", cliFlag.GetName(), cliFlag.GetValue())))
}

func (cliCommand cliCommand) ReadInt(cliFlag helpers.CLIFlag) int {
	switch cliFlag.GetValue().(type) {
	case int:
		for _, registeredCliFlag := range cliCommand.cliFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(int)
			}
		}
	default:
		panic(errors.New(fmt.Sprintf("Flag %v not an int flag, Flag type: %T, ", cliFlag.GetName(), cliFlag.GetValue())))
	}
	panic(errors.New(fmt.Sprintf("Uregistered flag %v type %T", cliFlag.GetName(), cliFlag.GetValue())))
}

func (cliCommand cliCommand) ReadBool(cliFlag helpers.CLIFlag) bool {
	switch cliFlag.GetValue().(type) {
	case bool:
		for _, registeredCliFlag := range cliCommand.cliFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(bool)
			}
		}
	default:
		panic(errors.New(fmt.Sprintf("Falg %v not an bool flag, Flag type: %T, ", cliFlag.GetName(), cliFlag.GetValue())))
	}
	panic(errors.New(fmt.Sprintf("Uregistered flag %v type %T", cliFlag.GetName(), cliFlag.GetValue())))
}

func (cliCommand cliCommand) ReadString(cliFlag helpers.CLIFlag) string {
	switch cliFlag.GetValue().(type) {
	case string:
		for _, registeredCliFlag := range cliCommand.cliFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadCLIValue().(string)
			}
		}
	default:
		panic(errors.New(fmt.Sprintf("Falg %v not an string flag, Flag type: %T, ", cliFlag.GetName(), cliFlag.GetValue())))
	}
	panic(errors.New(fmt.Sprintf("Uregistered flag %v type %T", cliFlag.GetName(), cliFlag.GetValue())))
}

func (cliCommand cliCommand) ReadBaseReq(cliContext context.CLIContext) rest.BaseReq {
	return rest.BaseReq{
		From:     cliContext.GetFromAddress().String(),
		ChainID:  cliContext.ChainID,
		Simulate: cliContext.Simulate,
	}
}
func (cliCommand cliCommand) CreateCommand(runE func(command *cobra.Command, args []string) error) *cobra.Command {
	command := &cobra.Command{
		Use:   cliCommand.use,
		Short: cliCommand.short,
		Long:  cliCommand.long,
		RunE:  runE,
	}
	cliCommand.registerFlags(command)
	return flags.PostCommands(command)[0]
}

func NewCLICommand(use string, short string, long string, cliFlagList []helpers.CLIFlag) helpers.CLICommand {
	return cliCommand{
		use:         use,
		short:       short,
		long:        long,
		cliFlagList: cliFlagList,
	}
}
