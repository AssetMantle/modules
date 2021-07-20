/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/pkg/errors"
)

func parseGasAdjustment(s string) (float64, error) {
	if len(s) == 0 {
		return flags.DefaultGasAdjustment, nil
	}

	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return n, err
	}

	return n, nil
}

func simulationResponse(cdc *codec.Codec, gas uint64) ([]byte, error) {
	gasEst := rest.GasEstimateResponse{GasEstimate: gas}
	resp, err := cdc.MarshalJSON(gasEst)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return resp, nil
}
