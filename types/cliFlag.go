package types

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

type CLIFlag interface {
	Name() string
	Value() interface{}
	Usage() string
	ReadCLIValue() interface{}
}
type BaseCLIFlag struct {
	name  string
	value interface{}
	usage string
}

var _ CLIFlag = (*BaseCLIFlag)(nil)

func (baseCLIFlag BaseCLIFlag) Name() string       { return baseCLIFlag.name }
func (baseCLIFlag BaseCLIFlag) Value() interface{} { return baseCLIFlag.value }
func (baseCLIFlag BaseCLIFlag) Usage() string      { return baseCLIFlag.usage }
func (baseCLIFlag BaseCLIFlag) ReadCLIValue() interface{} {
	switch value := baseCLIFlag.value.(type) {
	case int:
		return viper.GetInt(baseCLIFlag.name)
	case bool:
		return viper.GetBool(baseCLIFlag.name)
	case string:
		return viper.GetString(baseCLIFlag.name)
	default:
		panic(errors.New(fmt.Sprintf("Unhandled flag type %T for flag %v", value, baseCLIFlag.name)))
	}
}

func NewCLIFlag(name string, value interface{}, usage string) CLIFlag {
	return &BaseCLIFlag{
		name:  name,
		value: value,
		usage: usage,
	}
}
