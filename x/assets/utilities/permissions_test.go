// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/lists/base"
	"reflect"
	"testing"
)

func TestSetPermissions(t *testing.T) {
	idList := base.NewIDList()
	type args struct {
		mint       bool
		burn       bool
		renumerate bool
		add        bool
		remove     bool
		mutate     bool
	}
	tests := []struct {
		name string
		args args
		want lists.IDList
	}{
		{"+ve for can Mint", args{true, false, false, false, false, false}, idList.Add(constants.CanMintAssetPermission)},
		{"+ve for can Burn", args{false, true, false, false, false, false}, idList.Add(constants.CanBurnAssetPermission)},
		{"+ve for can Remunerate", args{false, false, true, false, false, false}, idList.Add(constants.CanRenumerateAssetPermission)},
		{"+ve", args{true, true, true, true, true, true}, idList.Add([]ids.ID{constants.CanMintAssetPermission, constants.CanBurnAssetPermission, constants.CanRenumerateAssetPermission}...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetModulePermissions(tt.args.mint, tt.args.burn, tt.args.renumerate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetModulePermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
