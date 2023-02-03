package charge

import (
	"context"
	"fmt"
	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	mappable2 "github.com/AssetMantle/modules/modules/classifications/internal/mappable"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(auxiliaryRequest.ClassificationID))

	mappable := mappable2.GetClassification(classifications.Get(key.NewKey(auxiliaryRequest.ClassificationID)))
	for _, i := range mappable.GetImmutables().GetImmutablePropertyList().GetList() {
		x := i.Get().GetID().AsString()
		if x == "BondingAmount.S" {
			prop := i.Get().(properties.MetaProperty)
			dat := prop.GetData()
			findat := dat.Get()
			str := findat.AsString()
			coins, err := sdkTypes.ParseCoinsNormalized(str) //i.Get().(properties.MetaProperty).GetData().Get().AsString()
			if err != nil {
				fmt.Println("IDK")
			}
			if err := auxiliaryRequest.bankKeeper.SendCoinsFromAccountToModule(sdkTypes.UnwrapSDKContext(context), auxiliaryRequest.fromAddress, auxiliaryRequest.moduleName, coins); err != nil {
				fmt.Println("error")
			}
			return newAuxiliaryResponse(i.Get().(properties.MetaProperty).GetData().AsString(), nil)
		}
	}
	return newAuxiliaryResponse("", nil)
}

func (auxiliaryKeeper auxiliaryKeeper) FetchCollection(context context.Context, classificationID baseIDs.ClassificationID) helpers.Collection {
	return auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(&classificationID))
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterList, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
