/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/stretchr/testify/require"
)

func Test_Rest_Utils(t *testing.T) {

	value, err := parseGasAdjustment("")
	require.Equal(t, flags.DefaultGasAdjustment, value)
	require.Equal(t, nil, err)

	value2, error2 := parseGasAdjustment("test")
	require.Equal(t, float64(0), value2)
	require.NotNil(t, error2)

	value3, error3 := parseGasAdjustment("0.3")
	require.Equal(t, 0.3, value3)
	require.Equal(t, nil, error3)

	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	gas := uint64(123)
	response, err := simulationResponse(Codec, gas)
	gasEst := rest.GasEstimateResponse{GasEstimate: gas}
	resp, _ := Codec.MarshalJSON(gasEst)
	require.Equal(t, resp, response)
	require.Equal(t, nil, err)

}
