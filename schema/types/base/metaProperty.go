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

type metaProperty struct {
	ID       types.ID       `json:"id"`
	MetaFact types.MetaFact `json:"metaFact"`
}

var _ types.MetaProperty = (*metaProperty)(nil)

func (metaProperty metaProperty) GetMetaFact() types.MetaFact { return metaProperty.MetaFact }
func (metaProperty metaProperty) GetID() types.ID             { return metaProperty.ID }
func (metaProperty metaProperty) RemoveData() types.Property {
	return NewProperty(metaProperty.ID, metaProperty.MetaFact.RemoveData())
}

func NewMetaProperty(id types.ID, metaFact types.MetaFact) types.MetaProperty {
	return metaProperty{
		ID:       id,
		MetaFact: metaFact,
	}
}
func ReadMetaProperty(metaPropertyString string) (types.MetaProperty, error) {
	propertyIDAndData := strings.Split(metaPropertyString, constants.PropertyIDAndDataSeparator)
	if len(propertyIDAndData) == 2 && propertyIDAndData[0] != "" {
		metaFact, Error := ReadMetaFact(propertyIDAndData[1])
		if Error != nil {
			return nil, Error
		}

		return NewMetaProperty(NewID(propertyIDAndData[0]), metaFact), nil
	}

	return nil, errors.IncorrectFormat
}
