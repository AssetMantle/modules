// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/schema/ids"
	"testing"
	"github.com/stretchr/testify/assert"
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
		{"can mint permission", args{true, false, false}, []ids.StringID{constants.CanMintAssetPermission}},
		{"can renumerate permission", args{false, true, false}, []ids.StringID{constants.CanRenumerateAssetPermission}},
		{"can burn permission", args{false, false, true}, []ids.StringID{constants.CanBurnAssetPermission}},
		{"valid", args{true, true, true}, []ids.StringID{constants.CanMintAssetPermission, constants.CanRenumerateAssetPermission, constants.CanBurnAssetPermission}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetModulePermissions(tt.args.canMintAsset, tt.args.canRenumerateAsset, tt.args.canBurnAsset)
			assert.Equal(t, tt.want, got, "SetModulePermissions()")
		})
	}
}
