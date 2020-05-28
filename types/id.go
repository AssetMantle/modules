package types

import (
	"bytes"
)

type ID interface {
	String() string
	Bytes() []byte
	Compare(ID) int
}

type BaseID struct {
	BaseString string
}

var _ ID = (*BaseID)(nil)

func (baseID BaseID) String() string    { return baseID.BaseString }
func (baseID BaseID) Bytes() []byte     { return []byte(baseID.BaseString) }
func (baseID BaseID) Compare(id ID) int { return bytes.Compare(baseID.Bytes(), id.Bytes()) }
