// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/lists/base"

	"github.com/AssetMantle/modules/x/maintainers/constants"
)

func TestSetPermissions(t *testing.T) {
	type args struct {
		canAddMaintainer    bool
		canMutateMaintainer bool
		canRemoveMaintainer bool
	}
	tests := []struct {
		name string
		args args
		want lists.IDList
	}{
		{"+ve for can Add", args{true, false, false}, base.NewIDList().Add(constants.CanAddMaintainerPermission)},
		{"+ve for can Mutate", args{false, true, false}, base.NewIDList().Add(constants.CanMutateMaintainerPermission)},
		{"+ve for can Remove", args{false, false, true}, base.NewIDList().Add(constants.CanRemoveMaintainerPermission)},
		{"+ve", args{true, true, true}, base.NewIDList().Add(constants.CanAddMaintainerPermission).Add(constants.CanMutateMaintainerPermission).Add(constants.CanRemoveMaintainerPermission)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetModulePermissions(tt.args.canAddMaintainer, tt.args.canMutateMaintainer, tt.args.canRemoveMaintainer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetModulePermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
