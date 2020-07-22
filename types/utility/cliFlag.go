package utility

import (
	"github.com/spf13/cobra"
)

type CLIFlag interface {
	GetName() string
	GetValue() interface{}
	Register(*cobra.Command)
	ReadCLIValue() interface{}
}
