// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package testutil

import (
	"github.com/cometbft/cometbft/crypto/ed25519"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/ids"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/AssetMantle/schema/qualified"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

// TestAddress generates a random AccAddress for testing.
func TestAddress() sdkTypes.AccAddress {
	return sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
}

// TestImmutables returns a standard test Immutables with one string property.
func TestImmutables() qualified.Immutables {
	return baseQualified.NewImmutables(baseLists.NewPropertyList(
		baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")),
	))
}

// TestMutables returns a standard test Mutables with one string property.
func TestMutables() qualified.Mutables {
	return baseQualified.NewMutables(baseLists.NewPropertyList(
		baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData")),
	))
}

// TestClassificationID returns a classification ID built from TestImmutables and TestMutables.
func TestClassificationID() ids.ClassificationID {
	return baseIDs.NewClassificationID(TestImmutables(), TestMutables())
}

// TestIdentityID returns an identity ID built from TestClassificationID and TestImmutables.
func TestIdentityID() ids.IdentityID {
	return baseIDs.NewIdentityID(TestClassificationID(), TestImmutables())
}
