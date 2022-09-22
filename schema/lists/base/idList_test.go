// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/traits"
)

func Test_idList_Add(t *testing.T) {
	type fields struct {
		List lists.List
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
		// TODO: Add test cases.
		{"+ve for nil", fields{NewList()}, args{[]ids.ID{base.NewID("ID")}}, idList{NewList(idsToListables([]ids.ID{base.NewID("ID")}...)...)}},                                                             // TODO: panic for nil
		{"+ve", fields{NewList(idsToListables([]ids.ID{base.NewID("ID")}...)...)}, args{[]ids.ID{base.NewID("ID1")}}, idList{NewList(idsToListables([]ids.ID{base.NewID("ID"), base.NewID("ID1")}...)...)}}, // TODO: report
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idList := idList{
				List: tt.fields.List,
			}
			if got := idList.Add(tt.args.ids...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idList_GetList(t *testing.T) {
	type fields struct {
		List lists.List
	}
	tests := []struct {
		name   string
		fields fields
		want   []ids.ID
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{NewList()}, []ids.ID{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idList := idList{
				List: tt.fields.List,
			}
			if got := idList.GetList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idList_Remove(t *testing.T) {
	type fields struct {
		List lists.List
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
		// TODO: Add test cases.
		{"+ve with no removal", fields{NewList(idsToListables(base.NewID("ID1"), base.NewID("ID2"), base.NewID("ID3"))...)}, args{}, idList{NewList(idsToListables(base.NewID("ID1"), base.NewID("ID2"), base.NewID("ID3"))...)}},
		{"+ve with removal", fields{NewList(idsToListables(base.NewID("ID1"), base.NewID("ID2"), base.NewID("ID3"))...)}, args{[]ids.ID{base.NewID("ID3")}}, idList{NewList(idsToListables(base.NewID("ID1"), base.NewID("ID2"))...)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idList := idList{
				List: tt.fields.List,
			}
			if got := idList.Remove(tt.args.ids...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idList_Search(t *testing.T) {
	type fields struct {
		List lists.List
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
		// TODO: Add test cases.
		{"+ve with nil", fields{NewList(idsToListables([]ids.ID{}...)...)}, args{base.NewID("ID")}, 0, false}, // TODO report issue
		{"+ve", fields{NewList(idsToListables([]ids.ID{base.NewID("ID")}...)...)}, args{base.NewID("ID")}, 0, true},
		{"+ve with no entry", fields{NewList(idsToListables([]ids.ID{base.NewID("ID")}...)...)}, args{base.NewID("ID1")}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idList := idList{
				List: tt.fields.List,
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

func Test_idsToListables(t *testing.T) {
	type args struct {
		ids []ids.ID
	}
	tests := []struct {
		name string
		args args
		want []traits.Listable
	}{
		// TODO: Add test cases.
		{"+ve with nil", args{}, []traits.Listable{}},
		{"+ve", args{[]ids.ID{base.NewID("ID")}}, []traits.Listable{base.NewID("ID")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := idsToListables(tt.args.ids...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("idsToListables() = %v, want %v", got, tt.want)
			}
		})
	}
}
