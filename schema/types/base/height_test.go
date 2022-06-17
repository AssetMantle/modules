// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/types"
	"reflect"
	"testing"
)

func TestNewHeight(t *testing.T) {

	h := height{Value: 10}

	type args struct {
		value int64
	}
	tests := []struct {
		name string
		args args
		want types.Height
	}{
		{"Test for New Height", args{10}, h},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHeight(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_height_Compare(t *testing.T) {

	type fields struct {
		Value int64
	}
	type args struct {
		compareHeight types.Height
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"Test for Greater case", fields{10}, args{height{Value: 12}}, -1},
		{"Test for Lower case", fields{10}, args{height{Value: 8}}, 1},
		{"Test for Equal case", fields{10}, args{height{Value: 10}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			height := height{
				Value: tt.fields.Value,
			}
			if got := height.Compare(tt.args.compareHeight); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_height_Get(t *testing.T) {
	type fields struct {
		Value int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{"Test for Get Height", fields{10}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			height := height{
				Value: tt.fields.Value,
			}
			if got := height.Get(); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
