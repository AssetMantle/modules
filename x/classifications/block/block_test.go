// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package block

import (
	"context"
	storeTypes "cosmossdk.io/store/types"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base/testutil"
	"github.com/AssetMantle/modules/x/classifications/mapper"
	"github.com/AssetMantle/modules/x/classifications/parameters"
)

func createTestInput(t *testing.T) context.Context {
	storeKey := storeTypes.NewKVStoreKey("test")
	ctx := testutil.NewTestContext(t, storeKey)
	return sdkTypes.WrapSDKContext(ctx)
}

func Test_Block_Methods(t *testing.T) {
	block := Prototype()
	block.Initialize(mapper.Prototype(), parameters.Prototype(), []helpers.Auxiliary{})
	context := createTestInput(t)
	block.Begin(context)
	block.End(context)
}
