// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/x/ids"
	"github.com/AssetMantle/schema/x/lists"
	"github.com/AssetMantle/schema/x/lists/base"

	"github.com/AssetMantle/modules/x/maintainers/internal/module"
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
		{"+ve for can Mint", args{true, false, false, false, false, false}, idList.Add(module.Mint)},
		{"+ve for can Burn", args{false, true, false, false, false, false}, idList.Add(module.Burn)},
		{"+ve for can Remunerate", args{false, false, true, false, false, false}, idList.Add(module.Renumerate)},
		{"+ve for can Add", args{false, false, false, true, false, false}, idList.Add(module.Add)},
		{"+ve for can remove", args{false, false, false, false, true, false}, idList.Add(module.Remove)},
		{"+ve for can mutate", args{false, false, false, false, false, true}, idList.Add(module.Mutate)},
		{"+ve", args{true, true, true, true, true, true}, idList.Add([]ids.ID{module.Mint, module.Burn, module.Renumerate, module.Add, module.Remove, module.Mutate}...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetPermissions(tt.args.mint, tt.args.burn, tt.args.renumerate, tt.args.add, tt.args.remove, tt.args.mutate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
