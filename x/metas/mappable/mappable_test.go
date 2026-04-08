// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"testing"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/data"
	"github.com/AssetMantle/schema/data/base"
	"github.com/stretchr/testify/assert"
)

func TestNewMappable(t *testing.T) {
	type args struct {
		data data.Data
	}
	tests := []struct {
		name string
		args args
		want helpers.Mappable
	}{
		{"valid", args{base.NewStringData("data")}, &Mappable{base.NewStringData("data").ToAnyData().(*base.AnyData)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMappable(tt.args.data)
			assert.Equal(t, tt.want, got, "NewMappable()")
		})
	}
}

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Mappable
	}{
		{"valid", &Mappable{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Prototype()
			assert.Equal(t, tt.want, got, "Prototype()")
		})
	}
}
