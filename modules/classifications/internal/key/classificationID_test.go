package key

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func Test_ClassificationID_Methods(t *testing.T) {
	chainID := base.NewID("chainID")
	immutables := base.NewImmutables(base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("ImmutableData")))))
	mutables := base.NewMutables(base.NewProperties(base.NewProperty(base.NewID("ID2"), base.NewFact(base.NewStringData("MutableData")))))

	testClassificationID := NewClassificationID(chainID, immutables, mutables).(classificationID)
	require.Equal(t, classificationID{ChainID: chainID, HashID: base.NewID(metaUtilities.Hash(metaUtilities.Hash("ID1"), metaUtilities.Hash("ID2"), immutables.GetHashID().String()))}, testClassificationID)
	require.Equal(t, strings.Join([]string{chainID.String(), base.NewID(metaUtilities.Hash(metaUtilities.Hash("ID1"), metaUtilities.Hash("ID2"), immutables.GetHashID().String())).String()}, constants.IDSeparator), testClassificationID.String())
	require.Equal(t, false, testClassificationID.Matches(classificationID{ChainID: base.NewID("chainID"), HashID: base.NewID("hashID")}))
	require.Equal(t, false, testClassificationID.Equals(base.NewID("id")))
	require.Equal(t, true, testClassificationID.Equals(testClassificationID))
	require.Equal(t, false, testClassificationID.IsPartial())
	require.Equal(t, true, classificationID{ChainID: chainID, HashID: base.NewID("")}.IsPartial())
	require.Equal(t, testClassificationID, New(testClassificationID))
	require.Equal(t, classificationID{ChainID: base.NewID(""), HashID: base.NewID("")}, New(base.NewID("tempID")))
	require.Equal(t, testClassificationID, readClassificationID(testClassificationID.String()))
}
