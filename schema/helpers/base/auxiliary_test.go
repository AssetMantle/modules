package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"testing"
)

type auxiliaryKeeperMock struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeperMock)(nil)

func (auxiliaryKeeper auxiliaryKeeperMock) Help(_ sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) helpers.AuxiliaryResponse {
	return newAuxiliaryResponse
}

func (auxiliaryKeeper auxiliaryKeeperMock) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeperMock{mapper: mapper}
}

func keeperPrototypeMock() helpers.Auxiliary {
	return auxiliaryKeeperMock{}
}

func Test_ClassificationID_Help(t *testing.T) {

	classificationID := base.NewID("classificationID")
	//ownerID:= base.NewID("ownerID")

	//testAux:=NewAuxiliary("testAux",)
}
