/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"fmt"

	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type cliFlag struct {
	name  string
	value interface{}
	usage string
}

var _ helpers.CLIFlag = (*cliFlag)(nil)

func (cliFlag cliFlag) GetName() string { return cliFlag.name }

func (cliFlag cliFlag) GetValue() interface{} { return cliFlag.value }

func (cliFlag cliFlag) Register(command *cobra.Command) {
	switch value := cliFlag.value.(type) {
	case int64:
		command.Flags().Int64(cliFlag.name, value, cliFlag.usage)
	case int:
		command.Flags().Int(cliFlag.name, value, cliFlag.usage)
	case bool:
		command.Flags().Bool(cliFlag.name, value, cliFlag.usage)
	case string:
		command.Flags().String(cliFlag.name, value, cliFlag.usage)
	default:
		panic(value)
	}
}

func (cliFlag cliFlag) ReadCLIValue() interface{} {
	switch value := cliFlag.value.(type) {
	case int64:
		return viper.GetInt64(cliFlag.name)
	case int:
		return viper.GetInt(cliFlag.name)
	case bool:
		return viper.GetBool(cliFlag.name)
	case string:
		return viper.GetString(cliFlag.name)
	default:
		panic(fmt.Errorf("unhandled flag type %T for flag %v", value, cliFlag.name))
	}
}

func NewCLIFlag(name string, value interface{}, usage string) helpers.CLIFlag {
	return cliFlag{
		name:  name,
		value: value,
		usage: usage,
	}
}
