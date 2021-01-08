/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package rest

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func Test_Rest_Utils(t *testing.T) {

	value, status, Error := ParseFloat64OrReturnBadRequest("", 0.4)
	require.Equal(t, 0.4, value)
	require.Equal(t, http.StatusAccepted, status)
	require.Equal(t, nil, Error)

	value2, status2, error2 := ParseFloat64OrReturnBadRequest("test", 0.5)
	require.Equal(t, float64(0), value2)
	require.Equal(t, http.StatusBadRequest, status2)
	require.NotNil(t, error2)

	value3, status3, error3 := ParseFloat64OrReturnBadRequest("0.3", 0.4)
	require.Equal(t, 0.3, value3)
	require.Equal(t, http.StatusAccepted, status3)
	require.Equal(t, nil, error3)

	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	gas := uint64(123)
	response, Error := SimulationResponse(Codec, gas)
	gasEst := rest.GasEstimateResponse{GasEstimate: gas}
	resp, _ := Codec.MarshalJSON(gasEst)
	require.Equal(t, resp, response)
	require.Equal(t, nil, Error)

}
