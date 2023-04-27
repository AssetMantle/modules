// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package wasm

import (
	"encoding/json"
	"strings"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/CosmWasm/wasmd/x/wasm"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

func CustomEncoder(moduleList ...helpers.Module) wasm.CustomEncoder {
	return func(sender sdkTypes.AccAddress, rawMessage json.RawMessage) ([]sdkTypes.Msg, error) {
		wasmMessage := baseHelpers.WasmMessagePrototype()

		err := json.Unmarshal(rawMessage, &wasmMessage)
		if err != nil {
			return nil, errorConstants.IncorrectMessage.Wrapf(err.Error())
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

		return nil, errorConstants.IncorrectMessage.Wrapf("module not found")
	}
}
