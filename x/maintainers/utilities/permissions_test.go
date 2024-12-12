// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/lists/base"

	"github.com/AssetMantle/modules/x/maintainers/constants"
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
		{"+ve for can Add", args{false, false, false, true, false, false}, idList.Add(constants.CanAddMaintainerPermission)},
		{"+ve for can remove", args{false, false, false, false, true, false}, idList.Add(constants.CanRemoveMaintainerPermission)},
		{"+ve for can mutate", args{false, false, false, false, false, true}, idList.Add(constants.CanMutateMaintainerPermission)},
		{"+ve", args{true, true, true, true, true, true}, idList.Add([]ids.ID{constants.CanAddMaintainerPermission, constants.CanRemoveMaintainerPermission, constants.CanMutateMaintainerPermission}...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetModulePermissions(tt.args.add, tt.args.remove, tt.args.mutate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetModulePermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
