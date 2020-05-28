package types

import (
	"bytes"
	"encoding/base64"
)

type ID interface {
	String() string
	Bytes() []byte
	Compare(ID) int
}

type BaseID struct {
	BaseBytes []byte
}

var _ ID = (*BaseID)(nil)

func (baseID BaseID) String() string    { return base64.URLEncoding.EncodeToString(baseID.BaseBytes) }
func (baseID BaseID) Bytes() []byte     { return baseID.BaseBytes }
func (baseID BaseID) Compare(id ID) int { return bytes.Compare(baseID.BaseBytes, id.Bytes()) }
