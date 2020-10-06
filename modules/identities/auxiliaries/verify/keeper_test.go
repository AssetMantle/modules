package verify

import (
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
	"reflect"
	"testing"
)

type TestKeepers struct {
	IdentitiesKeeper helpers.AuxiliaryKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {

	keyIdentity := mapper.Mapper.GetKVStoreKey()

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyIdentity, sdkTypes.StoreTypeIAVL, db)
	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	ctx := sdkTypes.NewContext(ms, abci.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		IdentitiesKeeper: initializeAuxiliaryKeeper(mapper.Mapper, parameters.Prototype, []interface{}{}),
	}

	return ctx, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	ctx, keepers := CreateTestInput(t)
	defaultAddr := sdkTypes.AccAddress("addr")
	unprovisionedAddr := sdkTypes.AccAddress("unProvisionedAddr")
	defaultClassificationID := base.NewID("test.cGn3HMW8M3t5gMDv-wXa9sseHnA=")
	defaultIdentityID := mapper.NewIdentityID(defaultClassificationID, base.NewID("d0Jhri_bOd3EEPXpyPUpNpGiQ1U="))
	mapper.NewIdentities(mapper.Mapper, ctx).Add(mapper.NewIdentity(defaultIdentityID, []sdkTypes.AccAddress{defaultAddr},
		[]sdkTypes.AccAddress{unprovisionedAddr}, base.NewImmutables(base.NewProperties()), base.NewMutables(base.NewProperties())))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newAuxiliaryResponse(nil)
		if got := keepers.IdentitiesKeeper.Help(ctx, NewAuxiliaryRequest(defaultAddr, defaultIdentityID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Nil Identity", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(errors.EntityNotFound)
		if got := keepers.IdentitiesKeeper.Help(ctx, NewAuxiliaryRequest(defaultAddr, base.NewID("id"))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Unprovisioned Address", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(errors.NotAuthorized)
		if got := keepers.IdentitiesKeeper.Help(ctx, NewAuxiliaryRequest(unprovisionedAddr, defaultIdentityID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
