// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/modules/schema/data"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

var (
	testImmutables        = NewImmutables(NewPropertyList(NewMetaProperty(NewStringID("testImutable"), NewStringData("testImmutable"))))
	testImmutables2       = NewImmutables(NewPropertyList())
	testMutables          = NewMutables(NewPropertyList(NewMetaProperty(NewStringID("testMutable"), NewStringData("testImmutable"))))
	testMutables2         = NewMutables(NewPropertyList())
	testClassificationID  = NewClassificationID(testImmutables, testMutables)
	testClassificationID2 = PrototypeClassificationID()
	testAssetID           = NewAssetID(testClassificationID, testImmutables)
	testAssetID2          = PrototypeAssetID()
)

func TestNewAssetID(t *testing.T) {
	type args struct {
		classificationID ids.ClassificationID
		immutables       qualified.Immutables
	}
	tests := []struct {
		name string
		args args
		want ids.AssetID
	}{
		{"+ve with empty immutable", args{testClassificationID, testImmutables2}, assetID{HashID: GenerateHashID(testClassificationID.Bytes())}},
		{"+ve", args{testClassificationID, testImmutables}, NewAssetID(testClassificationID, testImmutables)},
		{"+ve with empty classificationID", args{testClassificationID2, testImmutables}, assetID{HashID: GenerateHashID(testClassificationID2.Bytes(), testImmutables.GenerateHashID().Bytes())}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewAssetID(tt.args.classificationID, tt.args.immutables), "NewAssetID(%v, %v)", tt.args.classificationID, tt.args.immutables)
		})
	}
}

func TestPrototypeAssetID(t *testing.T) {
	tests := []struct {
		name string
		want ids.AssetID
	}{
		{"+ve", assetID{HashID: PrototypeHashID()}},
		{"+ve", NewAssetID(NewClassificationID(NewImmutables(NewPropertyList()), NewMutables(NewPropertyList())), NewImmutables(NewPropertyList()))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, PrototypeAssetID(), "PrototypeAssetID()")
		})
	}
}

func TestReadAssetID(t *testing.T) {
	type args struct {
		assetIDString string
	}
	tests := []struct {
		name    string
		args    args
		want    ids.AssetID
		wantErr assert.ErrorAssertionFunc
	}{
		{"+ve", args{testAssetID.String()}, testAssetID, assert.NoError},
		{"+ve with empty immutable", args{NewAssetID(testClassificationID, testImmutables2).String()}, NewAssetID(testClassificationID, testImmutables2), assert.NoError},
		{"+ve with empty classificationID", args{NewAssetID(testClassificationID2, testImmutables).String()}, NewAssetID(testClassificationID2, testImmutables), assert.NoError},
		{"-ve with empty String", args{""}, PrototypeAssetID(), assert.NoError},
		{"-ve with nil", args{}, assetID{}, assert.NoError},
		{"+ve with Random String", args{"Random String"}, assetID{}, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadAssetID(tt.args.assetIDString)
			if !tt.wantErr(t, err, fmt.Sprintf("ReadAssetID(%v)", tt.args.assetIDString)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ReadAssetID(%v)", tt.args.assetIDString)
		})
	}
}

func Test_assetIDFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name        string
		args        args
		want        assetID
		shouldPanic bool
	}{
		{"+ve", args{testAssetID}, testAssetID.(assetID), false},
		{"+ve with PrototypeAssetID", args{PrototypeAssetID()}, PrototypeAssetID().(assetID), false},
		{"+ve with nil", args{}, PrototypeAssetID().(assetID), true},
		{"+ve with identityID", args{testIdentityID}, PrototypeAssetID().(assetID), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				assert.Panicsf(t, func() { assetIDFromInterface(tt.args.i) }, "assetIDFromInterface(%v)", tt.args.i)
			} else {
				assert.Equalf(t, tt.want, assetIDFromInterface(tt.args.i), "assetIDFromInterface(%v)", tt.args.i)
			}
		})
	}
}

func Test_assetID_Compare(t *testing.T) {
	type fields struct {
		HashID ids.HashID
	}
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"-ve", fields{GenerateHashID(testClassificationID2.Bytes(), testImmutables2.GenerateHashID().Bytes())}, args{testAssetID}, -1},
		{"+ve", fields{GenerateHashID(testClassificationID.Bytes(), testImmutables.GenerateHashID().Bytes())}, args{testAssetID}, 0},
		{"+ve prototype", fields{PrototypeHashID()}, args{testAssetID2}, 0},
		{"-ve prototype", fields{PrototypeHashID()}, args{testAssetID}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assetID := assetID{
				HashID: tt.fields.HashID,
			}
			assert.Equalf(t, tt.want, assetID.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
		})
	}
}

// Immutable Mocks
type immutables struct {
	lists.PropertyList
}

func (i immutables) GetImmutablePropertyList() lists.PropertyList {
	//if i.PropertyList.GetList() == nil {
	//	return base.NewPropertyList()
	//}

	return i.PropertyList
}

func (i immutables) GenerateHashID() ids.HashID {
	metaList := make([][]byte, len(i.PropertyList.GetList()))

	for i, property := range i.PropertyList.GetList() {
		metaList[i] = property.GetDataID().GetHashID().Bytes()
	}
	return GenerateHashID(metaList...)
}

func NewImmutables(propertyList lists.PropertyList) qualified.Immutables {
	return immutables{
		PropertyList: propertyList,
	}
}

// Mutable Mocks
type mutables struct {
	lists.PropertyList
}

var _ qualified.Mutables = (*mutables)(nil)

func (m mutables) GetMutablePropertyList() lists.PropertyList {
	//TODO implement me
	return m.PropertyList
}

func (m mutables) Mutate(property ...properties.Property) qualified.Mutables {
	//TODO implement me
	panic("implement me")
}

func NewMutables(propertyList lists.PropertyList) qualified.Mutables {
	return mutables{
		PropertyList: propertyList,
	}
}

// PropertyList Mocks
type propertyList struct {
	lists.List
}

func (p propertyList) Sanitize() (lists.PropertyList, error) {
	//TODO implement me
	panic("implement me")
}

func (p propertyList) GetProperty(id ids.PropertyID) properties.Property {
	//TODO implement me
	panic("implement me")
}

func (p propertyList) GetList() []properties.Property {
	Properties := make([]properties.Property, p.Size())
	for i, property := range p.List.Get() {
		Properties[i] = property.(properties.Property)
	}
	return Properties
}

func (p propertyList) GetPropertyIDList() lists.IDList {
	//TODO implement me
	panic("implement me")
}

func (p propertyList) Add(property ...properties.Property) lists.PropertyList {
	//TODO implement me
	panic("implement me")
}

func (p propertyList) Remove(property ...properties.Property) lists.PropertyList {
	//TODO implement me
	panic("implement me")
}

func (p propertyList) Mutate(property ...properties.Property) lists.PropertyList {
	//TODO implement me
	panic("implement me")
}

func (p propertyList) ScrubData() lists.PropertyList {
	//TODO implement me
	panic("implement me")
}

func propertiesToListables(properties ...properties.Property) []traits.Listable {
	listables := make([]traits.Listable, len(properties))
	for i, property := range properties {
		listables[i] = property
	}
	return listables
}

func NewPropertyList(properties ...properties.Property) lists.PropertyList {
	return propertyList{List: NewList(propertiesToListables(properties...)...)}
}

// NewList Mocks
type list []traits.Listable

func (l list) Sanitize() (lists.List, error) {
	//TODO implement me
	panic("implement me")
}

func (l list) Get() []traits.Listable {
	//TODO implement me
	List := make([]traits.Listable, l.Size())
	for i, listable := range l {
		List[i] = listable
	}
	return List
}

func (l list) Size() int {
	return len(l)
}

func (l list) Search(listable traits.Listable) (index int, found bool) {
	//TODO implement me
	panic("implement me")
}

func (l list) Add(listable ...traits.Listable) lists.List {
	//TODO implement me
	panic("implement me")
}

func (l list) Remove(listable ...traits.Listable) lists.List {
	//TODO implement me
	panic("implement me")
}

func (l list) Mutate(listable ...traits.Listable) lists.List {
	//TODO implement me
	panic("implement me")
}

func NewList(listables ...traits.Listable) lists.List {
	list := list(listables)
	sort.Slice(list, func(i, j int) bool {
		return list[i].Compare(list[j]) <= 0
	})

	return list
}

// MetaProperty Mocks
type metaProperty struct {
	ID   ids.PropertyID `json:"id"`
	Data data.Data      `json:"data"`
}

var _ properties.Property = (*metaProperty)(nil)

func (m metaProperty) GetData() data.Data {
	//TODO implement me
	panic("implement me")
}

func (m metaProperty) ScrubData() properties.MesaProperty {
	//TODO implement me
	panic("implement me")
}

func (m metaProperty) GetID() ids.PropertyID {
	//TODO implement me
	return m.ID
	//panic("implement me")
}

func (m metaProperty) GetDataID() ids.DataID {
	//TODO implement me
	return m.Data.GetID()
}

func (m metaProperty) GetKey() ids.StringID {
	//TODO implement me
	panic("implement me")
}

func (m metaProperty) GetType() ids.StringID {
	//TODO implement me
	panic("implement me")
}

func (m metaProperty) IsMeta() bool {
	//TODO implement me
	panic("implement me")
}

func (m metaProperty) Compare(listable traits.Listable) int {
	//TODO implement me
	panic("implement me")
}

func NewMetaProperty(key ids.StringID, data data.Data) properties.MetaProperty {
	if data == nil || key == nil {
		panic(errorConstants.MetaDataError)
	}
	return metaProperty{
		ID:   NewPropertyID(key, data.GetType()),
		Data: data,
	}
}

//StringData Mocks
type stringData struct {
	Value string `json:"value"`
}

func (s stringData) GetID() ids.DataID {
	//TODO implement me
	return NewDataID(s)
}

func (s stringData) String() string {
	//TODO implement me
	panic("implement me")
}

func (s stringData) Bytes() []byte {
	//TODO implement me
	return []byte(s.Value)
}

func (s stringData) GetType() ids.StringID {
	//TODO implement me
	return NewStringID("S")
}

func (s stringData) ZeroValue() data.Data {
	//TODO implement me
	panic("implement me")
}

func (s stringData) GenerateHashID() ids.HashID {
	//TODO implement me
	return GenerateHashID(s.Bytes())
}

func (s stringData) Compare(listable traits.Listable) int {
	//TODO implement me
	panic("implement me")
}

func (s stringData) Get() string {
	//TODO implement me
	panic("implement me")
}

func NewStringData(value string) data.StringData {
	return stringData{
		Value: value,
	}
}
