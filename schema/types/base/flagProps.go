package base

import (
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func setFlags(name string, value interface{}, usage string) {
	cflags := []helpers.CLIFlag{flags.RemoveMaintainer, flags.MutateMaintainer, flags.KafkaBool, flags.AddMaintainer}
	for _, flag := range cflags {
		if name == flag.GetName() {
			flag = base.NewCLIFlag(name, value, usage)
		}
	}
}

func GetValue(name string) interface{} {
	cflags := []helpers.CLIFlag{flags.RemoveMaintainer, flags.MutateMaintainer, flags.KafkaBool, flags.AddMaintainer}
	for _, flag := range cflags {
		if name == flag.GetName() {
			return flag.GetValue()
		}
	}

	return nil
}
