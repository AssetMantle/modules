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

func ReadProperties(Properties string) types.Properties {
	var propertyList []types.Property
	properties := strings.Split(Properties, constants.PropertiesSeparator)
	for _, property := range properties {
		propertyList = append(propertyList, ReadProperty(property))
	}
	return base.NewProperties(propertyList)
}
