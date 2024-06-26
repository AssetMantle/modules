// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wasm

import (
	"encoding/json"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"strings"

	"github.com/CosmWasm/wasmd/x/wasm"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
)

func CustomEncoder(moduleList ...helpers.Module) wasm.CustomEncoder {
	return func(sender sdkTypes.AccAddress, rawMessage json.RawMessage) ([]sdkTypes.Msg, error) {
		wasmMessage := baseHelpers.WasmMessagePrototype()

		err := json.Unmarshal(rawMessage, &wasmMessage)
		if err != nil {
			return nil, errorConstants.InvalidMessage.Wrapf(err.Error())
		}

		path := strings.Split(wasmMessage.GetType(), "/")

		for _, module := range moduleList {
			if module.Name() == path[0] {
				msg, err := module.DecodeModuleTransactionRequest(path[1], wasmMessage.GetRawMessage())
				if err != nil {
					return nil, err
				}

				return []sdkTypes.Msg{msg}, nil
			}
		}

		return nil, errorConstants.InvalidMessage.Wrapf("module not found")
	}
}
