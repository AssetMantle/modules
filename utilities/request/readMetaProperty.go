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

func ReadMetaProperty(PropertyIDAndMetaFact string) types.Property {
	propertyIDAndFactList := strings.Split(PropertyIDAndMetaFact, constants.PropertyIDAndMetaFactSeparator)
	if len(propertyIDAndFactList) == 2 && propertyIDAndFactList[0] != "" {
		return base.NewProperty(base.NewID(propertyIDAndFactList[0]), base.NewMetaFact(propertyIDAndFactList[1]))
	}
	return nil
}
