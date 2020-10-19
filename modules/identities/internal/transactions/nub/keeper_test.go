/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package nub

import (
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/classifications/auxiliaries/define"
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

	scrub.AuxiliaryMock.Initialize(mapper.Mapper, parameters.Prototype)
	define.AuxiliaryMock.Initialize(mapper.Mapper, parameters.Prototype)
	keepers := TestKeepers{
		IdentitiesKeeper: initializeTransactionKeeper(mapper.Mapper, parameters.Prototype,
			[]interface{}{scrub.AuxiliaryMock,
				define.AuxiliaryMock}),
	}

	return ctx, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	ctx, keepers := CreateTestInput(t)

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.IdentitiesKeeper.Transact(ctx, newMessage(sdkTypes.AccAddress("addr"), base.NewID("nubID"))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Duplicate", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.EntityAlreadyExists)
		if got := keepers.IdentitiesKeeper.Transact(ctx, newMessage(sdkTypes.AccAddress("addr"), base.NewID("nubID"))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
