package types

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CLIFlag interface {
	GetName() string
	GetValue() interface{}
	Register(*cobra.Command)
	ReadCLIValue() interface{}
}
type cliFlag struct {
	Name  string
	Value interface{}
	Usage string
}

var _ CLIFlag = (*cliFlag)(nil)

func (cliFlag cliFlag) GetName() string { return cliFlag.Name }

func (cliFlag cliFlag) GetValue() interface{} { return cliFlag.Value }

func (cliFlag cliFlag) Register(command *cobra.Command) {
	switch value := cliFlag.Value.(type) {
	case int64:
		command.Flags().Int64(cliFlag.Name, value, cliFlag.Usage)
	case int:
		command.Flags().Int(cliFlag.Name, value, cliFlag.Usage)
	case bool:
		command.Flags().Bool(cliFlag.Name, value, cliFlag.Usage)
	case string:
		command.Flags().String(cliFlag.Name, value, cliFlag.Usage)
	default:
		panic(value)
	}
}

func (cliFlag cliFlag) ReadCLIValue() interface{} {
	switch value := cliFlag.Value.(type) {
	case int64:
		return viper.GetInt64(cliFlag.Name)
	case int:
		return viper.GetInt(cliFlag.Name)
	case bool:
		return viper.GetBool(cliFlag.Name)
	case string:
		return viper.GetString(cliFlag.Name)
	default:
		panic(errors.New(fmt.Sprintf("Unhandled flag type %T for flag %v", value, cliFlag.Name)))
	}
}

func NewCLIFlag(name string, value interface{}, usage string) CLIFlag {
	return cliFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}
