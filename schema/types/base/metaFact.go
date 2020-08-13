/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type metaFact struct {
	Fact       string           `json:"fact"`
	Hash       string           `json:"hash"`
	Signatures types.Signatures `json:"signatures"`
}

var _ types.Fact = (*metaFact)(nil)

func (metaFact metaFact) Get() string                     { return metaFact.Fact }
func (metaFact metaFact) GetHash() string                 { return metaFact.Hash }
func (metaFact metaFact) GetSignatures() types.Signatures { return metaFact.Signatures }
func (metaFact metaFact) IsMeta() bool {
	return true
}
func (metaFact metaFact) Sign(_ keyring.Keyring) types.Fact {
	//TODO implement signing
	return metaFact
}

func NewMetaFact(fact string) types.Fact {
	return metaFact{
		Hash:       metaUtilities.Hash(fact),
		Fact:       fact,
		Signatures: signatures{},
	}
}
