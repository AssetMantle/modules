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
	SplitList []mappables.Split
}

var _ helpers.Genesis = (*genesis)(nil)

func (genesis genesis) Default() helpers.Genesis {
	return genesis
}

func (genesis genesis) Validate() error {
	for _, split := range genesis.SplitList {
		var _, Error = govalidator.ValidateStruct(split)
		if Error != nil {
			return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
		}
		if split.GetSplit().LT(sdkTypes.ZeroDec()) {
			return xprtErrors.InsufficientBalance
		}
	}
	return nil
}

func (genesis genesis) Initialize(ctx sdkTypes.Context, mapper helpers.Mapper) {
	for _, split := range genesis.SplitList {
		mapper.Create(ctx, split)
	}
}

func (genesis genesis) Export(context sdkTypes.Context, mapper helpers.Mapper) helpers.Genesis {
	assetsID := base.NewID("")

	appendableAssetList := func(mappable traits.Mappable) bool {
		genesis.SplitList = append(genesis.SplitList, mappable.(mappables.Split))
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

func newGenesis(splitList []mappables.Split) helpers.Genesis {
	return genesis{
		SplitList: splitList,
	}
}

var Genesis = newGenesis([]mappables.Split{})
