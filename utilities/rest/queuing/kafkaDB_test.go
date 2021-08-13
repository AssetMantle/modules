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
	"github.com/persistenceOne/persistenceSDK/utilities/random"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"
	"testing"
)

func Test_Kafka_DB(t *testing.T) {
	require.Panics(t, func() {
		var Codec = codec.NewLegacyAmino()
		schema.RegisterLegacyAminoCodec(Codec)
		sdkTypes.RegisterLegacyAminoCodec(Codec)
		codec.RegisterCrypto(Codec)
		codec.RegisterEvidences(Codec)
		vesting.RegisterCodec(Codec)
		Codec.Seal()
		ticketID := TicketID(random.GenerateID("name"))
		kafkaDB, _ := dbm.NewGoLevelDB("KafkaDB", defaultCLIHome)
		setTicketIDtoDB(ticketID, kafkaDB, Codec, []byte{})
		addResponseToDB(ticketID, base.NewID("").Bytes(), kafkaDB, Codec)
		require.Equal(t, base.NewID("").Bytes(), getResponseFromDB(ticketID, kafkaDB, Codec))
	})
}
