// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"testing"

	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/lists/base"

	"github.com/AssetMantle/modules/x/maintainers/constants"
	"github.com/stretchr/testify/assert"
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
		{"can add permission", args{true, false, false}, base.NewIDList().Add(constants.CanAddMaintainerPermission)},
		{"can mutate permission", args{false, true, false}, base.NewIDList().Add(constants.CanMutateMaintainerPermission)},
		{"can remove permission", args{false, false, true}, base.NewIDList().Add(constants.CanRemoveMaintainerPermission)},
		{"valid", args{true, true, true}, base.NewIDList().Add(constants.CanAddMaintainerPermission).Add(constants.CanMutateMaintainerPermission).Add(constants.CanRemoveMaintainerPermission)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetModulePermissions(tt.args.canAddMaintainer, tt.args.canMutateMaintainer, tt.args.canRemoveMaintainer)
			assert.Equal(t, tt.want, got, "SetModulePermissions()")
		})
	}
}
