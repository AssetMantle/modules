package key

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func Test_SplitID_Methods(t *testing.T) {

	ownerID := base.NewID("ownerID")
	ownableID := base.NewID("ownableID")

	testSplitID := NewSplitID(ownerID, ownableID).(splitID)
	testSplitID2 := NewSplitID(base.NewID(""), base.NewID("")).(splitID)
	require.Equal(t, strings.Join([]string{ownerID.String(), ownableID.String()}, constants.FirstOrderCompositeIDSeparator), testSplitID.String())
	require.Equal(t, true, testSplitID.Equals(testSplitID))
	require.Equal(t, false, testSplitID.Equals(testSplitID2))
	require.Equal(t, false, testSplitID.IsPartial())
	require.Equal(t, true, testSplitID2.IsPartial())

	require.Equal(t, true, testSplitID.Matches(testSplitID))
	require.Equal(t, false, testSplitID.Matches(testSplitID2))
	require.Equal(t, testSplitID, New(testSplitID))
	require.Equal(t, testSplitID2, New(base.NewID("")))
}
