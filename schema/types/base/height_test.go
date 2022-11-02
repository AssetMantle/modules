// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math"
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/schema/types"
)

func TestNewHeight(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name string
		args args
		want types.Height
	}{
		// Fails in case of overflow conditions
		{"Testing with value 10", args{
			10,
		}, height{Value: 10}},
		{"Testing with value -10", args{
			-10,
		}, height{Value: -10}},
		{"Testing with value math.MaxInt64", args{
			math.MaxInt64,
		}, height{Value: math.MaxInt64}},
		{"Testing with value math.MaxInt64", args{
			math.MinInt64,
		}, height{Value: math.MinInt64}},
		{"Testing with value math.MaxInt64", args{
			math.MinInt64,
		}, height{Value: math.MinInt64}},
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
		{
			"compareHeight > height", fields{
				10,
			}, args{height{11}}, -1,
		},
		{
			"compareHeight < height", fields{
				10,
			}, args{height{9}}, 1,
		},
		{
			"compareHeight === height", fields{
				10,
			}, args{height{10}}, 0,
		},
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
		{
			"Testing with 100", fields{
				100,
			}, 100,
		},
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
