package base

import (
	"github.com/AssetMantle/modules/helpers"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ helpers.Error = (*sdkErrors.Error)(nil)

func NewError(codeSpace string, code uint32, description string) helpers.Error {
	return sdkErrors.Register(codeSpace, code, description)
}
