package define

import (
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/define"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/super"
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

	scrub.AuxiliaryMock.InitializeKeeper(mapper.Mapper)
	define.AuxiliaryMock.InitializeKeeper(mapper.Mapper)
	super.AuxiliaryMock.InitializeKeeper(mapper.Mapper)
	keepers := TestKeepers{
		IdentitiesKeeper: initializeTransactionKeeper(mapper.Mapper,
			[]interface{}{scrub.AuxiliaryMock,
				define.AuxiliaryMock, super.AuxiliaryMock}),
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

	defaultIdentityID := mapper.NewIdentityID(base.NewID("test.cGn3HMW8M3t5gMDv-wXa9sseHnA="), base.NewID("d0Jhri_bOd3EEPXpyPUpNpGiQ1U="))
	mapper.NewIdentities(mapper.Mapper, ctx).Add(mapper.NewIdentity(defaultIdentityID, []sdkTypes.AccAddress{sdkTypes.AccAddress("addr")},
		[]sdkTypes.AccAddress{}, base.NewImmutables(base.NewProperties()), base.NewMutables(base.NewProperties())))

	type fields struct {
		keeper helpers.TransactionKeeper
	}

	type args struct {
		context sdkTypes.Context
		msg     sdkTypes.Msg
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.TransactionResponse
	}{
		{
			name:   "Expected Case",
			fields: fields{keepers.IdentitiesKeeper},
			args: args{
				context: ctx,
				msg: message{
					From:                sdkTypes.AccAddress("addr"),
					FromID:              defaultIdentityID,
					ImmutableMetaTraits: immutableMetaTraits,
					ImmutableTraits:     immutableTraits,
					MutableMetaTraits:   mutableMetaTraits,
					MutableTraits:       mutableTraits,
				},
			},
			want: transactionResponse{
				Success: true,
				Error:   nil,
			},
		},
		{
			name:   "Identity nil case",
			fields: fields{keepers.IdentitiesKeeper},
			args: args{
				context: ctx,
				msg: message{
					From:                sdkTypes.AccAddress("addr"),
					FromID:              base.NewID(""),
					ImmutableMetaTraits: immutableMetaTraits,
					ImmutableTraits:     immutableTraits,
					MutableMetaTraits:   mutableMetaTraits,
					MutableTraits:       mutableTraits,
				},
			},
			want: transactionResponse{
				Success: false,
				Error:   errors.EntityNotFound,
			},
		},
		{
			name:   "Identity unprovisioned address case",
			fields: fields{keepers.IdentitiesKeeper},
			args: args{
				context: ctx,
				msg: message{
					From:                sdkTypes.AccAddress("addr1"),
					FromID:              defaultIdentityID,
					ImmutableMetaTraits: immutableMetaTraits,
					ImmutableTraits:     immutableTraits,
					MutableMetaTraits:   mutableMetaTraits,
					MutableTraits:       mutableTraits,
				},
			},
			want: transactionResponse{
				Success: false,
				Error:   errors.NotAuthorized,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.keeper.Transact(tt.args.context, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
