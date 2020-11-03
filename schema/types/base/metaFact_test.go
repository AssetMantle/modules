package base

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Meta_Fact(t *testing.T) {

	stringData := NewStringData("testString")
	//decData:= NewDecData(sdkTypes.NewDec(12))
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
	//require.Equal(t,DecType,NewMetaFact(decData).GetType())
	require.Equal(t, IDType, NewMetaFact(idData).GetType())
	require.Equal(t, HeightType, NewMetaFact(heightData).GetType())

}
