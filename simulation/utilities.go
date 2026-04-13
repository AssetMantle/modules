// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulation

import (
	"context"
	"cosmossdk.io/math"
	"fmt"
	goGoProto "github.com/cosmos/gogoproto/proto"
	"math/rand"
	"strings"

	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/qualified"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"

	"github.com/AssetMantle/modules/helpers"
	baseSimulation "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/x/classifications/parameters/bond_rate"
)

var (
	Immutables qualified.Immutables = &baseQualified.Immutables{}
	Mutables   qualified.Mutables   = &baseQualified.Mutables{}

	// SimTxConfig and SimAccountKeeper are set by module WeightedOperations
	// from SimulationState and the node's application. Used by DeliverSimTx
	// to create properly signed transactions delivered through the full ABCI flow.
	SimTxConfig     client.TxConfig
	SimAccountKeeper AccountKeeper
)

// AccountKeeper defines the interface for looking up accounts during simulation.
type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdkTypes.AccAddress) sdkTypes.AccountI
}

func RandomBool(r *rand.Rand) bool {
	return r.Intn(2) == 0
}

func GenerateRandomAddresses(r *rand.Rand) []sdkTypes.AccAddress {
	randomAccounts := simulationTypes.RandomAccounts(r, r.Intn(99))
	addresses := make([]sdkTypes.AccAddress, len(randomAccounts))

	for i, account := range randomAccounts {
		addresses[i] = account.Address
	}

	return addresses
}

func generateGenesisProperties(r *rand.Rand) {
	Immutables = baseQualified.NewImmutables(baseSimulation.GenerateRandomMetaPropertyListWithoutData(r))
	Mutables = baseQualified.NewMutables(baseSimulation.GenerateRandomPropertyList(r))
}

func GetGenesisProperties(r *rand.Rand) (qualified.Immutables, qualified.Mutables) {
	if Immutables.(*baseQualified.Immutables).PropertyList == nil {
		generateGenesisProperties(r)
	}
	return Immutables, Mutables
}

func CalculateBondAmount(immutables qualified.Immutables, mutables qualified.Mutables) data.NumberData {
	totalWeight := math.ZeroInt()
	for _, property := range append(immutables.GetImmutablePropertyList().Get(), mutables.GetMutablePropertyList().Get()...) {
		if inner := property.Get(); inner != nil {
			totalWeight = totalWeight.Add(inner.GetBondWeight())
		}
	}

	bondRateData := bond_rate.Parameter.GetMetaProperty().GetData()
	if bondRateData == nil || bondRateData.Get() == nil {
		return baseData.NewNumberData(totalWeight)
	}
	if numData, ok := bondRateData.Get().(data.NumberData); ok {
		return baseData.NewNumberData(numData.Get().Mul(totalWeight))
	}
	return baseData.NewNumberData(totalWeight)
}

func ExecuteMessage(context sdkTypes.Context, module helpers.Module, message helpers.Message) (*sdkTypes.Result, error) {
	if module == nil || message == nil {
		return nil, fmt.Errorf("nil module or message")
	}
	transactions := module.GetTransactions()
	if transactions == nil {
		return nil, fmt.Errorf("module %s has no transactions", module.Name())
	}

	// Derive service path from proto message name.
	// Proto name: "AssetMantle.modules.x.assets.transactions.burn.Message"
	// Service path: "/assets/burn"
	msgName := goGoProto.MessageName(message)
	parts := strings.Split(msgName, ".")
	if len(parts) >= 6 {
		servicePath := "/" + parts[3] + "/" + parts[5]
		if tx := transactions.GetTransaction(servicePath); tx != nil {
			return tx.HandleMessage(sdkTypes.WrapSDKContext(context), message)
		}
	}
	return nil, fmt.Errorf("no matching transaction for message %s in module %s", msgName, module.Name())
}

// DeliverSimTx creates a properly signed transaction and delivers it through the
// full ABCI flow (AnteHandler, authentication, etc.). This is the correct way to
// execute messages during simulation — unlike ExecuteMessage which bypasses the
// transaction pipeline. Returns OperationMsg with success/failure indication.
func DeliverSimTx(r *rand.Rand, app *baseapp.BaseApp, ctx sdkTypes.Context, account simulationTypes.Account, msg sdkTypes.Msg, moduleName string) (simulationTypes.OperationMsg, error) {
	if SimTxConfig == nil || SimAccountKeeper == nil {
		return simulationTypes.NoOpMsg(moduleName, sdkTypes.MsgTypeURL(msg), "sim infrastructure not set"), fmt.Errorf("SimTxConfig or SimAccountKeeper not initialized")
	}

	acc := SimAccountKeeper.GetAccount(ctx, account.Address)
	if acc == nil {
		return simulationTypes.NoOpMsg(moduleName, sdkTypes.MsgTypeURL(msg), "account not found"), fmt.Errorf("account %s not found", account.Address)
	}

	tx, err := simtestutil.GenSignedMockTx(
		r,
		SimTxConfig,
		[]sdkTypes.Msg{msg},
		sdkTypes.NewCoins(sdkTypes.NewInt64Coin(sdkTypes.DefaultBondDenom, 0)),
		simtestutil.DefaultGenTxGas,
		ctx.ChainID(),
		[]uint64{acc.GetAccountNumber()},
		[]uint64{acc.GetSequence()},
		account.PrivKey,
	)
	if err != nil {
		return simulationTypes.NoOpMsg(moduleName, sdkTypes.MsgTypeURL(msg), "unable to generate mock tx"), err
	}

	_, _, err = app.SimDeliver(SimTxConfig.TxEncoder(), tx)
	if err != nil {
		return simulationTypes.NoOpMsg(moduleName, sdkTypes.MsgTypeURL(msg), err.Error()), nil
	}

	return simulationTypes.NewOperationMsg(msg, true, ""), nil
}

// SafeOperation wraps a simulation operation to catch panics from nil data in
// simulated databases. Early blocks have empty databases, causing operations to
// panic on nil type assertions. This wrapper converts panics to no-op results.
func SafeOperation(op simulationTypes.Operation) simulationTypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdkTypes.Context, accs []simulationTypes.Account, chainID string) (opMsg simulationTypes.OperationMsg, futures []simulationTypes.FutureOperation, err error) {
		defer func() {
			if rec := recover(); rec != nil {
				opMsg = simulationTypes.NoOpMsg("recovered", "panic", fmt.Sprintf("%v", rec))
				err = nil
			}
		}()
		return op(r, app, ctx, accs, chainID)
	}
}
