/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/spf13/cobra"
)

type CLIFlag interface {
	GetName() string
	GetValue() interface{}
	Register(*cobra.Command)
	ReadCLIValue() interface{}
}
