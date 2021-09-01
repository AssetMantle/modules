/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"strings"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.MetaProperty = (*MetaProperty)(nil)

func (metaProperty MetaProperty) GetMetaFact() types.MetaFact { return &metaProperty.MetaFact }
func (metaProperty MetaProperty) GetID() types.ID             { return &metaProperty.Id }
func (metaProperty MetaProperty) RemoveData() types.Property {
	return NewProperty(metaProperty.Id, *NewFact(&metaProperty.MetaFact.Data))
}

func NewMetaProperty(id ID, metaFact MetaFact) *MetaProperty {
	return &MetaProperty{
		Id:       id,
		MetaFact: metaFact,
	}
}
func ReadMetaProperty(metaPropertyString string) (*MetaProperty, error) {
	propertyIDAndData := strings.Split(metaPropertyString, constants.PropertyIDAndDataSeparator)
	if len(propertyIDAndData) == 2 && propertyIDAndData[0] != "" {
		metaFact, Error := ReadMetaFact(propertyIDAndData[1])
		if Error != nil {
			return nil, Error
		}

		return NewMetaProperty(*NewID(propertyIDAndData[0]), *metaFact), nil
	}

	return nil, errors.IncorrectFormat
}
