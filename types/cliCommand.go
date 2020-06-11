package types

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

type CLICommand struct {
	Use         string
	Short       string
	Long        string
	CLIFlagList []CLIFlag
}

func (cliCommand CLICommand) RegisterFlags(command *cobra.Command) {
	for _, cliFlag := range cliCommand.CLIFlagList {
		switch value := cliFlag.Value.(type) {
		case int:
			command.Flags().Int(cliFlag.Name, cliFlag.Value.(int), cliFlag.Usage)
		case bool:
			command.Flags().Bool(cliFlag.Name, cliFlag.Value.(bool), cliFlag.Usage)
		case string:
			command.Flags().String(cliFlag.Name, cliFlag.Value.(string), cliFlag.Usage)
		default:
			panic(value)
		}
	}
}

func (cliCommand CLICommand) GetInt(cliFlag CLIFlag) int {
	switch _ := cliFlag.Value.(type) {
	case int:
		for _, registeredCliFlag := range cliCommand.CLIFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadValue().(int)
			}
		}
	default:
		panic(errors.New(fmt.Sprintf("Falg %v not an int flag, Flag type: %T, ", cliFlag.Name, cliFlag.Value)))
	}
	panic(errors.New(fmt.Sprintf("Uregistered flag %v type %T", cliFlag.Name, cliFlag.Value)))
}

func (cliCommand CLICommand) GetBool(cliFlag CLIFlag) bool {
	switch _ := cliFlag.Value.(type) {
	case bool:
		for _, registeredCliFlag := range cliCommand.CLIFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadValue().(bool)
			}
		}
	default:
		panic(errors.New(fmt.Sprintf("Falg %v not an bool flag, Flag type: %T, ", cliFlag.Name, cliFlag.Value)))
	}
	panic(errors.New(fmt.Sprintf("Uregistered flag %v type %T", cliFlag.Name, cliFlag.Value)))
}

func (cliCommand CLICommand) GetString(cliFlag CLIFlag) string {
	switch _ := cliFlag.Value.(type) {
	case string:
		for _, registeredCliFlag := range cliCommand.CLIFlagList {
			if registeredCliFlag == cliFlag {
				return cliFlag.ReadValue().(string)
			}
		}
	default:
		panic(errors.New(fmt.Sprintf("Falg %v not an string flag, Flag type: %T, ", cliFlag.Name, cliFlag.Value)))
	}
	panic(errors.New(fmt.Sprintf("Uregistered flag %v type %T", cliFlag.Name, cliFlag.Value)))
}
