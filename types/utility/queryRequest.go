package utility

import (
	"github.com/cosmos/cosmos-sdk/client/context"
)

type QueryRequest interface {
	Request
	FromCLI(CLICommand, context.CLIContext) QueryRequest
	FromMap(map[string]string) QueryRequest
}
