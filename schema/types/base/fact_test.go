package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
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
	require.Equal(t, StringType, testFact.GetType())
	require.Equal(t, DecType, NewFact(decData).GetType())
	require.Equal(t, IDType, NewFact(idData).GetType())
	require.Equal(t, HeightType, NewFact(heightData).GetType())
}
