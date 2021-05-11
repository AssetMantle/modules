package key

import (
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromID(t *testing.T) {

	classificationID := base.NewID("classificationID")
	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("ImmutableData"))))
	assetID := NewAssetID(classificationID, immutableProperties)
	require.Equal(t, assetIDFromInterface(assetID), FromID(assetID))

}

func TestReadClassificationID(t *testing.T) {
	classificationID := base.NewID("classificationID")
	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("ImmutableData"))))
	assetID := NewAssetID(classificationID, immutableProperties)

	require.Equal(t, assetIDFromInterface(assetID).ClassificationID, ReadClassificationID(assetID))

}

func Test_assetIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	var tests []struct {
		name string
		args args
		want assetID
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, assetIDFromInterface(tt.args.i), FromID(base.NewID("")))
		})
	}
}

func Test_readAssetID(t *testing.T) {
	classificationID := base.NewID("")
	immutableProperties := base.NewProperties(base.NewProperty(base.NewID(""), base.NewFact(base.NewStringData(""))))
	tassetID := NewAssetID(classificationID, immutableProperties)
	assetID2 := assetID{ClassificationID: base.NewID(""), HashID: base.NewID("")}

	require.Equal(t, tassetID, assetID2)
}
