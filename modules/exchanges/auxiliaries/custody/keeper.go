package custody

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/burn"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type auxiliaryKeeper struct {
	mapper              helpers.Mapper
	splitsMintAuxiliary helpers.Auxiliary
	splitsBurnAuxiliary helpers.Auxiliary
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) error {
	//auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)

	return nil
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper, externalKeepers []interface{}) helpers.AuxiliaryKeeper {
	auxiliaryKeeper := auxiliaryKeeper{mapper: mapper}
	for _, externalKeeper := range externalKeepers {
		switch value := externalKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case mint.Auxiliary.GetName():
				auxiliaryKeeper.splitsMintAuxiliary = value
			case burn.Auxiliary.GetName():
				auxiliaryKeeper.splitsBurnAuxiliary = value
			}
		}
	}
	return auxiliaryKeeper
}
