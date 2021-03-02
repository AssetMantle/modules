/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"testing"

	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_Asset_Methods(t *testing.T) {
	classificationID := base.NewID("classificationID")
	immutables := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("ImmutableData"))))
	mutables := base.NewProperties(base.NewProperty(base.NewID("ID2"), base.NewFact(base.NewStringData("MutableData"))))

	assetID := key.NewAssetID(classificationID, immutables)
	testAsset := NewAsset(assetID, immutables, mutables).(asset)

	require.Equal(t, asset{ID: assetID, Immutables: baseTraits.Immutables{Properties: immutables}, Mutables: baseTraits.Mutables{Properties: mutables}}, testAsset)
	require.Equal(t, assetID, testAsset.GetID())
	require.Equal(t, classificationID, testAsset.GetClassificationID())
	require.Equal(t, immutables, testAsset.GetImmutableProperties())
	require.Equal(t, mutables, testAsset.GetMutableProperties())
	data, _ := base.ReadHeightData("")
	require.Equal(t, base.NewProperty(base.NewID(properties.Burn), base.NewFact(data)), testAsset.GetBurn())
	require.Equal(t, base.NewProperty(base.NewID(properties.Burn), base.NewFact(base.NewStringData("BurnImmutableData"))), asset{ID: assetID, Immutables: baseTraits.Immutables{Properties: base.NewProperties(base.NewProperty(base.NewID(properties.Burn), base.NewFact(base.NewStringData("BurnImmutableData"))))}, Mutables: baseTraits.Mutables{Properties: mutables}}.GetBurn())
	require.Equal(t, base.NewProperty(base.NewID(properties.Burn), base.NewFact(base.NewStringData("BurnMutableData"))), asset{ID: assetID, Immutables: baseTraits.Immutables{Properties: immutables}, Mutables: baseTraits.Mutables{Properties: base.NewProperties(base.NewProperty(base.NewID(properties.Burn), base.NewFact(base.NewStringData("BurnMutableData"))))}}.GetBurn())
	require.Equal(t, base.NewProperty(base.NewID(properties.Lock), base.NewFact(data)), testAsset.GetLock())
	require.Equal(t, base.NewProperty(base.NewID(properties.Lock), base.NewFact(base.NewStringData("LockImmutableData"))), asset{ID: assetID, Immutables: baseTraits.Immutables{Properties: base.NewProperties(base.NewProperty(base.NewID(properties.Lock), base.NewFact(base.NewStringData("LockImmutableData"))))}, Mutables: baseTraits.Mutables{Properties: mutables}}.GetLock())
	require.Equal(t, base.NewProperty(base.NewID(properties.Lock), base.NewFact(base.NewStringData("LockMutableData"))), asset{ID: assetID, Immutables: baseTraits.Immutables{Properties: immutables}, Mutables: baseTraits.Mutables{Properties: base.NewProperties(base.NewProperty(base.NewID(properties.Lock), base.NewFact(base.NewStringData("LockMutableData"))))}}.GetLock())
	require.Equal(t, assetID, testAsset.GetKey())

}
