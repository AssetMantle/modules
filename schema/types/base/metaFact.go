/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"strings"

	"github.com/99designs/keyring"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var (
	_ types.MetaFact = (*MetaFact)(nil)
)

func (metaFact MetaFact) GetHashID() types.ID {
	return metaFact.GetData().GenerateHashID()
}
func (metaFact MetaFact) GetTypeID() types.ID {
	return metaFact.GetData().GetTypeID()
}
func (metaFact MetaFact) GetSignatures() types.Signatures { return metaFact.Signatures }

func (metaFact MetaFact) Sign(_ keyring.Keyring) types.Fact {
	// TODO implement signing
	return &metaFact
}
func (metaFact MetaFact) GetData() types.Data {
	if metaFact.Data == nil {
		return nil
	}

	if metaFact.Data.GetCachedValue() == nil {
		panic(errors.EmptyDataCacheValue)
	}
	data, ok := metaFact.Data.GetCachedValue().(types.Data)
	if !ok {
		panic(errors.InvalidDataCacheValue)
	}
	return data
}
func (metaFact MetaFact) setData(data types.Data) MetaFact {
	if data == nil {
		metaFact.Data = nil
		return metaFact
	}
	any, err := codecTypes.NewAnyWithValue(data)
	if err != nil {
		panic(err)
	}
	metaFact.Data = any
	return metaFact
}

func (metaFact MetaFact) ToFact() types.Fact {
	return NewFact(metaFact.GetData())
}

func NewMetaFact(data types.Data) MetaFact {
	return MetaFact{}.setData(data)
}

func ReadMetaFact(metaFactString string) (types.MetaFact, error) {
	dataTypeAndString := strings.SplitN(metaFactString, constants.DataTypeAndValueSeparator, 2)
	if len(dataTypeAndString) == 2 {
		dataType, dataString := dataTypeAndString[0], dataTypeAndString[1]

		var data types.Data

		var Error error

		switch NewID(dataType).String() {
		case DecData{}.GetTypeID().String():
			data, Error = ReadDecData(dataString)
		case IDData{}.GetTypeID().String():
			data, Error = ReadIDData(dataString)
		case HeightData{}.GetTypeID().String():
			data, Error = ReadHeightData(dataString)
		case StringData{}.GetTypeID().String():
			data, Error = ReadStringData(dataString)
		case AccAddressData{}.GetTypeID().String():
			data, Error = ReadAccAddressData(dataString)
		case ListData{}.GetTypeID().String():
			data, Error = ReadAccAddressListData(dataString)
		default:
			data, Error = nil, errors.UnsupportedParameter
		}

		if Error != nil {
			return nil, Error
		}

		metaFact := NewMetaFact(data)
		return &metaFact, nil
	}

	return nil, errors.IncorrectFormat
}

func (metaFact MetaFact) UnpackInterfaces(unpacker codecTypes.AnyUnpacker) error {
	var data types.Data
	err := unpacker.UnpackAny(metaFact.Data, &data)
	if err != nil {
		return err
	}
	return nil
}
