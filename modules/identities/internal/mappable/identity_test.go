/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_Identity_Methods(t *testing.T) {

	classificationID := base.NewID("classificationID")
	immutableProperties, _ := base.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	testIdentityID := key.NewIdentityID(classificationID, base.NewImmutables(immutableProperties))
	provisionedAddressList := []sdkTypes.AccAddress{sdkTypes.AccAddress("provAddr")}
	unProvisionedAddressList := []sdkTypes.AccAddress{sdkTypes.AccAddress("unProvAddr")}
	immutables := base.NewImmutables(base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("ImmutableData")))))
	mutables := base.NewMutables(base.NewProperties(base.NewProperty(base.NewID("ID2"), base.NewFact(base.NewStringData("MutableData")))))

	testIdentity := NewIdentity(testIdentityID, provisionedAddressList, unProvisionedAddressList, immutables, mutables)
	require.Equal(t, testIdentity, identity{ID: testIdentityID, ProvisionedAddressList: provisionedAddressList, UnprovisionedAddressList: unProvisionedAddressList, Immutables: immutables, Mutables: mutables})
	require.Equal(t, testIdentity.(identity).GetID(), testIdentityID)
	require.Equal(t, testIdentity.GetProvisionedAddressList(), provisionedAddressList)
	require.Equal(t, testIdentity.GetUnprovisionedAddressList(), unProvisionedAddressList)
	require.Equal(t, testIdentity.ProvisionAddress(sdkTypes.AccAddress("newAddr")).GetProvisionedAddressList(), append(provisionedAddressList, sdkTypes.AccAddress("newAddr")))
	require.Equal(t, testIdentity.UnprovisionAddress(sdkTypes.AccAddress("provAddr")).GetProvisionedAddressList(), []sdkTypes.AccAddress{})
	require.Equal(t, testIdentity.UnprovisionAddress(sdkTypes.AccAddress("provAddr")).GetUnprovisionedAddressList(), append(unProvisionedAddressList, sdkTypes.AccAddress("provAddr")))
	require.Equal(t, testIdentity.GetImmutables(), immutables)
	require.Equal(t, testIdentity.GetMutables(), mutables)
	require.Equal(t, testIdentity.IsProvisioned(sdkTypes.AccAddress("provAddr")), true)
	require.Equal(t, testIdentity.IsProvisioned(sdkTypes.AccAddress("unProvAddr")), false)
	require.Equal(t, testIdentity.IsUnprovisioned(sdkTypes.AccAddress("unProvAddr")), true)
	require.Equal(t, testIdentity.IsUnprovisioned(sdkTypes.AccAddress("provAddr")), false)
	require.Equal(t, testIdentity.GetKey(), testIdentityID)
}
