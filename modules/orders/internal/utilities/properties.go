/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package utilities

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func GetOrderMakerOwnableSplitAndExpiry(supplementAuxiliary helpers.Auxiliary, context sdkTypes.Context, order mappables.Order) (sdkTypes.Dec, int64, error) {
	metaProperties, Error := supplement.GetMetaPropertiesFromResponse(supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(order.GetMakerOwnableSplit(), order.GetExpiry())))
	if Error != nil {
		return sdkTypes.Dec{}, 0, Error
	}

	makerOwnableSplitProperty := metaProperties.GetMetaProperty(base.NewID(properties.MakerOwnableSplit))
	if makerOwnableSplitProperty == nil {
		return sdkTypes.Dec{}, 0, Error
	}

	makerOwnableSplit, Error := makerOwnableSplitProperty.GetMetaFact().GetData().AsDec()
	if Error != nil {
		return sdkTypes.Dec{}, 0, Error
	}

	expiryProperty := metaProperties.GetMetaProperty(base.NewID(properties.Expiry))
	if expiryProperty == nil {
		return sdkTypes.Dec{}, 0, Error
	}

	expiry, Error := expiryProperty.GetMetaFact().GetData().AsHeight()
	if Error != nil {
		return sdkTypes.Dec{}, 0, Error
	}

	return makerOwnableSplit, expiry.Get(), nil
}
