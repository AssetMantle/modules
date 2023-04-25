// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package asset

import (
	"testing"

	"github.com/AssetMantle/modules/x/assets/common"
	"github.com/AssetMantle/modules/x/assets/mapper"

	schema "github.com/AssetMantle/schema/go"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
)

func CreateTestInput(t *testing.T) sdkTypes.Context {
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

	return context
}

func Test_Asset_Response(t *testing.T) {
	context := CreateTestInput(t)
	collection := mapper.Prototype().NewCollection(sdkTypes.WrapSDKContext(context))

	testQueryResponse := newQueryResponse(collection, nil)
	testQueryResponseWithError := newQueryResponse(collection, errorConstants.IncorrectFormat)

	require.Equal(t, true, testQueryResponse.IsSuccessful())
	require.Equal(t, false, testQueryResponseWithError.IsSuccessful())
	require.Equal(t, nil, testQueryResponse.GetError())
	require.Equal(t, errorConstants.IncorrectFormat, testQueryResponseWithError.GetError())

	encodedResponse, _ := testQueryResponse.Encode()
	bytes, _ := common.LegacyAmino.MarshalJSON(testQueryResponse)
	require.Equal(t, bytes, encodedResponse)

	decodedResponse, _ := (&QueryResponse{}).Decode(bytes)
	require.Equal(t, testQueryResponse, decodedResponse)

	decodedResponse2, _ := (&QueryResponse{}).Decode([]byte{})
	require.Equal(t, nil, decodedResponse2)

	require.Equal(t, &QueryResponse{}, responsePrototype())
}
