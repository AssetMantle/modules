package verify

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/identities/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) error {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	identities := mapper.NewIdentities(auxiliaryKeeper.mapper, context).Fetch(auxiliaryRequest.IdentityID)
	identity := identities.Get(auxiliaryRequest.IdentityID)
	if identity == nil {
		return constants.EntityNotFound
	}
	if identity.IsProvisioned(auxiliaryRequest.Address) {
		return nil
	} else {
		return constants.NotAuthorized
	}
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper, _ []interface{}) helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{mapper: mapper}
}
