/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/99designs/keyring"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"strings"
)

type metaFact struct {
	Data       types.Data       `json:"data"`
	Signatures types.Signatures `json:"signatures"`
}

var _ types.MetaFact = (*metaFact)(nil)

func (metaFact metaFact) GetData() types.Data             { return metaFact.Data }
func (metaFact metaFact) RemoveData() types.Fact          { return NewFact(metaFact.Data) }
func (metaFact metaFact) GetHash() string                 { return metaFact.Data.GenerateHash() }
func (metaFact metaFact) GetType() string                 { return metaFact.Data.Type() }
func (metaFact metaFact) GetSignatures() types.Signatures { return metaFact.Signatures }

func (metaFact metaFact) Sign(_ keyring.Keyring) types.Fact {
	//TODO implement signing
	return metaFact
}

func NewMetaFact(data types.Data) types.MetaFact {
	return metaFact{
		Data:       data,
		Signatures: signatures{},
	}
}

func ReadMetaFact(DataTypeAndString string) (types.MetaFact, error) {
	dataTypeAndString := strings.Split(DataTypeAndString, constants.DataTypeAndValueSeparator)
	if len(dataTypeAndString) == 2 {
		dataType, dataString := dataTypeAndString[0], dataTypeAndString[1]
		var data types.Data
		var Error error
		switch dataType {
		case decData{}.Type():
			data, Error = ReadDecData(dataString)
		case idData{}.Type():
			data, Error = ReadIDData(dataString)
		case heightData{}.Type():
			data, Error = ReadHeightData(dataString)
		case stringData{}.Type():
			data, Error = ReadStringData(dataString)
		default:
			data, Error = nil, errors.UnsupportedParameter
		}
		if Error != nil {
			return nil, Error
		} else {
			return NewMetaFact(data), nil
		}
	} else {
		return nil, errors.IncorrectFormat
	}
}
