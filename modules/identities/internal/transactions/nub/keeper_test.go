package nub

import (
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/define"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
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
	keepers := TestKeepers{
		IdentitiesKeeper: initializeTransactionKeeper(mapper.Mapper,
			[]interface{}{scrub.AuxiliaryMock,
				define.AuxiliaryMock}),
	}

	return ctx, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	ctx, keepers := CreateTestInput(t)

	defaultIdentityID := mapper.NewIdentityID(base.NewID("test.cGn3HMW8M3t5gMDv-wXa9sseHnA="), base.NewID("d0Jhri_bOd3EEPXpyPUpNpGiQ1U="))
	mapper.NewIdentities(mapper.Mapper, ctx).Add(mapper.NewIdentity(defaultIdentityID, []sdkTypes.AccAddress{sdkTypes.AccAddress("addr")},
		[]sdkTypes.AccAddress{}, base.NewImmutables(base.NewProperties()), base.NewMutables(base.NewProperties())))

	response := keepers.IdentitiesKeeper.Transact(ctx, newMessage(sdkTypes.AccAddress("addr"), base.NewID("nubID1")))
	require.Equal(t, true, response.IsSuccessful())
	require.Nil(t, response.GetError())

	response = keepers.IdentitiesKeeper.Transact(ctx, newMessage(sdkTypes.AccAddress("addr"), base.NewID("nubID")))
	//require.Equal(t, true,reflect.DeepEqual(response, transactionResponse{
	//	Success: false,
	//	Error:   errors.EntityAlreadyExists,
	//}))
	require.Equal(t, false, response.IsSuccessful())
	require.Equal(t, response.GetError(), errors.EntityAlreadyExists)

	//type fields struct {
	//	keeper helpers.TransactionKeeper
	//}
	//
	//type args struct {
	//	context sdkTypes.Context
	//	msg     sdkTypes.Msg
	//}
	//
	//tests := []struct {
	//	name   string
	//	fields fields
	//	args   args
	//	want   helpers.TransactionResponse
	//}{
	//	{
	//		name:   "Expected Case",
	//		fields: fields{keepers.IdentitiesKeeper},
	//		args: args{
	//			context: ctx,
	//			msg:     newMessage(sdkTypes.AccAddress("addr"), base.NewID("nubID1")),
	//		},
	//		want: transactionResponse{
	//			Success: true,
	//			Error:   nil,
	//		},
	//	},
	//	{
	//		name:   "Duplicate nub identity addition.",
	//		fields: fields{keepers.IdentitiesKeeper},
	//		args: args{
	//			context: ctx,
	//			msg: message{
	//				From:  sdkTypes.AccAddress("addr"),
	//				NubID: base.NewID("nubID"),
	//			},
	//		},
	//		want: transactionResponse{
	//			Success: false,
	//			Error:   errors.EntityAlreadyExists,
	//		},
	//	},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := tt.fields.keeper.Transact(tt.args.context, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("Transact() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}
