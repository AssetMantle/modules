package mapper

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var Mapper = base.NewMapper(
	ModuleName,
	generateKey,
	exchangePrototype,
	registerCodec,
)

func exchangePrototype() traits.Mappable {
	return nil
}

func generateKey(exchangeID types.ID) []byte {
	return StoreKeyPrefix
}
