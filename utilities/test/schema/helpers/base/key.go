// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/schema/x/helpers"
)

// key struct, implements helpers.Key
type testKey struct {
	ID string
}

var _ helpers.Key = (*testKey)(nil)

func (t testKey) String() string {
	return t.ID
}

func (t testKey) GenerateStoreKeyBytes() []byte {
	return append([]byte{0x11}, []byte(t.ID)...)
}

func (t testKey) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	legacyAmino.RegisterConcrete(testKey{}, "test/testKey", nil)
}

func (t testKey) IsPartial() bool {
	return t.ID != ""
}

func (t testKey) Equals(key helpers.Key) bool {
	return bytes.Equal([]byte(t.ID), []byte(key.(testKey).ID))
}

func NewKey(id string) helpers.Key {
	return testKey{ID: id}
}

func KeyPrototype() helpers.Key {
	return testKey{}
}
