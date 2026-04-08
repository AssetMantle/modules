// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"cosmossdk.io/math"
	"testing"

	"github.com/AssetMantle/schema/types"
	baseTypes "github.com/AssetMantle/schema/types/base"
"github.com/AssetMantle/modules/helpers"
	"github.com/stretchr/testify/assert"
)

var (
	testRate = math.NewInt(1)
	split    = baseTypes.NewSplit(testRate).(*baseTypes.Split)
)

func TestNewMappable(t *testing.T) {
	type args struct {
		split types.Split
	}
	tests := []struct {
		name string
		args args
		want helpers.Mappable
	}{
		{"valid", args{split}, &Mappable{split}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMappable(tt.args.split)
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
