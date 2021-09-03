/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/99designs/keyring"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.Fact = (*Fact)(nil)

func (fact Fact) GetHashID() types.ID             { return &fact.HashId }
func (fact Fact) GetTypeID() types.ID             { return &fact.TypeId }
func (fact Fact) GetSignatures() types.Signatures { return &fact.Signatures }
func (fact Fact) IsMeta() bool {
	return false
}
func (fact Fact) Sign(_ keyring.Keyring) types.Fact {
	// TODO use keyring to sign
	cliContext := client.Context{}
	sign, _, _ := cliContext.Keyring.Sign(cliContext.FromName, fact.GetHashID().Bytes())
	Signature := Signature{
		Id:             ID{IdString: fact.GetHashID().String()},
		SignatureBytes: sign,
		ValidityHeight: Height{cliContext.Height},
	}
	fact.GetSignatures().Add(&Signature)

	return &fact
}

func NewFact(data types.Data) *Fact {
	return &Fact{
		HashId:     *NewID(data.GenerateHashID().String()),
		TypeId:     *NewID(data.GetTypeID().String()),
		Signatures: Signatures{},
	}
}

func NewFactProperty(hashID types.ID, typeID types.ID, signatures types.Signatures) *Fact {
	return &Fact{
		HashId:     *NewID(hashID.String()),
		TypeId:     *NewID(typeID.String()),
		Signatures: *NewSignatures(signatures.GetList()),
	}
}

func ReadFact(metaFactString string) (types.Fact, error) {
	metaFact, Error := ReadMetaFact(metaFactString)
	if Error != nil {
		return nil, Error
	}

	return metaFact.RemoveData(), nil
}
