// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	base2 "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func TestDocument_GetClassificationID(t *testing.T) {
	type fields struct {
		ID               ids.ID
		ClassificationID ids.ID
		Immutables       Immutables
		Mutables         Mutables
	}

	creationID := baseIDs.NewID("100")
	classificationID := baseIDs.NewID("c100")

	takerIDImmutableProperty := base2.NewProperty(constants.TakerIDProperty.GetKey(), baseData.NewStringData("takerIDImmutableProperty"))
	exchangeRateImmutableProperty := base2.NewMetaProperty(constants.ExchangeRateProperty.GetKey(), baseData.NewDecData(sdkTypes.OneDec()))
	creationImmutableProperty := base2.NewMetaProperty(constants.CreationProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(100)))
	expiryImmutableProperty := base2.NewProperty(constants.ExpiryProperty.GetKey(), baseData.NewStringData("expiryImmutableProperty"))
	makerOwnableSplitImmutableProperty := base2.NewProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewStringData("makerOwnableSplitImmutableProperty"))

	immutableProperties := base.NewPropertyList(takerIDImmutableProperty, exchangeRateImmutableProperty.RemoveData(), creationImmutableProperty.RemoveData(), expiryImmutableProperty, makerOwnableSplitImmutableProperty)

	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{

		{"Test1", fields{ID: creationID, ClassificationID: classificationID, Immutables: Immutables{PropertyList: immutableProperties}, Mutables: Mutables{Properties: base.NewPropertyList()}}, classificationID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := Document{
				ID:         tt.fields.ID,
				Immutables: tt.fields.Immutables,
				Mutables:   tt.fields.Mutables,
			}
			if got := document.GetClassificationID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClassificationID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocument_GetID(t *testing.T) {
	creationID := baseIDs.NewID("100")
	classificationID := baseIDs.NewID("c100")

	takerIDImmutableProperty := base2.NewProperty(constants.TakerIDProperty.GetKey(), baseData.NewStringData("takerIDImmutableProperty"))
	exchangeRateImmutableProperty := base2.NewMetaProperty(constants.ExchangeRateProperty.GetKey(), baseData.NewDecData(sdkTypes.OneDec()))
	creationImmutableProperty := base2.NewMetaProperty(constants.CreationProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(100)))
	expiryImmutableProperty := base2.NewProperty(constants.ExpiryProperty.GetKey(), baseData.NewStringData("expiryImmutableProperty"))
	makerOwnableSplitImmutableProperty := base2.NewProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewStringData("makerOwnableSplitImmutableProperty"))

	immutableProperties := base.NewPropertyList(takerIDImmutableProperty, exchangeRateImmutableProperty.RemoveData(), creationImmutableProperty.RemoveData(), expiryImmutableProperty, makerOwnableSplitImmutableProperty)
	type fields struct {
		ID               ids.ID
		ClassificationID ids.ID
		Immutables       Immutables
		Mutables         Mutables
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		{"Test for GetID", fields{ID: creationID, ClassificationID: classificationID, Immutables: Immutables{PropertyList: immutableProperties}, Mutables: Mutables{Properties: base.NewPropertyList()}}, creationID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := Document{
				ID:         tt.fields.ID,
				Immutables: tt.fields.Immutables,
				Mutables:   tt.fields.Mutables,
			}
			if got := document.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocument_GetProperty(t *testing.T) {

	type fields struct {
		ID               ids.ID
		ClassificationID ids.ID
		Immutables       Immutables
		Mutables         Mutables
	}
	type args struct {
		propertyID ids.PropertyID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   properties.Property
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := Document{
				ID:         tt.fields.ID,
				Immutables: tt.fields.Immutables,
				Mutables:   tt.fields.Mutables,
			}
			if got := document.GetProperty(tt.args.propertyID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocument_Mutate(t *testing.T) {

	type fields struct {
		ID               ids.ID
		ClassificationID ids.ID
		Immutables       Immutables
		Mutables         Mutables
	}
	type args struct {
		propertyList []properties.Property
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   qualified.Document
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			document := Document{
				ID:         tt.fields.ID,
				Immutables: tt.fields.Immutables,
				Mutables:   tt.fields.Mutables,
			}
			if got := document.Mutate(tt.args.propertyList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}
