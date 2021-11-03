package helpers

import paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"

type Parameters interface {
	Validate() error
	paramTypes.ParamSet
}
