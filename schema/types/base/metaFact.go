/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type metaFact struct {
	Data       types.Data       `json:"meta"`
	Signatures types.Signatures `json:"signatures"`
}

var _ types.MetaFact = (*metaFact)(nil)

func (metaFact metaFact) GetData() types.Data             { return metaFact.Data }
func (metaFact metaFact) RemoveData() types.Fact          { return NewFact(metaFact.Data) }
func (metaFact metaFact) GetHash() string                 { return metaFact.Data.GenerateHash() }
func (metaFact metaFact) GetSignatures() types.Signatures { return metaFact.Signatures }

func (metaFact metaFact) Sign(_ keyring.Keyring) types.Fact {
	//TODO implement signing
	return metaFact
}

func NewMetaFact(data types.Data) types.MetaFact {
	return metaFact{
		Data:       data,
		Signatures: signatures{},
	}
}
