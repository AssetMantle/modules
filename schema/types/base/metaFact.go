/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"strings"

	"github.com/99designs/keyring"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ types.MetaFact = (*MetaFact)(nil)

func (metaFact MetaFact) GetHashID() types.ID             { return metaFact.Data.GenerateHashID() }
func (metaFact MetaFact) GetTypeID() types.ID             { return metaFact.Data.GetTypeID() }
func (metaFact MetaFact) GetSignatures() types.Signatures { return metaFact.Signatures }
func (metaFact MetaFact) Sign(_ keyring.Keyring) types.MetaFact {
	// TODO implement signing
	return &metaFact
}
func (metaFact MetaFact) GetData() types.Data    { return metaFact.Data }
func (metaFact MetaFact) RemoveData() types.Fact { return NewFact(metaFact.Data) }

func NewMetaFact(data types.Data) types.MetaFact {
	return &MetaFact{
		Data:       data,
		Signatures: &Signatures{},
	}
}

func ReadMetaFact(metaFactString string) (types.MetaFact, error) {
	dataTypeAndString := strings.SplitN(metaFactString, constants.DataTypeAndValueSeparator, 2)
	if len(dataTypeAndString) == 2 {
		dataType, dataString := dataTypeAndString[0], dataTypeAndString[1]

		var data types.Data

		var Error error

		switch NewID(dataType) {
		case DecData{}.GetTypeID():
			data, Error = ReadDecData(dataString)
		case IDData{}.GetTypeID():
			data, Error = ReadIDData(dataString)
		case HeightData{}.GetTypeID():
			data, Error = ReadHeightData(dataString)
		case StringData{}.GetTypeID():
			data, Error = ReadStringData(dataString)
		case AccAddressData{}.GetTypeID():
			data, Error = ReadAccAddressData(dataString)
		case ListData{}.GetTypeID():
			data, Error = ReadAccAddressListData(dataString)
		default:
			data, Error = nil, errors.UnsupportedParameter
		}

		if Error != nil {
			return nil, Error
		}

		return NewMetaFact(data), nil
	}

	return nil, errors.IncorrectFormat
}
