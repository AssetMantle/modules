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
	ClassificationList []mappables.Classification
}

var _ helpers.Genesis = (*genesis)(nil)

func (genesis genesis) Default() helpers.Genesis {
	return genesis
}

func (genesis genesis) Validate() error {

	for _, classification := range genesis.ClassificationList {
		var _, Error = govalidator.ValidateStruct(classification)
		if Error != nil {
			return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
		}
	}
	return nil
}

func (genesis genesis) Initialize(ctx sdkTypes.Context, mapper helpers.Mapper) {

	for _, cls := range genesis.ClassificationList {
		mapper.Create(ctx, cls)
	}
}

func (genesis genesis) Export(context sdkTypes.Context, mapper helpers.Mapper) helpers.Genesis {
	assetsID := base.NewID("")

	appendableAssetList := func(mappable traits.Mappable) bool {
		genesis.ClassificationList = append(genesis.ClassificationList, mappable.(mappables.Classification))
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

func newGenesis(classificationList []mappables.Classification) helpers.Genesis {
	return genesis{
		ClassificationList: classificationList,
	}
}

var Genesis = newGenesis([]mappables.Classification{})
