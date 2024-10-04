package helpers

import (
	"context"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/parameters"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/gogoproto/proto"
)

type Genesis interface {
	GetRecords() []Record
	GetParameters() []parameters.Parameter

	Default() Genesis

	ValidateBasic(ParameterManager) error

	Import(context.Context, Mapper, ParameterManager)
	Export(context.Context, Mapper, ParameterManager) Genesis

	Encode(sdkCodec.JSONCodec) []byte
	Decode(sdkCodec.JSONCodec, []byte) Genesis

	Initialize([]Record, lists.ParameterList) Genesis

	proto.Message
}

func ValidateGenesis[T Genesis](genesis T, parameterManager ParameterManager) error {
	if err := parameterManager.ValidateGenesisParameters(genesis); err != nil {
		return err
	}

	for _, record := range genesis.GetRecords() {
		if err := record.GetMappable().ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}
