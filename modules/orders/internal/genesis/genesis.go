/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package genesis

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type genesis struct {
	OrderList []mappables.Order
}

var _ helpers.Genesis = (*genesis)(nil)

func (genesis genesis) Default() helpers.Genesis {
	return genesis
}

func (genesis genesis) Validate() error {
	for _, order := range genesis.OrderList {
		var _, Error = govalidator.ValidateStruct(order)
		if Error != nil {
			return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
		}
	}
	return nil
}

func (genesis genesis) Initialize(ctx sdkTypes.Context, mapper helpers.Mapper) {

	for _, order := range genesis.OrderList {
		mapper.Create(ctx, order)
	}
}

func (genesis genesis) Export(context sdkTypes.Context, mapper helpers.Mapper) helpers.Genesis {
	assetsID := base.NewID("")

	appendableAssetList := func(mappable traits.Mappable) bool {
		genesis.OrderList = append(genesis.OrderList, mappable.(mappables.Order))
		return false
	}
	mapper.Iterate(context, assetsID, appendableAssetList)
	return genesis
}

func (genesis genesis) Marshall() []byte {
	return packageCodec.MustMarshalJSON(genesis)
}
func (genesis genesis) Unmarshall(byte []byte) helpers.Genesis {
	if Error := packageCodec.UnmarshalJSON(byte, &genesis); Error != nil {
		return nil
	}
	return genesis
}

func newGenesis(orderList []mappables.Order) helpers.Genesis {
	return genesis{
		OrderList: orderList,
	}
}

var Genesis = newGenesis([]mappables.Order{})
