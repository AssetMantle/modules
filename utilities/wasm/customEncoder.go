/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package wasm

import (
	"encoding/json"
	"strings"

	"github.com/CosmWasm/wasmd/x/wasm"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func CustomEncoder(moduleList ...helpers.Module) wasm.CustomEncoder {
	return func(sender sdkTypes.AccAddress, rawMessage json.RawMessage) ([]sdkTypes.Msg, error) {
		wasmMessage := base.WasmMessagePrototype()

		err := json.Unmarshal(rawMessage, &wasmMessage)
		if err != nil {
			return nil, errors.IncorrectMessage
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

		return nil, errors.IncorrectMessage
	}
}
