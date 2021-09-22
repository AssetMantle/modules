package base

import (
	"bytes"
	"strings"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type propertyID struct {
	KeyID  types.ID `json:"keyID"`
	TypeID types.ID `json:"typeID"`
}

var _ types.ID = (*propertyID)(nil)

func (propertyID propertyID) String() string {
	var values []string
	values = append(values, propertyID.KeyID.String())
	values = append(values, propertyID.TypeID.String())

	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}
func (propertyID propertyID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, propertyID.KeyID.Bytes()...)
	Bytes = append(Bytes, propertyID.TypeID.Bytes()...)

	return Bytes
}
func (propertyID propertyID) Compare(id types.ID) int {
	return bytes.Compare(propertyID.Bytes(), id.Bytes())
}

func propertyIDFromInterface(id types.ID) propertyID {
	switch value := id.(type) {
	case propertyID:
		return value
	default:
		panic(id)
	}
}
