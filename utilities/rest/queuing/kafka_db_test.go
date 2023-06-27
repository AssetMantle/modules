// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queuing

import (
	"github.com/AssetMantle/modules/helpers/base"
	"testing"

	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/utilities/random"
)

func Test_Kafka_DB(t *testing.T) {
	require.Panics(t, func() {
		var legacyAmino = base.CodecPrototype().GetLegacyAmino()
		ticketID := TicketID(random.GenerateUniqueIdentifier("name"))
		kafkaDB, _ := dbm.NewGoLevelDB("KafkaDB", defaultCLIHome)
		setTicketIDtoDB(ticketID, kafkaDB, legacyAmino, []byte{})
		addResponseToDB(ticketID, baseIDs.NewStringID("").Bytes(), kafkaDB, legacyAmino)
		require.Equal(t, baseIDs.NewStringID("").Bytes(), getResponseFromDB(ticketID, kafkaDB, legacyAmino))
	})
}
