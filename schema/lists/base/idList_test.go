// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
)

func Test_idList_Add(t *testing.T) {
	type fields struct {
		List []*baseIDs.AnyID
	}
	type args struct {
		ids []ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   lists.IDList
	}{
		{"+ve for nil", fields{[]*baseIDs.AnyID{}}, args{[]ids.ID{baseIDs.NewStringID("ID")}}, &IDList{[]*baseIDs.AnyID{baseIDs.NewStringID("ID").ToAnyID().(*baseIDs.AnyID)}}},                                                                                                     // TODO: panic for nil
		{"+ve", fields{[]*baseIDs.AnyID{baseIDs.NewStringID("ID").ToAnyID().(*baseIDs.AnyID)}}, args{[]ids.ID{baseIDs.NewStringID("ID1")}}, &IDList{[]*baseIDs.AnyID{baseIDs.NewStringID("ID").ToAnyID().(*baseIDs.AnyID), baseIDs.NewStringID("ID1").ToAnyID().(*baseIDs.AnyID)}}}, // TODO: report
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idList := &IDList{
				IDList: tt.fields.List,
			}
			if got := idList.Add(tt.args.ids...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idList_GetList(t *testing.T) {
	type fields struct {
		List []*baseIDs.AnyID
	}
	tests := []struct {
		name   string
		fields fields
		want   []ids.AnyID
	}{
		{"+ve with nil", fields{[]*baseIDs.AnyID{}}, []ids.AnyID{}},
		{"+ve", fields{[]*baseIDs.AnyID{baseIDs.NewStringID("Data").ToAnyID().(*baseIDs.AnyID)}}, []ids.AnyID{baseIDs.NewStringID("Data").(ids.ID).ToAnyID()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idList := IDList{
				IDList: tt.fields.List,
			}
			if got := idList.GetList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idList_Remove(t *testing.T) {
	type fields struct {
		List []*baseIDs.AnyID
	}
	type args struct {
		ids []ids.ID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   lists.IDList
	}{
		{"-ve with no removal", fields{[]*baseIDs.AnyID{baseIDs.NewStringID("ID1").ToAnyID().(*baseIDs.AnyID), baseIDs.NewStringID("ID2").ToAnyID().(*baseIDs.AnyID), baseIDs.NewStringID("ID3").ToAnyID().(*baseIDs.AnyID)}}, args{}, &IDList{[]*baseIDs.AnyID{baseIDs.NewStringID("ID1").ToAnyID().(*baseIDs.AnyID), baseIDs.NewStringID("ID2").ToAnyID().(*baseIDs.AnyID), baseIDs.NewStringID("ID3").ToAnyID().(*baseIDs.AnyID)}}},
		{"+ve with removal", fields{[]*baseIDs.AnyID{baseIDs.NewStringID("ID1").ToAnyID().(*baseIDs.AnyID), baseIDs.NewStringID("ID2").ToAnyID().(*baseIDs.AnyID), baseIDs.NewStringID("ID3").ToAnyID().(*baseIDs.AnyID)}}, args{[]ids.ID{baseIDs.NewStringID("ID3")}}, &IDList{[]*baseIDs.AnyID{baseIDs.NewStringID("ID1").ToAnyID().(*baseIDs.AnyID), baseIDs.NewStringID("ID2").ToAnyID().(*baseIDs.AnyID)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idList := IDList{
				IDList: tt.fields.List,
			}
			if got := idList.Remove(tt.args.ids...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idList_Search(t *testing.T) {
	type fields struct {
		List []*baseIDs.AnyID
	}
	type args struct {
		id ids.ID
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantIndex int
		wantFound bool
	}{
		{"+ve with nil", fields{[]*baseIDs.AnyID{}}, args{baseIDs.NewStringID("ID")}, 0, false}, // TODO report issue
		{"+ve", fields{[]*baseIDs.AnyID{baseIDs.NewStringID("ID").ToAnyID().(*baseIDs.AnyID)}}, args{baseIDs.NewStringID("ID")}, 0, true},
		{"-ve with no entry", fields{[]*baseIDs.AnyID{baseIDs.NewStringID("ID").ToAnyID().(*baseIDs.AnyID)}}, args{baseIDs.NewStringID("ID1")}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idList := IDList{
				IDList: tt.fields.List,
			}
			gotIndex, gotFound := idList.Search(tt.args.id)
			if gotIndex != tt.wantIndex {
				t.Errorf("Search() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Search() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}
