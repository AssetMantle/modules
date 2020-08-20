/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/99designs/keyring"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type fact struct {
	Hash       string           `json:"hash"`
	Signatures types.Signatures `json:"signatures"`
}

var _ types.Fact = (*fact)(nil)

func (fact fact) Get() string                     { return "" }
func (fact fact) GetHash() string                 { return fact.Hash }
func (fact fact) GetSignatures() types.Signatures { return fact.Signatures }
func (fact fact) IsMeta() bool {
	return false
}
func (fact fact) Sign(_ keyring.Keyring) types.Fact {
	//TODO implement signing
	return fact
}

func NewFact(data types.Data) types.Fact {
	return fact{
		Hash:       data.GenerateHash(),
		Signatures: signatures{},
	}
}

func ReadFact(DataTypeAndString string) types.Fact {
	return ReadMetaFact(DataTypeAndString).RemoveData()
}
