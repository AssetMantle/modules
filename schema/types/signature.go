// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/tendermint/tendermint/crypto"

	"github.com/AssetMantle/modules/schema/ids"
)

type Signature interface {
	String() string
	Bytes() []byte

	GetID() ids.ID

	Verify(crypto.PubKey, []byte) bool
	GetValidityHeight() Height
	HasExpired(Height) bool
}
