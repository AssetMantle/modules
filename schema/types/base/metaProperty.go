/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"strings"
)

var _ types.Property = (*property)(nil)

type metaProperty struct {
	ID       types.ID       `json:"id"`
	MetaFact types.MetaFact `json:"metaFact"`
}

func (metaProperty metaProperty) GetMetaFact() types.MetaFact { return metaProperty.MetaFact }

func (metaProperty metaProperty) RemoveData() types.Property {
	return NewProperty(metaProperty.ID, metaProperty.MetaFact.RemoveData())
}

func (metaProperty metaProperty) GetID() types.ID { return metaProperty.ID }

func (metaProperty metaProperty) GetFact() types.Fact { return metaProperty.MetaFact }

func NewMetaProperty(id types.ID, metaFact types.MetaFact) types.MetaProperty {
	return metaProperty{
		ID:       id,
		MetaFact: metaFact,
	}
}
func ReadMetaProperty(PropertyIDAndStringData string) types.MetaProperty {
	propertyIDAndStringDataList := strings.Split(PropertyIDAndStringData, constants.PropertyIDAndDataSeparator)
	if len(propertyIDAndStringDataList) == 2 && propertyIDAndStringDataList[0] != "" {
		return NewMetaProperty(NewID(propertyIDAndStringDataList[0]), NewMetaFact(NewStringData(propertyIDAndStringDataList[1])))
	}
	return nil
}
