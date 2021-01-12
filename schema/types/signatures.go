/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

type Signatures interface {
	Get(ID) Signature

	GetList() []Signature

	Add(Signature) Signatures
	Remove(Signature) Signatures
	Mutate(Signature) Signatures
}
