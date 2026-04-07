// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/schema/ids"
	"reflect"
	"testing"
)

func TestSetPermissions(t *testing.T) {
	type args struct {
		canMintAsset       bool
		canRenumerateAsset bool
		canBurnAsset       bool
	}
	tests := []struct {
		name string
		args args
		want []ids.StringID
	}{
		{"+ve for can Mint", args{true, false, false}, []ids.StringID{constants.CanMintAssetPermission}},
		{"+ve for can Renumerate", args{false, true, false}, []ids.StringID{constants.CanRenumerateAssetPermission}},
		{"+ve for can Burn", args{false, false, true}, []ids.StringID{constants.CanBurnAssetPermission}},
		{"+ve", args{true, true, true}, []ids.StringID{constants.CanMintAssetPermission, constants.CanRenumerateAssetPermission, constants.CanBurnAssetPermission}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetModulePermissions(tt.args.canMintAsset, tt.args.canRenumerateAsset, tt.args.canBurnAsset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetModulePermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
