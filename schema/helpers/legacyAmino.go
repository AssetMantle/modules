package helpers

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

type LegacyAmino interface {
	GetLegacyAmino() *codec.LegacyAmino
}
