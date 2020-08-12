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

func ReadMetaProperties(MetaProperties string) []types.Property {
	properties := strings.Split(MetaProperties, constants.PropertiesSeparator)
	var propertyList []types.Property
	for _, property := range properties {
		propertyIDAndFact := strings.Split(property, constants.PropertyIDAndFactSeparator)
		if len(propertyIDAndFact) == 2 && propertyIDAndFact[0] != "" {
			propertyList = append(propertyList, base.NewProperty(base.NewID(propertyIDAndFact[0]), base.NewMetaFact(propertyIDAndFact[1])))
		}
	}
	return propertyList
}
