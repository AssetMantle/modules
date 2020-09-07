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
	MaintainersList []mappables.Maintainer
}

var _ helpers.Genesis = (*genesis)(nil)

func (genesis genesis) Default() helpers.Genesis {
	return genesis
}

func (genesis genesis) Validate() error {
	for _, maintainers := range genesis.MaintainersList {
		var _, Error = govalidator.ValidateStruct(maintainers)
		if Error != nil {
			return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
		}
	}
	return nil
}

func (genesis genesis) Initialize(ctx sdkTypes.Context, mapper helpers.Mapper) {

	for _, cls := range genesis.MaintainersList {
		mapper.Create(ctx, cls)
	}
}

func (genesis genesis) Export(context sdkTypes.Context, mapper helpers.Mapper) helpers.Genesis {
	assetsID := base.NewID("")

	appendableAssetList := func(mappable traits.Mappable) bool {
		genesis.MaintainersList = append(genesis.MaintainersList, mappable.(mappables.Maintainer))
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

func newGenesis(MaintainersList []mappables.Maintainer) helpers.Genesis {
	return genesis{
		MaintainersList: MaintainersList,
	}
}

var Genesis = newGenesis([]mappables.Maintainer{})
