package charge

import (
	"context"
	"fmt"
	"github.com/AssetMantle/modules/modules/classifications/internal/key"
	"github.com/AssetMantle/modules/modules/classifications/internal/mappable"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(auxiliaryRequest.ClassificationID))

	mappable := mappable.GetClassification(classifications.Get(key.NewKey(auxiliaryRequest.ClassificationID)))
	for _, i := range mappable.GetImmutables().GetImmutablePropertyList().GetList() {
		if i.Get().GetID().AsString() == constants.BondingIDString {
			coins, err := sdkTypes.ParseCoinsNormalized(i.Get().(properties.MetaProperty).GetData().Get().AsString() + "stake")
			if err != nil {
				fmt.Println("Incorrect format: ", err.Error())
			}
			if auxiliaryRequest.bondingMode {
				if err := auxiliaryRequest.bankKeeper.SendCoinsFromAccountToModule(sdkTypes.UnwrapSDKContext(context), auxiliaryRequest.address, auxiliaryRequest.moduleName, coins); err != nil {
					fmt.Println("error")
				}
			} else {
				if err := auxiliaryRequest.bankKeeper.SendCoinsFromModuleToAccount(sdkTypes.UnwrapSDKContext(context), auxiliaryRequest.moduleName, auxiliaryRequest.address, coins); err != nil {
					fmt.Println("error")
				}
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
