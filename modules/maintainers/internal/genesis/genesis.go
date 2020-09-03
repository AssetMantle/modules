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

type genesisState struct {
	MaintainersList []mappables.Maintainer
}

var _ helpers.GenesisState = (*genesisState)(nil)

func (genesisState genesisState) Default() helpers.GenesisState {
	return genesisState
}

func (genesisState genesisState) Validate() error {
	for _, maintainers := range genesisState.MaintainersList {
		var _, Error = govalidator.ValidateStruct(maintainers)
		if Error != nil {
			return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
		}
	}
	return nil
}

func (genesisState genesisState) Initialize(ctx sdkTypes.Context, mapper helpers.Mapper) {

	for _, cls := range genesisState.MaintainersList {
		mapper.Create(ctx, cls)
	}
}

func (genesisState genesisState) Export(context sdkTypes.Context, mapper helpers.Mapper) helpers.GenesisState {
	assetsID := base.NewID("")

	appendableAssetList := func(mappable traits.Mappable) bool {
		genesisState.MaintainersList = append(genesisState.MaintainersList, mappable.(mappables.Maintainer))
		return false
	}
	mapper.Iterate(context, assetsID, appendableAssetList)
	return genesisState
}

func (genesisState genesisState) Marshall() []byte {
	return packageCodec.MustMarshalJSON(genesisState)
}
func (genesisState genesisState) Unmarshall(byte []byte) helpers.GenesisState {
	if Error := packageCodec.UnmarshalJSON(byte, &genesisState); Error != nil {
		return nil
	}
	return genesisState
}

func newGenesisState(MaintainersList []mappables.Maintainer) helpers.GenesisState {
	return genesisState{
		MaintainersList: MaintainersList,
	}
}

var State = newGenesisState([]mappables.Maintainer{})
