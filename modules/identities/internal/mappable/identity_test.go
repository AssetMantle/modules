package mappable

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Identity_Methods(t *testing.T) {

	classificationID := base.NewID("classificationID")
	hashID := base.NewID("hashID")
	testIdentityID := key.NewIdentityID(classificationID, hashID)
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
