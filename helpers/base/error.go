package base

import (
	sdkErrors "cosmossdk.io/errors"
	"github.com/AssetMantle/modules/helpers"
)

var _ helpers.Error = (*sdkErrors.Error)(nil)

func NewError(codeSpace string, code uint32, description string) helpers.Error {
	return sdkErrors.Register(codeSpace, code, description)
}
