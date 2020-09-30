/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	"testing"
)

func CreateTestInput(t *testing.T) sdkTypes.Context {

	keyIdentity := Mapper.GetKVStoreKey()

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyIdentity, sdkTypes.StoreTypeIAVL, db)
	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	ctx := sdkTypes.NewContext(ms, abci.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	return ctx
}

func Test_transactionKeeper_Transact(t *testing.T) {
	ctx := CreateTestInput(t)
	idString := "classification|hashID"
	idByteString := "classificationhashID"

	addr := sdkTypes.AccAddress("addr")

	//Test Identity ID
	id := NewIdentityID(base.NewID("classification"), base.NewID("hashID"))
	require.Equal(t, idString, id.String())
	require.Equal(t, []byte(idByteString), id.Bytes())
	require.Equal(t, false, id.Equal(base.NewID("")))

	//Test Identity
	testIdentity := NewIdentity(id, []sdkTypes.AccAddress{addr},
		[]sdkTypes.AccAddress{}, base.NewImmutables(base.NewProperties()), base.NewMutables(base.NewProperties()))
	testProperty, Error := base.ReadMetaProperty("test:S|test")
	require.Equal(t, nil, Error)
	identityMutated := NewIdentity(id, []sdkTypes.AccAddress{addr},
		[]sdkTypes.AccAddress{}, base.NewImmutables(base.NewProperties(testProperty)), base.NewMutables(base.NewProperties()))

	require.Equal(t, idString, testIdentity.GetID().String())
	require.Equal(t, true, testIdentity.IsProvisioned(addr))
	require.Equal(t, false, testIdentity.IsProvisioned(sdkTypes.AccAddress("addr1")))
	require.Equal(t, false, testIdentity.IsUnprovisioned(sdkTypes.AccAddress("addr")))
	require.Equal(t, base.NewMutables(base.NewProperties()), testIdentity.GetMutables())
	require.Equal(t, base.NewImmutables(base.NewProperties()), testIdentity.GetImmutables())
	require.Equal(t, []sdkTypes.AccAddress{addr}, testIdentity.GetProvisionedAddressList())
	require.Equal(t, []sdkTypes.AccAddress{}, testIdentity.GetUnprovisionedAddressList())
	testIdentity = testIdentity.ProvisionAddress(sdkTypes.AccAddress("addr1"))
	require.Equal(t, true, testIdentity.IsProvisioned(sdkTypes.AccAddress("addr1")))
	testIdentity = testIdentity.UnprovisionAddress(sdkTypes.AccAddress("addr1"))
	require.Equal(t, true, testIdentity.IsUnprovisioned(sdkTypes.AccAddress("addr1")))
	testIdentity = testIdentity.UnprovisionAddress(sdkTypes.AccAddress("addr1"))
	require.Equal(t, true, testIdentity.IsUnprovisioned(sdkTypes.AccAddress("addr1")))

	//Test Identities
	testIdentities := NewIdentities(Mapper, ctx)
	require.Equal(t, "|", testIdentities.GetID().String())
	require.Equal(t, nil, testIdentities.Get(base.NewID("")))

	testIdentities = testIdentities.Add(testIdentity)
	require.Equal(t, testIdentity, testIdentities.Get(id))
	require.Equal(t, NewIdentityID(base.NewID(""), base.NewID("")), testIdentities.GetID())
	require.Equal(t, []mappables.InterIdentity{testIdentity}, testIdentities.GetList())
	require.Equal(t, base.NewID("someID"), testIdentities.Fetch(base.NewID("someID")).GetID())
	require.Equal(t, id, testIdentities.Fetch(id).GetID())
	require.Equal(t, id.String(), testIdentities.Fetch(base.NewID(idString)).GetID().String())

	testIdentities = testIdentities.Mutate(identityMutated)
	require.Equal(t, identityMutated.GetImmutables().Get().GetList(), testIdentities.Get(id).GetImmutables().Get().GetList())
	require.Equal(t, NewIdentities(Mapper, ctx).GetList(), testIdentities.Remove(identityMutated).GetList())

}
