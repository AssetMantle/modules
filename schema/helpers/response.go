// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/gogo/protobuf/proto"
)

type Response interface {
	proto.Message
	IsSuccessful() bool
	GetError() error
}
