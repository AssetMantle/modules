package issue

import (
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/conform"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
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
	IdentitiesKeeper helpers.TransactionKeeper
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

	scrub.AuxiliaryMock.InitializeKeeper(mapper.Mapper, parameters.Prototype)
	conform.AuxiliaryMock.InitializeKeeper(mapper.Mapper, parameters.Prototype)
	keepers := TestKeepers{
		IdentitiesKeeper: initializeTransactionKeeper(mapper.Mapper, parameters.Prototype,
			[]interface{}{scrub.AuxiliaryMock,
				conform.AuxiliaryMock}),
	}

	return ctx, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	ctx, keepers := CreateTestInput(t)

	immutableMetaTraits, Error := base.ReadMetaProperties("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, Error)
	immutableTraits, Error := base.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, Error)
	mutableMetaTraits, Error := base.ReadMetaProperties("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, Error)
	mutableTraits, Error := base.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, Error)
	scrubMockErrorTraits, Error := base.ReadMetaProperties("scrubError:S|mockError")
	require.Equal(t, nil, Error)
	conformMockErrorTraits, Error := base.ReadMetaProperties("conformError:S|mockError")
	require.Equal(t, nil, Error)
	nubImmutables, Error := base.ReadMetaProperties("nubID:I|nubID")
	require.Equal(t, nil, Error)
	emptyProperties := base.NewProperties()
	require.Equal(t, nil, Error)
	defaultAddr := sdkTypes.AccAddress("addr")
	defaultIdentityID := mapper.NewIdentityID(base.NewID("test.cGn3HMW8M3t5gMDv-wXa9sseHnA="), base.NewID("d0Jhri_bOd3EEPXpyPUpNpGiQ1U="))
	mapper.NewIdentities(mapper.Mapper, ctx).Add(mapper.NewIdentity(defaultIdentityID, []sdkTypes.AccAddress{defaultAddr},
		[]sdkTypes.AccAddress{}, base.NewImmutables(base.NewProperties()), base.NewMutables(base.NewProperties())))
	defaultClassificationID := base.NewID("test.cGn3HMW8M3t5gMDv-wXa9sseHnA=")

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.IdentitiesKeeper.Transact(ctx, newMessage(defaultAddr, defaultAddr, defaultIdentityID, defaultClassificationID,
			immutableMetaTraits, immutableTraits, mutableMetaTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-DuplicateIdentity", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.EntityAlreadyExists)
		if got := keepers.IdentitiesKeeper.Transact(ctx, newMessage(defaultAddr, defaultAddr, defaultIdentityID, defaultClassificationID,
			nubImmutables, emptyProperties, mutableMetaTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-ImmutableScrub Error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.IdentitiesKeeper.Transact(ctx, newMessage(defaultAddr, defaultAddr, defaultIdentityID, defaultClassificationID,
			scrubMockErrorTraits, immutableTraits, mutableMetaTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-MutableScrubError", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.IdentitiesKeeper.Transact(ctx, newMessage(defaultAddr, defaultAddr, defaultIdentityID, base.NewID("newClassificationID"),
			immutableMetaTraits, immutableTraits, scrubMockErrorTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-ConformError", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.IdentitiesKeeper.Transact(ctx, newMessage(defaultAddr, defaultAddr, defaultIdentityID, base.NewID("newClassificationID"),
			immutableMetaTraits, immutableTraits, conformMockErrorTraits, mutableTraits)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
