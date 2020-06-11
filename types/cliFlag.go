package types

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

type CLIFlag struct {
	Name  string
	Value interface{}
	Usage string
}

func (cliFlag CLIFlag) ReadValue() interface{} {
	switch value := cliFlag.Value.(type) {
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
