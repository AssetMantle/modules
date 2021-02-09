/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package auxiliaries

import (
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/deputize"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/maintain"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/revoke"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/auxiliaries/super"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func Prototype() helpers.Auxiliaries {
	return base.NewAuxiliaries(
		deputize.Auxiliary,
		maintain.Auxiliary,
		revoke.Auxiliary,
		super.Auxiliary,
	)
}
