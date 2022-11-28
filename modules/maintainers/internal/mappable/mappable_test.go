// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"github.com/AssetMantle/modules/schema/ids/constansts"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/maintainers/internal/key"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

func createTestData(t *testing.T) documents.Maintainer {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testIdentityID := baseIDs.NewIdentityID(testClassificationID, immutables)
	testMaintainer := base.NewMaintainer(testIdentityID, testClassificationID, baseLists.NewIDList(baseIDs.NewStringID("ID2")), baseLists.NewIDList(baseIDs.NewStringID("ID2")))
	return testMaintainer
}

func TestNewMappable(t *testing.T) {
	testMaintainer := createTestData(t)
	type args struct {
		maintainer documents.Maintainer
	}
	tests := []struct {
		name string
		args args
		want helpers.Mappable
	}{
		{"+ve", args{testMaintainer}, mappable{testMaintainer}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMappable(tt.args.maintainer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMappable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Mappable
	}{
		{"+ve", mappable{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mappable_GetKey(t *testing.T) {
	testMaintainer := createTestData(t)
	type fields struct {
		Maintainer documents.Maintainer
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.Key
	}{
		{"+ve", fields{testMaintainer}, key.NewKey(baseIDs.NewMaintainerID(constansts.MaintainerClassificationID,
			baseQualified.NewImmutables(baseLists.NewPropertyList(
				baseProperties.NewMetaProperty(constantProperties.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(testMaintainer.GetMaintainedClassificationID())),
				baseProperties.NewMetaProperty(constantProperties.IdentityIDProperty.GetKey(), baseData.NewIDData(testMaintainer.GetIdentityID())),
			))))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maintainer := mappable{
				Maintainer: tt.fields.Maintainer,
			}
			if got := maintainer.GetKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mappable_RegisterCodec(t *testing.T) {
	testMaintainer := createTestData(t)
	type fields struct {
		Maintainer documents.Maintainer
	}
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testMaintainer}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ma := mappable{
				Maintainer: tt.fields.Maintainer,
			}
			ma.RegisterCodec(tt.args.codec)
		})
	}
}
