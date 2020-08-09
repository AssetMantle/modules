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

var _ types.Fact = (*fact)(nil)

type fact struct {
	Hash       string           `json:"hash"`
	Signatures types.Signatures `json:"signatures"`
	Meta       bool
}

func (fact fact) GetHash() string                 { return fact.Hash }
func (fact fact) GetSignatures() types.Signatures { return fact.Signatures }
func (fact fact) IsMeta() bool {
	return fact.Meta
}
func (fact fact) Sign(_ keyring.Keyring) types.Fact {
	//TODO implement signing
	return fact
}
func NewFact(Fact string, Meta bool) types.Fact {
	return fact{
		Hash:       metaUtilities.Hash(Fact),
		Signatures: signatures{},
		Meta:       Meta,
	}
}
