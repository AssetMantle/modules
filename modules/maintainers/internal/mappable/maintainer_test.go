package mappable

import (
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Maintainer_Methods(t *testing.T) {
	classificationID := base.NewID("classificationID")
	identityID := base.NewID("identityID")
	mutables := base.NewMutables(base.NewProperties(base.NewProperty(base.NewID("ID"), base.NewFact(base.NewStringData("MutableData")))))

	testMaintainerID := key.NewMaintainerID(classificationID, identityID)
	addMaintainer := true
	removeMaintainer := true
	mutateMaintainer := true
	testMaintainer := NewMaintainer(testMaintainerID, mutables, addMaintainer, removeMaintainer, mutateMaintainer).(maintainer)

	require.Equal(t, maintainer{ID: testMaintainerID, MaintainedTraits: mutables, AddMaintainer: addMaintainer, RemoveMaintainer: removeMaintainer, MutateMaintainer: mutateMaintainer}, testMaintainer)
	require.Equal(t, testMaintainerID, testMaintainer.GetID())
	require.Equal(t, classificationID, testMaintainer.GetClassificationID())
	require.Equal(t, identityID, testMaintainer.GetIdentityID())
	require.Equal(t, addMaintainer, testMaintainer.CanAddMaintainer())
	require.Equal(t, mutateMaintainer, testMaintainer.CanMutateMaintainer())
	require.Equal(t, removeMaintainer, testMaintainer.CanRemoveMaintainer())
	require.Equal(t, true, testMaintainer.MaintainsTrait(base.NewID("ID")))
	require.Equal(t, false, testMaintainer.MaintainsTrait(base.NewID("ID2")))
	require.Equal(t, testMaintainerID, testMaintainer.GetKey())
}
