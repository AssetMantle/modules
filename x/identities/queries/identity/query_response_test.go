// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"context"
	"reflect"
	"testing"

	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/identities/mappable"
	"github.com/AssetMantle/modules/x/identities/mapper"
)

func CreateTestInputContext(t *testing.T) context.Context {
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

		{"+ve", args{collection: collection, error: nil}, &QueryResponse{}},
		{"-ve with error", args{collection: collection, error: errorConstants.IncorrectFormat}, &QueryResponse{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newQueryResponse(tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newQueryResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResponse_Decode(t *testing.T) {
	context := CreateTestInputContext(t)
	collection := mapper.Prototype().NewCollection(context)
	testQueryResponse := newQueryResponse(collection)
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
				List: mappable.MappablesFromInterface(tt.fields.List),
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
	encodedByte, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(&QueryResponse{List: mappable.MappablesFromInterface(collection.Get())})
	encodedByteWithError, _err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(&QueryResponse{List: mappable.MappablesFromInterface(collection.Get())})
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

		{"+ve", fields{Success: true, Error: nil, List: collection.Get()}, encodedByte, false},
		{"-ve with error", fields{Error: _err, List: collection.Get()}, encodedByteWithError, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryResponse := &QueryResponse{
				List: mappable.MappablesFromInterface(tt.fields.List),
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
