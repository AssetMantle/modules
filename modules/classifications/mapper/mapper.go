/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Mapper = base.NewMapper(
	ModuleName,
	generateKey,
	classificationPrototype,
	registerCodec,
)
