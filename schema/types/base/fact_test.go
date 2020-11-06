package base

import (
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
	require.Equal(t, fact{Hash: stringData.GenerateHash(), Type: StringType, Signatures: signatures{}}, testFact)
	require.Equal(t, stringData.GenerateHash(), testFact.GetHash())
	require.Equal(t, signatures{}, testFact.GetSignatures())
	require.Equal(t, false, testFact.(fact).IsMeta())
	require.Equal(t, StringType, testFact.GetType())
	require.Equal(t, DecType, NewFact(decData).GetType())
	require.Equal(t, IDType, NewFact(idData).GetType())
	require.Equal(t, HeightType, NewFact(heightData).GetType())

	readFact, error := ReadFact("S|testString")
	require.Equal(t, testFact, readFact)
	require.Nil(t, error)

	readFact2, error := ReadFact("")
	require.Equal(t, nil, readFact2)
	require.Equal(t, errors.IncorrectFormat, error)
}
