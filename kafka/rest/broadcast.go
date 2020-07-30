package rest

import (
	//cTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"

	context "github.com/cosmos/cosmos-sdk/client/context"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmos/cosmos-sdk/x/auth"
)

const DefaultCodeSpace = "commit"

func BroadcastRest(cliCtx context.CLIContext, cdc *codec.Codec, stdTx auth.StdTx, mode string) ([]byte, error) {

	txBytes, err := cdc.MarshalBinaryLengthPrefixed(stdTx)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	cliCtx = cliCtx.WithBroadcastMode(mode)

	res, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return PostProcessResponse(cliCtx, res)
}
