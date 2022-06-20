package base

import (
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/AssetMantle/modules/schema/helpers"
)

var _ helpers.Error = (*errors.Error)(nil)

func NewError(codeSpace string, code uint32, description string) helpers.Error {
	return errors.Register(codeSpace, code, description)
}
