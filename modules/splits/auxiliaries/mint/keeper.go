package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type auxiliaryKeeper struct {
	mapper utilities.Mapper
}

var _ utilities.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest utilities.AuxiliaryRequest) error {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	splitID := mapper.NewSplitID(auxiliaryRequest.OwnerID, auxiliaryRequest.OwnableID)
	splits := mapper.NewSplits(auxiliaryKeeper.mapper, context).Fetch(splitID)
	split := splits.Get(splitID)
	if split == nil {
		splits.Add(mapper.NewSplit(splitID, auxiliaryRequest.Split))
	} else {
		splits.Mutate(split.Receive(auxiliaryRequest.Split).(mappables.Split))
	}
	return nil
}

func initializeAuxiliaryKeeper(mapper utilities.Mapper) utilities.AuxiliaryKeeper {
	return auxiliaryKeeper{mapper: mapper}
}
