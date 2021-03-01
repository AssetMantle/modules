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

type metaFact struct {
	Data       types.Data       `json:"data"`
	Signatures types.Signatures `json:"signatures"`
}

var _ types.MetaFact = (*metaFact)(nil)

func (metaFact metaFact) GetHashID() types.ID             { return metaFact.Data.GenerateHashID() }
func (metaFact metaFact) GetTypeID() types.ID             { return metaFact.Data.GetTypeID() }
func (metaFact metaFact) GetSignatures() types.Signatures { return metaFact.Signatures }
func (metaFact metaFact) Sign(_ keyring.Keyring) types.MetaFact {
	// TODO implement signing
	return metaFact
}
func (metaFact metaFact) GetData() types.Data    { return metaFact.Data }
func (metaFact metaFact) RemoveData() types.Fact { return NewFact(metaFact.Data) }

func NewMetaFact(data types.Data) types.MetaFact {
	return metaFact{
		Data:       data,
		Signatures: signatures{},
	}
}

func ReadMetaFact(metaFactString string) (types.MetaFact, error) {
	dataTypeAndString := strings.SplitN(metaFactString, constants.DataTypeAndValueSeparator, 2)
	if len(dataTypeAndString) == 2 {
		dataType, dataString := dataTypeAndString[0], dataTypeAndString[1]

		var data types.Data

		var Error error

		switch NewID(dataType) {
		case decData{}.GetTypeID():
			data, Error = ReadDecData(dataString)
		case idData{}.GetTypeID():
			data, Error = ReadIDData(dataString)
		case heightData{}.GetTypeID():
			data, Error = ReadHeightData(dataString)
		case stringData{}.GetTypeID():
			data, Error = ReadStringData(dataString)
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
