// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"context"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/identities/internal/common"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/identities/internal/mapper"
	"github.com/AssetMantle/modules/schema"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
)

func CreateTestInputContext(t *testing.T) context.Context {
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	return sdkTypes.WrapSDKContext(context)
}

func Test_newQueryResponse(t *testing.T) {
	context := CreateTestInputContext(t)
	collection := mapper.Prototype().NewCollection(context)
	type args struct {
		collection helpers.Collection
		error      error
	}
	tests := []struct {
		name string
		args args
		want helpers.QueryResponse
	}{

		{"+ve", args{collection: collection, error: nil}, &QueryResponse{Success: true, Error: ""}},
		{"-ve with error", args{collection: collection, error: errorConstants.IncorrectFormat}, &QueryResponse{Success: false, Error: errorConstants.IncorrectFormat.Error()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newQueryResponse(tt.args.collection, tt.args.error); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newQueryResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResponse_Decode(t *testing.T) {
	context := CreateTestInputContext(t)
	collection := mapper.Prototype().NewCollection(context)
	testQueryResponse := newQueryResponse(collection, nil)
	encodedResponse, _ := testQueryResponse.Encode()
	type fields struct {
		Success bool
		Error   string
		List    []helpers.Mappable
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.QueryResponse
		wantErr bool
	}{

		{"+ve", fields{Success: true, Error: ""}, args{bytes: encodedResponse}, testQueryResponse, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryResponse := &QueryResponse{
				Success: tt.fields.Success,
				Error:   tt.fields.Error,
				List:    mappable.MappablesFromInterface(tt.fields.List),
			}
			got, err := queryResponse.Decode(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResponse_Encode(t *testing.T) {
	context := CreateTestInputContext(t)
	collection := mapper.Prototype().NewCollection(context)
	encodedByte, err := common.LegacyAmino.MarshalJSON(&QueryResponse{Success: true, Error: "", List: mappable.MappablesFromInterface(collection.GetList())})
	encodedByteWithError, _err := common.LegacyAmino.MarshalJSON(&QueryResponse{Success: false, Error: errorConstants.IncorrectFormat.Error(), List: mappable.MappablesFromInterface(collection.GetList())})
	require.Nil(t, err)
	type fields struct {
		Success bool
		Error   error
		List    []helpers.Mappable
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{

		{"+ve", fields{Success: true, Error: nil, List: collection.GetList()}, encodedByte, false},
		{"-ve with error", fields{Success: false, Error: _err, List: collection.GetList()}, encodedByteWithError, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryResponse := &QueryResponse{
				Success: tt.fields.Success,
				Error:   tt.fields.Error.Error(),
				List:    mappable.MappablesFromInterface(tt.fields.List),
			}
			got, err := queryResponse.Encode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResponse_GetError(t *testing.T) {
	type fields struct {
		Success bool
		Error   error
		List    []helpers.Mappable
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		{"+ve", fields{Success: true, Error: nil}, false},
		{"-ve", fields{Success: true, Error: errorConstants.IncorrectFormat}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryResponse := &QueryResponse{
				Success: tt.fields.Success,
				Error:   tt.fields.Error.Error(),
				List:    mappable.MappablesFromInterface(tt.fields.List),
			}
			if err := queryResponse.GetError(); (err != nil) != tt.wantErr {
				t.Errorf("GetError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_queryResponse_IsSuccessful(t *testing.T) {
	type fields struct {
		Success bool
		Error   error
		List    []helpers.Mappable
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{

		{"+ve", fields{Success: true, Error: nil}, true},
		{"+ve", fields{Success: false, Error: nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryResponse := &QueryResponse{
				Success: tt.fields.Success,
				Error:   tt.fields.Error.Error(),
				List:    mappable.MappablesFromInterface(tt.fields.List),
			}
			if got := queryResponse.IsSuccessful(); got != tt.want {
				t.Errorf("IsSuccessful() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_responsePrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.QueryResponse
	}{

		{"+ve", &QueryResponse{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := responsePrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("responsePrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
