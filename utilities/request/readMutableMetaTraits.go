/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package request

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

func ReadMutableMetaTraits(Properties string) types.Traits {
	properties := strings.Split(Properties, constants.PropertiesSeparator)
	var traitList []types.Trait
	for _, property := range properties {
		traitList = append(traitList, base.NewTrait(ReadMetaProperty(property), true))
	}
	return base.NewTraits(traitList)
}
