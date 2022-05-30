// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

// TODO move to list
type Signatures interface {
	Get(ID) Signature

	GetList() []Signature

	Add(Signature) Signatures
	Remove(Signature) Signatures
	Mutate(Signature) Signatures
}
