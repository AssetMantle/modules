package base

import (
	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Immutables(t *testing.T) {
	testProperty := NewProperty(NewID("ID"), NewFact(NewHeightData(NewHeight(123))))
	testImmutables := NewImmutables(NewProperties(testProperty))

	require.Equal(t, immutables{Properties: NewProperties(testProperty)}, testImmutables)
	require.Equal(t, NewProperties(testProperty), testImmutables.Get())
	require.Equal(t, id{IDString: metaUtilities.Hash([]string{testProperty.GetFact().GetHashID().String()}...)}, testImmutables.GetHashID())
}
