// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"github.com/gogo/protobuf/proto"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/properties"
)

type Parameter interface {
	proto.Message
	ValidateBasic() error
	GetMetaProperty() properties.MetaProperty
	Mutate(data.Data) Parameter
}
