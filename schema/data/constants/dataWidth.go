package constants

import (
	"github.com/cosmos/cosmos-sdk/types"
)

var (
	AccAddressDataWeight = types.NewDec(90)
	BooleanDataWeight    = types.NewDec(1)
	StringDataWeight     = types.NewDec(256)
	DecDataWeight        = types.NewDec(8)
	HeightDataWeight     = types.NewDec(8)
	IDDataWeight         = types.NewDec(64)
	ListDataWeight       = types.NewDec(1024)
)
