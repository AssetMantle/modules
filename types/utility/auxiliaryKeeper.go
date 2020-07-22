package utility

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type AuxiliaryKeeper interface {
	Help(sdkTypes.Context, AuxiliaryRequest) error
}
