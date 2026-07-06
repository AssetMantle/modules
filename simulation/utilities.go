// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulation

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"cosmossdk.io/math"
	goGoProto "github.com/cosmos/gogoproto/proto"

	"github.com/AssetMantle/schema/properties"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/AssetMantle/modules/helpers"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	assetsDB "github.com/AssetMantle/modules/simulation/simulated_database/assets"
	identitiesDB "github.com/AssetMantle/modules/simulation/simulated_database/identities"
	ordersDB "github.com/AssetMantle/modules/simulation/simulated_database/orders"
)

// MaxBondRate is the inclusive upper bound of the classifications bond rate
// randomized in simulation genesis. Operations that must reveal a bond amount
// use it to bound rate*totalWeight from above without knowing the drawn rate.
const MaxBondRate = 98

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

// ErrSimDBMiss is returned when a simulated database lookup finds no data for an account.
var ErrSimDBMiss = fmt.Errorf("simulated database miss")

// LookupIdentityID returns the first identity ID string for the given address, or error if none exists.
func LookupIdentityID(address string) (string, error) {
	idMap := identitiesDB.GetIDData(address)
	if idMap == nil {
		return "", ErrSimDBMiss
	}
	for _, id := range idMap {
		return id, nil
	}
	return "", ErrSimDBMiss
}

// LookupAssetID returns the first asset ID string for the given address, or error if none exists.
func LookupAssetID(address string) (string, error) {
	assetMap := assetsDB.GetAssetData(address)
	if assetMap == nil {
		return "", ErrSimDBMiss
	}
	for _, id := range assetMap {
		return id, nil
	}
	return "", ErrSimDBMiss
}

// LookupClassificationID returns the first classification ID string for the given address's assets, or error if none exists.
func LookupClassificationID(address string) (string, error) {
	assetMap := assetsDB.GetAssetData(address)
	if assetMap == nil {
		return "", ErrSimDBMiss
	}
	for class := range assetMap {
		return class, nil
	}
	return "", ErrSimDBMiss
}

// LookupOrderClassificationID returns the first order classification ID string for the given address, or error if none exists.
func LookupOrderClassificationID(address string) (string, error) {
	orderMap := ordersDB.GetOrderData(address)
	if orderMap == nil {
		return "", ErrSimDBMiss
	}
	for class := range orderMap {
		return class, nil
	}
	return "", ErrSimDBMiss
}

// IsBusinessRejection reports whether an error from executing a self-constructed
// simulation message is a legitimate business-rule rejection under randomized
// parameters and state (disabled gate, exhausted balance, random collision,
// property cap, simulated database miss) rather than a delivery bug.
func IsBusinessRejection(err error) bool {
	return errors.Is(err, ErrSimDBMiss) ||
		errors.Is(err, errorConstants.NotAuthorized) ||
		errors.Is(err, errorConstants.InsufficientBalance) ||
		errors.Is(err, errorConstants.EntityAlreadyExists) ||
		errors.Is(err, errorConstants.EntityNotFound) ||
		errors.Is(err, sdkErrors.ErrInsufficientFunds) ||
		(errors.Is(err, errorConstants.InvalidRequest) && strings.Contains(err.Error(), "property count"))
}

// RejectionOrError converts a failed self-constructed message execution into the
// standard operation result: a NoOp with nil error for legitimate business
// rejections, a NoOp carrying the error (failing the simulation) for bugs.
func RejectionOrError(moduleName, route string, err error) (simulationTypes.OperationMsg, []simulationTypes.FutureOperation, error) {
	if IsBusinessRejection(err) {
		return simulationTypes.NoOpMsg(moduleName, route, err.Error()), nil, nil
	}
	return simulationTypes.NoOpMsg(moduleName, route, err.Error()), nil, err
}

// SumBondWeights sums the bond weights of the given property snapshots plus the
// given extra properties, mirroring the classifications define auxiliary's
// weighing of a classification's immutable and mutable property lists.
func SumBondWeights(snaps [][]properties.AnyProperty, extras ...properties.Property) math.Int {
	totalWeight := math.ZeroInt()
	for _, snap := range snaps {
		for _, anyProperty := range snap {
			totalWeight = totalWeight.Add(anyProperty.Get().GetBondWeight())
		}
	}
	for _, property := range extras {
		totalWeight = totalWeight.Add(property.GetBondWeight())
	}
	return totalWeight
}

// SafeOperation wraps a simulation operation so that a panic surfaces as a
// simulation failure carrying the operation route and panic value, instead of
// crashing the process. It never swallows the panic into a passing no-op.
func SafeOperation(route string, op simulationTypes.Operation) simulationTypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdkTypes.Context, accs []simulationTypes.Account, chainID string) (opMsg simulationTypes.OperationMsg, futures []simulationTypes.FutureOperation, err error) {
		defer func() {
			if rec := recover(); rec != nil {
				opMsg = simulationTypes.NoOpMsg("recovered", "panic", fmt.Sprintf("%s: %v", route, rec))
				err = fmt.Errorf("panic in simulation operation %s: %v", route, rec)
			}
		}()
		return op(r, app, ctx, accs, chainID)
	}
}
