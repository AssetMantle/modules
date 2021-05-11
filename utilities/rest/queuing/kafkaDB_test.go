/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"
	"testing"
)

func Test_Kafka_DB(t *testing.T) {
	require.Panics(t, func() {
		var Codec = codec.New()
		schema.RegisterCodec(Codec)
		sdkTypes.RegisterCodec(Codec)
		codec.RegisterCrypto(Codec)
		codec.RegisterEvidences(Codec)
		vesting.RegisterCodec(Codec)
		Codec.Seal()
		ticketID := TicketIDGenerator("name")
		kafkaDB, _ := dbm.NewGoLevelDB("KafkaDB", DefaultCLIHome)
		SetTicketIDtoDB(ticketID, kafkaDB, Codec, []byte{})
		AddResponseToDB(ticketID, base.NewID("").Bytes(), kafkaDB, Codec)
		require.Equal(t, base.NewID("").Bytes(), GetResponseFromDB(ticketID, kafkaDB, Codec))
	})
}
