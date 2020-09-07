/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package genesis

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type genesis struct {
	AssetList     []mappables.InterNFT `json:"assetList"`
	ParameterList []types.Parameter    `json:"parameterList"`
}

var _ helpers.Genesis = (*genesis)(nil)

func (genesis genesis) Default() helpers.Genesis {
	return genesis
}

func (genesis genesis) Validate() error {
	_, Error := govalidator.ValidateStruct(genesis)
	return Error
}

func (genesis genesis) Initialize(context sdkTypes.Context, Mapper helpers.Mapper) {
	assets := mapper.NewAssets(context, Mapper)
	for _, asset := range genesis.AssetList {
		assets = assets.Add(asset)
	}
}

func (genesis genesis) Export(context sdkTypes.Context, Mapper helpers.Mapper) helpers.Genesis {
	assets := mapper.NewAssets(context, Mapper).Fetch(base.NewID(""))
	//TODO add parameters
	return NewGenesis(assets.GetList(), nil)
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

func NewGenesis(assetList []mappables.InterNFT, parameterList ...types.Parameter) helpers.Genesis {
	return genesis{
		AssetList:     assetList,
		ParameterList: parameterList,
	}
}

var Genesis = NewGenesis([]mappables.InterNFT{}, []types.Parameter{}...)
