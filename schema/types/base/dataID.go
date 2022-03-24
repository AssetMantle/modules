package base

import (
	"bytes"
	"strings"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type dataID struct {
	TypeID types.ID `json:"typeID"`
	HashID types.ID `json:"hashID"`
}

var _ types.ID = (*dataID)(nil)

func (dataID dataID) String() string {
	var values []string
	values = append(values, dataID.TypeID.String())
	values = append(values, dataID.HashID.String())

	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}
func (dataID dataID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, dataID.TypeID.Bytes()...)
	Bytes = append(Bytes, dataID.HashID.Bytes()...)

	return Bytes
}
func (dataID dataID) Compare(id types.ID) int {
	return bytes.Compare(dataID.Bytes(), id.Bytes())
}

func dataIDFromInterface(id types.ID) dataID {
	switch value := id.(type) {
	case dataID:
		return value
	default:
		panic(id)
	}
}
