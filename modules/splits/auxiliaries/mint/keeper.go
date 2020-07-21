package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type auxiliaryKeeper struct {
	mapper utility.Mapper
}

var _ utility.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest utility.AuxiliaryRequest) error {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	splitID := mapper.NewSplitID(auxiliaryRequest.OwnerID, auxiliaryRequest.OwnableID)
	splits := mapper.NewSplits(auxiliaryKeeper.mapper, context).Fetch(splitID)
	split := splits.Get(splitID)
	if split == nil {
		splits.Add(mapper.NewSplit(splitID, auxiliaryRequest.Split))
	} else {
		splits.Mutate(split.Receive(auxiliaryRequest.Split).(schema.Split))
	}
	return nil
}

func initializeAuxiliaryKeeper(mapper utility.Mapper) utility.AuxiliaryKeeper {
	return auxiliaryKeeper{mapper: mapper}
}
