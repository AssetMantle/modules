package base

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	base2 "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Document(t *testing.T) {
	testProperty := base2.NewProperty(baseIDs.NewID("ID"), baseData.NewHeightData(baseTypes.NewHeight(123)))
	testImmutables := Immutables{base.NewPropertyList(testProperty)}
	testProperties := base.NewPropertyList(testProperty)
	testMutables := Mutables{testProperties}

	creationID := baseIDs.NewID("100")

	takerIDImmutableProperty := base2.NewProperty(constants.TakerIDProperty, baseData.NewStringData("takerIDImmutableProperty"))
	exchangeRateImmutableProperty := base2.NewMetaProperty(constants.ExchangeRateProperty, baseData.NewDecData(sdkTypes.OneDec()))
	creationImmutableProperty := base2.NewMetaProperty(constants.CreationProperty, baseData.NewHeightData(baseTypes.NewHeight(100)))
	expiryImmutableProperty := base2.NewProperty(constants.ExpiryProperty, baseData.NewStringData("expiryImmutableProperty"))
	makerOwnableSplitImmutableProperty := base2.NewProperty(constants.MakerOwnableSplitProperty, baseData.NewStringData("makerOwnableSplitImmutableProperty"))

	immutableProperties := base.NewPropertyList(takerIDImmutableProperty, exchangeRateImmutableProperty.RemoveData(), creationImmutableProperty.RemoveData(), expiryImmutableProperty, makerOwnableSplitImmutableProperty)

	testDocument := Document{ID: creationID, Immutables: Immutables{PropertyList: immutableProperties}, Mutables: Mutables{Properties: base.NewPropertyList()}}
	testDocument1 := Document{ID: creationID, Immutables: testImmutables, Mutables: testMutables}

	require.Equal(t, Document{ID: creationID, Immutables: testImmutables, Mutables: testMutables}, testDocument1)
	require.Equal(t, Document{ID: creationID, Immutables: Immutables{PropertyList: immutableProperties}, Mutables: Mutables{Properties: base.NewPropertyList()}}, testDocument)
}
