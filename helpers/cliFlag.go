// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

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
