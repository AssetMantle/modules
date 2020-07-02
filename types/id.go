package types

import (
	"bytes"
)

type ID interface {
	String() string
	Bytes() []byte
	Compare(ID) int
}

type id struct {
	IDString string
}

var _ ID = (*id)(nil)

func (id id) String() string {
	return id.IDString
}

func (id id) Bytes() []byte {
	return []byte(id.IDString)
}

func (id id) Compare(ID ID) int {
	return bytes.Compare(id.Bytes(), ID.Bytes())
}

func NewID(idString string) ID {
	return &id{IDString: idString}
}
