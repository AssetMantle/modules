// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queuing

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/schema"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/utilities/random"
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
		ticketID := TicketID(random.GenerateUniqueIdentifier("name"))
		kafkaDB, _ := dbm.NewGoLevelDB("KafkaDB", defaultCLIHome)
		setTicketIDtoDB(ticketID, kafkaDB, Codec, []byte{})
		addResponseToDB(ticketID, baseIDs.NewID("").Bytes(), kafkaDB, Codec)
		require.Equal(t, baseIDs.NewID("").Bytes(), getResponseFromDB(ticketID, kafkaDB, Codec))
	})
}
