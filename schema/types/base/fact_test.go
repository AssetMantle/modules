/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Fact(t *testing.T) {

	stringData := NewStringData("testString")
	decData := NewDecData(sdkTypes.NewDec(12))
	idData := NewIDData(NewID("id"))
	heightData := NewHeightData(NewHeight(123))

	testFact := NewFact(stringData)
	require.Equal(t, Fact{HashId: stringData.GenerateHashID(), TypeId: NewID("S"), Signatures: Signatures{}}, testFact)
	require.Equal(t, stringData.GenerateHashID(), testFact.GetHashID())
	require.Equal(t, Signatures{}, testFact.GetSignatures())
	require.Equal(t, false, testFact.(Fact).IsMeta())
	require.Equal(t, NewID("S"), testFact.GetTypeID())
	require.Equal(t, NewID("D"), NewFact(decData).GetTypeID())
	require.Equal(t, NewID("I"), NewFact(idData).GetTypeID())
	require.Equal(t, NewID("H"), NewFact(heightData).GetTypeID())

	readFact, Error := ReadFact("S|testString")
	require.Equal(t, testFact, readFact)
	require.Nil(t, Error)
	require.Panics(t, func() {
		require.Equal(t, readFact, readFact.Sign(nil))

	})
	readFact2, Error := ReadFact("")
	require.Equal(t, nil, readFact2)
	require.Equal(t, errors.IncorrectFormat, Error)

	cliContext := client.Context{}
	require.Panics(t, func() {
		sign, _, _ := cliContext.Keyring.Sign(cliContext.FromName, readFact.GetHashID().Bytes())
		Signature := Signature{
			Id:             ID{IdString: readFact.GetHashID().String()},
			SignatureBytes: sign,
			ValidityHeight: Height{cliContext.Height},
		}
		readFact.GetSignatures().Add(Signature)
		require.Equal(t, readFact.GetSignatures().Get(readFact.GetHashID()), readFact.GetHashID().String())
	})

}
