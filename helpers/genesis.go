// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"
	"github.com/AssetMantle/schema/lists"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/gogoproto/proto"
)

type Genesis interface {
	GetRecords() []Record
	GetParameterList() lists.ParameterList

	SetRecords([]Record) Genesis
	SetParameters(list lists.ParameterList) Genesis

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
	if err := parameterManager.Set(genesis.GetParameterList().Get()...).Validate(); err != nil {
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

	if _, err := parameterManager.Set(genesis.GetParameterList().Get()...).Update(context); err != nil {
		panic(err)
	}
}

func ExportGenesis[T Genesis](genesis T, context context.Context, mapper Mapper, parameterManager ParameterManager) Genesis {
	return genesis.Initialize(mapper.NewCollection(context).FetchAll().Get(), parameterManager.Fetch(context).Get())
}

func EncodeGenesis[T Genesis](genesis T, jsonCodec sdkCodec.JSONCodec) []byte {
	bytes, err := jsonCodec.MarshalJSON(genesis)
	if err != nil {
		panic(err)
	}

	return bytes
}
func DecodeGenesis[T Genesis](genesis T, jsonCodec sdkCodec.JSONCodec, byte []byte) Genesis {
	if err := jsonCodec.UnmarshalJSON(byte, genesis); err != nil {
		panic(err)
	}

	return genesis
}

func InitializeGenesis[T Genesis](genesis T, records []Record, parameterList lists.ParameterList) Genesis {
	if len(records) == 0 {
		records = genesis.Default().GetRecords()
	}

	if len(parameterList.Get()) == 0 {
		parameterList = genesis.Default().GetParameterList()
	} else {
		parameterList = genesis.Default().GetParameterList().Mutate(parameterList.Get()...)
	}

	if err := parameterList.ValidateBasic(); err != nil {
		panic(err)
	}

	return genesis.SetRecords(records).SetParameters(parameterList)
}
