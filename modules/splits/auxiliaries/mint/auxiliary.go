package mint

import (
	"github.com/persistenceOne/persistenceSDK/modules/splits/mapper"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

var Auxiliary = utility.NewAuxiliary(
	mapper.ModuleName,
	AuxiliaryName,
	AuxiliaryRoute,
	initializeAuxiliaryKeeper,
)
