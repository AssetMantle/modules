// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queuing

import (
	"github.com/AssetMantle/modules/utilities/rest"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers/base"
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

	gas := uint64(123)
	response, err := simulationResponse(base.CodecPrototype().GetLegacyAmino(), gas)
	gasEst := rest.GasEstimateResponse{GasEstimate: gas}
	resp, _ := base.CodecPrototype().GetLegacyAmino().MarshalJSON(gasEst)
	require.Equal(t, resp, response)
	require.Equal(t, nil, err)

}
