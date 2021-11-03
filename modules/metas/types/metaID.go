package types

import (
	"bytes"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

var (
	_ types.ID                           = (*MetaID)(nil)
	_ helpers.Key                        = (*MetaID)(nil)
	_ codecTypes.UnpackInterfacesMessage = (*Meta)(nil)
)

func (metaID MetaID) String() string {
	var values []string
	values = append(values, metaID.TypeID.String())
	values = append(values, metaID.HashID.String())

	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}

func (metaID MetaID) Bytes() []byte {
	var b []byte
	b = append(b, metaID.TypeID.Bytes()...)
	b = append(b, metaID.HashID.Bytes()...)
	return b
}

func (metaID MetaID) Compare(id types.ID) int {
	return bytes.Compare(metaID.Bytes(), id.Bytes())
}

func (metaID MetaID) GenerateStoreKeyBytes() []byte {
	return StoreKeyPrefix.GenerateStoreKey(metaID.Bytes())
}

func (metaID MetaID) IsPartial() bool {
	return len(metaID.HashID.Bytes()) == 0
}
func (metaID MetaID) Equals(key helpers.Key) bool {
	m := MetaIDFromInterface(key)
	return metaID.Compare(&m) == 0
}

func readMetaID(metaIDString string) types.ID {
	idList := strings.Split(metaIDString, constants.FirstOrderCompositeIDSeparator)
	if len(idList) == 2 {
		return &MetaID{
			TypeID: base.NewID(idList[0]),
			HashID: base.NewID(idList[1]),
		}
	}

	return &MetaID{TypeID: base.NewID(""), HashID: base.NewID("")}
}
func MetaIDFromInterface(i interface{}) MetaID {
	switch value := i.(type) {
	case MetaID:
		return value
	case *MetaID:
		return *value
	case types.ID:
		return MetaIDFromInterface(readMetaID(value.String()))
	default:
		panic(i)
	}
}

func GenerateMetaID(data types.Data) MetaID {
	return MetaID{
		TypeID: base.NewID(data.GetTypeID().String()),
		HashID: base.NewID(data.GenerateHashID().String()),
	}
}

func FromID(id types.ID) helpers.Key {
	m := MetaIDFromInterface(id)
	return &m
}
