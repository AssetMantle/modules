// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mapper

import (
	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/assets/internal/mappable"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

//func TestPrototype(t *testing.T) {
//	storeKey := sdkTypes.NewKVStoreKey("test")
//	newMapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
//
//	tests := []struct {
//		name string
//		want helpers.Mapper
//	}{
//		// TODO: Add test cases.
//		{"Default Tests", newMapper},
//		//{"dummy tests", baseHelpers.NewMapper(key.Prototype, mappable.Prototype)},
//		//{"Dt1", Prototype()},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Prototype().Initialize(storeKey); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Prototype() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestPrototype(t *testing.T) {
	//storeKey := sdkTypes.NewKVStoreKey("test")
	require.Panics(t, func() {
		require.Equal(t, Prototype(), baseHelpers.NewMapper(key.Prototype, mappable.Prototype))
	})
}
