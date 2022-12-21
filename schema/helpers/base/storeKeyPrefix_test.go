// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/binary"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	assets int8 = iota + 8
	classifications
	identities
	maintainers
	metas
	orders
	splits
)

var (
	immutables             = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables               = baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID   = baseIDs.NewClassificationID(immutables, mutables)
	assetKeyBytes          = baseIDs.NewAssetID(testClassificationID, immutables).Bytes()
	classificationKeyBytes = baseIDs.NewClassificationID(immutables, mutables).Bytes()
	identityID             = baseIDs.NewIdentityID(testClassificationID, immutables)
	identityKeyBytes       = identityID.Bytes()
	maintainerKeyBytes     = baseIDs.NewMaintainerID(testClassificationID, immutables).Bytes()
	metaKeyBytes           = baseIDs.NewDataID(baseData.NewStringData("DataID")).Bytes()
	orderKeyBytes          = baseIDs.NewOrderID(testClassificationID, immutables).Bytes()
	splitsKeyBytes         = baseIDs.NewSplitID(identityID, baseIDs.NewOwnableID(baseIDs.NewStringID("OwnerID"))).Bytes()
)

func TestNewStoreKeyPrefix(t *testing.T) {
	type args struct {
		value int8
	}
	tests := []struct {
		name string
		args args
		want helpers.StoreKeyPrefix
	}{
		{"+ve with assets", args{assets}, storeKeyPrefix(assets)},
		{"+ve with classifications", args{classifications}, storeKeyPrefix(classifications)},
		{"+ve with identities", args{identities}, storeKeyPrefix(identities)},
		{"+ve with maintainers", args{maintainers}, storeKeyPrefix(maintainers)},
		{"+ve with metas", args{metas}, storeKeyPrefix(metas)},
		{"+ve with orders", args{orders}, storeKeyPrefix(orders)},
		{"+ve with splits", args{splits}, storeKeyPrefix(splits)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewStoreKeyPrefix(tt.args.value), "NewStoreKeyPrefix(%v)", tt.args.value)
		})
	}
}

func Test_storeKeyPrefix_GenerateStoreKey(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name           string
		storeKeyPrefix storeKeyPrefix
		args           args
		want           []byte
	}{
		{"+ve for AssetsStoreKeyPrefix", NewStoreKeyPrefix(assets).(storeKeyPrefix), args{assetKeyBytes}, append(getStorePrefixBytes(assets), assetKeyBytes...)},
		{"+ve for ClassificationsStoreKeyPrefix", NewStoreKeyPrefix(classifications).(storeKeyPrefix), args{classificationKeyBytes}, append(getStorePrefixBytes(classifications), classificationKeyBytes...)},
		{"+ve for IdentitiesStoreKeyPrefix", NewStoreKeyPrefix(identities).(storeKeyPrefix), args{identityKeyBytes}, append(getStorePrefixBytes(identities), identityKeyBytes...)},
		{"+ve for MaintainersStoreKeyPrefix", NewStoreKeyPrefix(maintainers).(storeKeyPrefix), args{maintainerKeyBytes}, append(getStorePrefixBytes(maintainers), maintainerKeyBytes...)},
		{"+ve for MetasStoreKeyPrefix", NewStoreKeyPrefix(metas).(storeKeyPrefix), args{metaKeyBytes}, append(getStorePrefixBytes(metas), metaKeyBytes...)},
		{"+ve for OrdersStoreKeyPrefix", NewStoreKeyPrefix(orders).(storeKeyPrefix), args{orderKeyBytes}, append(getStorePrefixBytes(orders), orderKeyBytes...)},
		{"+ve for SplitsStoreKeyPrefix", NewStoreKeyPrefix(splits).(storeKeyPrefix), args{splitsKeyBytes}, append(getStorePrefixBytes(splits), splitsKeyBytes...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.storeKeyPrefix.GenerateStoreKey(tt.args.key), "GenerateStoreKey(%v)", tt.args.key)
		})
	}
}

func getStorePrefixBytes(value int8) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, uint16(value))
	return bytes
}
