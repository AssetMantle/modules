// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/ids/constansts"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	schemaProperties "github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"
)

var (
	testImmutables       = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")), baseProperties.NewMetaProperty(constants.CreationHeightProperty.GetKey(), baseData.NewHeightData(base.NewHeight(1))), baseProperties.NewMetaProperty(constants.ExchangeRateProperty.GetKey(), baseData.NewDecData(sdkTypes.NewDec(int64(10)))), baseProperties.NewMetaProperty(constants.MakerOwnableIDProperty.GetKey(), baseData.NewIDData(baseIDs.NewOwnableID(baseIDs.NewStringID("MakerOwnableID")))), baseProperties.NewMetaProperty(constants.TakerOwnableIDProperty.GetKey(), baseData.NewIDData(baseIDs.NewOwnableID(baseIDs.NewStringID("TakerOwnableID")))), baseProperties.NewMetaProperty(constants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(base.NewHeight(100))), baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(sdkTypes.NewDec(int64(10))))))
	testMutables         = baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData")), baseProperties.NewMetaProperty(constants.TakerIDProperty.GetKey(), baseData.NewIDData(baseIDs.PrototypeIdentityID())), baseProperties.NewMetaProperty(constants.MakerIDProperty.GetKey(), baseData.NewIDData(baseIDs.PrototypeIdentityID()))))
	testPermissionsList  = baseLists.NewIDList([]ids.ID{constansts.Mint, constansts.Add, constansts.Remove, constansts.Mutate, constansts.Renumerate, constansts.Burn}...)
	testClassificationID = baseIDs.NewClassificationID(immutables, mutables)
	testIdentityID       = baseIDs.NewIdentityID(testClassificationID, testImmutables)
	testMaintainer       = NewMaintainer(testIdentityID, testClassificationID, testMutables.GetMutablePropertyList().GetPropertyIDList(), testPermissionsList)
)

func TestNewMaintainer(t *testing.T) {
	type args struct {
		identityID                 ids.IdentityID
		maintainedClassificationID ids.ClassificationID
		maintainedPropertyIDList   lists.IDList
		permissions                lists.IDList
	}
	tests := []struct {
		name string
		args args
		want documents.Maintainer
	}{
		{"+ve", args{testIdentityID, testClassificationID, testMutables.GetMutablePropertyList().GetPropertyIDList(), testPermissionsList}, maintainer{NewDocument(constansts.MaintainerClassificationID,
			baseQualified.NewImmutables(baseLists.NewPropertyList(
				baseProperties.NewMetaProperty(constants.IdentityIDProperty.GetKey(), baseData.NewIDData(testIdentityID)),
				baseProperties.NewMetaProperty(constants.MaintainedClassificationIDProperty.GetKey(), baseData.NewIDData(testClassificationID)),
			)),
			baseQualified.NewMutables(baseLists.NewPropertyList(
				baseProperties.NewMetaProperty(constants.MaintainedPropertiesProperty.GetKey(), baseData.NewListData(idListToDataList(testMutables.GetMutablePropertyList().GetPropertyIDList()))),
				baseProperties.NewMetaProperty(constants.PermissionsProperty.GetKey(), baseData.NewListData(idListToDataList(testPermissionsList))),
			)),
		)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMaintainer(tt.args.identityID, tt.args.maintainedClassificationID, tt.args.maintainedPropertyIDList, tt.args.permissions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMaintainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idListToDataList(t *testing.T) {
	type args struct {
		idList lists.IDList
	}
	tests := []struct {
		name string
		args args
		want lists.DataList
	}{
		{"+ve", args{testPermissionsList}, baseLists.NewDataList(baseData.NewIDData(constansts.Mint), baseData.NewIDData(constansts.Add), baseData.NewIDData(constansts.Remove), baseData.NewIDData(constansts.Mutate), baseData.NewIDData(constansts.Renumerate), baseData.NewIDData(constansts.Burn))},
		{"+ve with nil", args{baseLists.NewIDList()}, baseLists.NewDataList()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := idListToDataList(tt.args.idList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("idListToDataList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maintainer_CanAddMaintainer(t *testing.T) {
	testMaintainer2 := NewMaintainer(testIdentityID, testClassificationID, testMutables.GetMutablePropertyList().GetPropertyIDList(), baseLists.NewIDList())
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{testMaintainer}, true},
		{"+ve", fields{testMaintainer2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maintainer := maintainer{
				Document: tt.fields.Document,
			}
			if got := maintainer.CanAddMaintainer(); got != tt.want {
				t.Errorf("CanAddMaintainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maintainer_CanBurnAsset(t *testing.T) {
	testMaintainer2 := NewMaintainer(testIdentityID, testClassificationID, testMutables.GetMutablePropertyList().GetPropertyIDList(), baseLists.NewIDList())
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{testMaintainer}, true},
		{"+ve", fields{testMaintainer2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maintainer := maintainer{
				Document: tt.fields.Document,
			}
			if got := maintainer.CanBurnAsset(); got != tt.want {
				t.Errorf("CanBurnAsset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maintainer_CanMintAsset(t *testing.T) {
	testMaintainer2 := NewMaintainer(testIdentityID, testClassificationID, testMutables.GetMutablePropertyList().GetPropertyIDList(), baseLists.NewIDList())
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{testMaintainer}, true},
		{"+ve", fields{testMaintainer2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maintainer := maintainer{
				Document: tt.fields.Document,
			}
			if got := maintainer.CanMintAsset(); got != tt.want {
				t.Errorf("CanMintAsset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maintainer_CanMutateMaintainer(t *testing.T) {
	testMaintainer2 := NewMaintainer(testIdentityID, testClassificationID, testMutables.GetMutablePropertyList().GetPropertyIDList(), baseLists.NewIDList())
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{testMaintainer}, true},
		{"+ve", fields{testMaintainer2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maintainer := maintainer{
				Document: tt.fields.Document,
			}
			if got := maintainer.CanMutateMaintainer(); got != tt.want {
				t.Errorf("CanMutateMaintainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maintainer_CanRemoveMaintainer(t *testing.T) {
	testMaintainer2 := NewMaintainer(testIdentityID, testClassificationID, testMutables.GetMutablePropertyList().GetPropertyIDList(), baseLists.NewIDList())
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{testMaintainer}, true},
		{"+ve", fields{testMaintainer2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maintainer := maintainer{
				Document: tt.fields.Document,
			}
			if got := maintainer.CanRemoveMaintainer(); got != tt.want {
				t.Errorf("CanRemoveMaintainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maintainer_CanRenumerateAsset(t *testing.T) {
	testMaintainer2 := NewMaintainer(testIdentityID, testClassificationID, testMutables.GetMutablePropertyList().GetPropertyIDList(), baseLists.NewIDList())
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{testMaintainer}, true},
		{"+ve", fields{testMaintainer2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maintainer := maintainer{
				Document: tt.fields.Document,
			}
			if got := maintainer.CanRenumerateAsset(); got != tt.want {
				t.Errorf("CanRenumerateAsset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maintainer_GetIdentityID(t *testing.T) {
	testMaintainer2 := NewMaintainer(baseIDs.PrototypeIdentityID(), testClassificationID, testMutables.GetMutablePropertyList().GetPropertyIDList(), baseLists.NewIDList())
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.IdentityID
	}{
		{"+ve", fields{testMaintainer}, testIdentityID},
		{"+ve", fields{testMaintainer2}, baseIDs.PrototypeIdentityID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maintainer := maintainer{
				Document: tt.fields.Document,
			}
			if got := maintainer.GetIdentityID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIdentityID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maintainer_GetMaintainedClassificationID(t *testing.T) {
	testMaintainer2 := NewMaintainer(testIdentityID, baseIDs.PrototypeClassificationID(), testMutables.GetMutablePropertyList().GetPropertyIDList(), baseLists.NewIDList())
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ClassificationID
	}{
		{"+ve", fields{testMaintainer}, testClassificationID},
		{"+ve", fields{testMaintainer2}, baseIDs.PrototypeClassificationID()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maintainer := maintainer{
				Document: tt.fields.Document,
			}
			if got := maintainer.GetMaintainedClassificationID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMaintainedClassificationID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maintainer_GetMaintainedProperties(t *testing.T) {
	testMaintainer2 := NewMaintainer(testIdentityID, testClassificationID, baseQualified.NewMutables(baseLists.NewPropertyList()).GetMutablePropertyList().GetPropertyIDList(), baseLists.NewIDList())
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   data.ListData
	}{
		{"+ve", fields{testMaintainer}, testMaintainer.GetProperty(constants.MaintainedPropertiesProperty.GetID()).(schemaProperties.MetaProperty).GetData().(data.ListData)},
		{"+ve", fields{testMaintainer2}, constants.MaintainedPropertiesProperty.GetData().(data.ListData)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maintainer := maintainer{
				Document: tt.fields.Document,
			}
			if got := maintainer.GetMaintainedProperties(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMaintainedProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maintainer_GetPermissions(t *testing.T) {
	testMaintainer2 := NewMaintainer(testIdentityID, testClassificationID, testMutables.GetMutablePropertyList().GetPropertyIDList(), baseLists.NewIDList())
	type fields struct {
		Document documents.Document
	}
	tests := []struct {
		name   string
		fields fields
		want   data.ListData
	}{
		{"+ve", fields{testMaintainer}, testMaintainer.GetProperty(constants.PermissionsProperty.GetID()).(schemaProperties.MetaProperty).GetData().(data.ListData)},
		{"+ve", fields{testMaintainer2}, constants.PermissionsProperty.GetData().(data.ListData)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maintainer := maintainer{
				Document: tt.fields.Document,
			}
			if got := maintainer.GetPermissions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maintainer_MaintainsProperty(t *testing.T) {
	type fields struct {
		Document documents.Document
	}
	type args struct {
		propertyID ids.PropertyID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"+ve", fields{testMaintainer}, args{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData")).GetID()}, true},
		{"+ve", fields{testMaintainer}, args{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData("MutableData")).GetID()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maintainer := maintainer{
				Document: tt.fields.Document,
			}
			if got := maintainer.MaintainsProperty(tt.args.propertyID); got != tt.want {
				t.Errorf("MaintainsProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}
