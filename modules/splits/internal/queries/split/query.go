/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package split

import (
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Query = base.NewQuery(
	"splits",
	"",
	"",

	requestPrototype,
	responsePrototype,
	keeperPrototype,

	flags.SplitID,
)
