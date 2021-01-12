/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_MetaFact(t *testing.T) {

	stringData := NewStringData("testString")
	decData := NewDecData(sdkTypes.NewDec(12))
	idData := NewIDData(NewID("id"))
	heightData := NewHeightData(NewHeight(123))

	testMetaFact := NewMetaFact(stringData)
	require.Equal(t, metaFact{Data: stringData, Signatures: signatures{}}, testMetaFact)
	require.Equal(t, stringData, testMetaFact.GetData())
	require.Equal(t, NewFact(stringData), testMetaFact.RemoveData())
	require.Equal(t, stringData.GenerateHashID(), testMetaFact.GetHashID())
	require.Equal(t, signatures{}, testMetaFact.GetSignatures())
	require.Equal(t, NewID("S"), testMetaFact.GetTypeID())
	require.Equal(t, NewID("D"), NewMetaFact(decData).GetTypeID())
	require.Equal(t, NewID("I"), NewMetaFact(idData).GetTypeID())
	require.Equal(t, NewID("H"), NewMetaFact(heightData).GetTypeID())

	readMetaFact, Error := ReadMetaFact("S|testString")
	require.Equal(t, testMetaFact, readMetaFact)
	require.Nil(t, Error)

	readMetaFact2, Error := ReadMetaFact("H|123")
	require.Equal(t, NewMetaFact(heightData), readMetaFact2)
	require.Nil(t, Error)

	readMetaFact3, Error := ReadMetaFact("I|id")
	require.Equal(t, NewMetaFact(idData), readMetaFact3)
	require.Nil(t, Error)

	readMetaFact3, Error = ReadMetaFact("I|test.Class|hash")
	require.Equal(t, NewMetaFact(NewIDData(NewID("test.Class|hash"))), readMetaFact3)
	require.Nil(t, Error)

	//Fix the decData case in GetTypeID Method
	readMetaFact4, Error := ReadMetaFact("D|12.0")
	require.Equal(t, NewMetaFact(decData), readMetaFact4)
	require.Nil(t, Error)

	readMetaFact5, Error := ReadMetaFact("Z|12.0")
	require.Equal(t, nil, readMetaFact5)
	require.Equal(t, errors.UnsupportedParameter, Error)

	readMetaFact6, Error := ReadMetaFact("randomString")
	require.Equal(t, nil, readMetaFact6)
	require.Equal(t, errors.IncorrectFormat, Error)

	require.Equal(t, readMetaFact4, readMetaFact4.Sign(nil))
}
