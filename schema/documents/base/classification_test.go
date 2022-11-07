package base

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Classification_Methods(t *testing.T) {

	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))

	id := baseIDs.NewClassificationID(immutables, mutables)

	testClassification := NewClassification(immutables, mutables)
	require.Equal(t, classification{NewDocument(id, immutables, mutables)}, testClassification)
	require.Equal(t, immutables, testClassification.GetImmutables())
	require.Equal(t, mutables, testClassification.GetMutables())
}
