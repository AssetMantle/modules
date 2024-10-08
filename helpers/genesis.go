package helpers

import (
	"context"
	parametersSchema "github.com/AssetMantle/schema/parameters"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/gogoproto/proto"
)

type Genesis interface {
	GetRecords() []Record
	GetParameters() []parametersSchema.Parameter

	SetRecords([]Record) Genesis
	SetParameters([]parametersSchema.Parameter) Genesis

	Default() Genesis

	ValidateBasic(ParameterManager) error

	Import(context.Context, Mapper, ParameterManager)
	Export(context.Context, Mapper, ParameterManager) Genesis

	Encode(sdkCodec.JSONCodec) []byte
	Decode(sdkCodec.JSONCodec, []byte) Genesis

	Initialize([]Record, []parametersSchema.Parameter) Genesis

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

func ImportGenesis[T Genesis](genesis T, context context.Context, mapper Mapper, parameterManager ParameterManager) {
	for _, record := range genesis.GetRecords() {
		mapper.NewCollection(context).Add(record)
	}

	parameterManager.Set(context, genesis.GetParameters())
}
