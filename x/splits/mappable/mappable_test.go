// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/types"
	baseTypes "github.com/AssetMantle/schema/types/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

var (
	testRate = sdkTypes.NewInt(1)
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
		{"+ve", args{split}, &Mappable{split}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMappable(tt.args.split); !reflect.DeepEqual(got, tt.want) {
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
		{"+ve", &Mappable{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
