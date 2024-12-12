// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package v2

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var Migration = base.NewMigration(func(context sdkTypes.Context, mapper helpers.Mapper, parameterManager helpers.ParameterManager, paramsSubspace paramsTypes.Subspace) error {
	for _, parameter := range parameterManager.GetDefaultParameterList().Get() {
		var value string
		paramsSubspace.Get(sdkTypes.UnwrapSDKContext(context), parameter.GetMetaProperty().GetID().Bytes(), &value)
		if data, err := parameter.GetMetaProperty().GetData().Get().FromString(value); err != nil {
			return err
		} else if parameterManager, err = parameterManager.Set(parameter.Mutate(data)).Update(context); err != nil {
			return err
		}
	}

	return nil
})
