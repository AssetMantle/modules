package base

import (
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/AssetMantle/modules/schema/errors"
)

var _ errors.Error = (*sdkErrors.Error)(nil)

func NewError(codeSpace string, code uint32, description string) errors.Error {
	return sdkErrors.Register(codeSpace, code, description)
}
