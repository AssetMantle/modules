package mapper

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Mapper = base.NewMapper(
	ModuleName,
	generateKey,
	identityPrototype,
	RegisterCodec,
)
