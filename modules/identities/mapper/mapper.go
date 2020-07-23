package mapper

import (
	"github.com/persistenceOne/persistenceSDK/schema/utilities/base"
)

var Mapper = base.NewMapper(generateKey, identityPrototype, registerCodec)
