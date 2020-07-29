package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/pkg/errors"
)

func ParseFloat64OrReturnBadRequest(s string, defaultIfEmpty float64) (float64, int, error) {
	if len(s) == 0 {
		return defaultIfEmpty, http.StatusAccepted, nil
	}

	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return n, http.StatusBadRequest, err
	}

	return n, http.StatusAccepted, nil
}

func SimulationResponse(cdc *codec.Codec, gas uint64) ([]byte, error) {
	gasEst := rest.GasEstimateResponse{GasEstimate: gas}
	resp, err := cdc.MarshalJSON(gasEst)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return resp, nil
}

func PostProcessResponse(cliCtx context.CLIContext, response interface{}) ([]byte, error) {
	var output []byte

	if cliCtx.Height < 0 {
		return nil, errors.New("Negative height in response")
	}

	switch response.(type) {
	case []byte:
		output = response.([]byte)

	default:
		var err error
		if cliCtx.Indent {
			output, err = cliCtx.Codec.MarshalJSONIndent(response, "", "  ")
		} else {
			output, err = cliCtx.Codec.MarshalJSON(response)
		}

		if err != nil {
			return nil, errors.New(err.Error())
		}
	}

	if cliCtx.Height > 0 {
		m := make(map[string]interface{})
		err := json.Unmarshal(output, &m)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		m["height"] = cliCtx.Height

		if cliCtx.Indent {
			output, err = json.MarshalIndent(m, "", "  ")
		} else {
			output, err = json.Marshal(m)
		}
		if err != nil {
			return nil, errors.New(err.Error())
		}
	}

	return output, nil
}
