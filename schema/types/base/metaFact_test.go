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
	require.Equal(t, stringData.GenerateHash(), testMetaFact.GetHash())
	require.Equal(t, StringType, testMetaFact.GetType())
	require.Equal(t, signatures{}, testMetaFact.GetSignatures())
	//Fix the decData case in GetType Method
	require.Equal(t, DecType, NewMetaFact(decData).GetType())
	require.Equal(t, IDType, NewMetaFact(idData).GetType())
	require.Equal(t, HeightType, NewMetaFact(heightData).GetType())

	readMetaFact, error := ReadMetaFact("S|testString")
	require.Equal(t, testMetaFact, readMetaFact)
	require.Nil(t, error)

	readMetaFact2, error := ReadMetaFact("H|123")
	require.Equal(t, NewMetaFact(heightData), readMetaFact2)
	require.Nil(t, error)

	readMetaFact3, error := ReadMetaFact("I|id")
	require.Equal(t, NewMetaFact(idData), readMetaFact3)
	require.Nil(t, error)

	//Fix the decData case in GetType Method
	readMetaFact4, error := ReadMetaFact("D|12.0")
	require.Equal(t, NewMetaFact(decData), readMetaFact4)
	require.Nil(t, error)

	readMetaFact5, error := ReadMetaFact("randomString")
	require.Equal(t, nil, readMetaFact5)
	require.Equal(t, errors.IncorrectFormat, error)

}
