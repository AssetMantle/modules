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
	Type       string           `json:"type"`
	Signatures types.Signatures `json:"signatures"`
}

var _ types.Fact = (*fact)(nil)

func (fact fact) GetHash() string                 { return fact.Hash }
func (fact fact) GetType() string                 { return fact.Type }
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
		Type:       data.Type(),
		Signatures: signatures{},
	}
}

func ReadFact(MetaFact string) (types.Fact, error) {
	metaFact, Error := ReadMetaFact(MetaFact)
	if Error != nil {
		return nil, Error
	}
	return metaFact.RemoveData(), nil
}
