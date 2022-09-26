// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"github.com/AssetMantle/modules/modules/identities/internal/common"
	"github.com/AssetMantle/modules/modules/identities/internal/mapper"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
	"reflect"
	"testing"
)

func CreateTestInputContext(t *testing.T) sdkTypes.Context {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

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

	context := sdkTypes.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	return context
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

		{"+ve", args{collection: collection, error: nil}, queryResponse{Success: true, Error: nil}},
		{"-ve with error", args{collection: collection, error: constants.IncorrectFormat}, queryResponse{Success: false, Error: constants.IncorrectFormat}},
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
		Error   error
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

		{"+ve", fields{Success: true, Error: nil}, args{bytes: encodedResponse}, testQueryResponse, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryResponse := queryResponse{
				Success: tt.fields.Success,
				Error:   tt.fields.Error,
				List:    tt.fields.List,
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
	encodedByte, err := common.Codec.MarshalJSON(queryResponse{Success: true, Error: nil, List: collection.GetList()})
	encodedByteWithError, _err := common.Codec.MarshalJSON(queryResponse{Success: false, Error: constants.IncorrectFormat, List: collection.GetList()})
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
			queryResponse := queryResponse{
				Success: tt.fields.Success,
				Error:   tt.fields.Error,
				List:    tt.fields.List,
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
		{"-ve", fields{Success: true, Error: constants.IncorrectFormat}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryResponse := queryResponse{
				Success: tt.fields.Success,
				Error:   tt.fields.Error,
				List:    tt.fields.List,
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
			queryResponse := queryResponse{
				Success: tt.fields.Success,
				Error:   tt.fields.Error,
				List:    tt.fields.List,
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

		{"+ve", queryResponse{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := responsePrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("responsePrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
