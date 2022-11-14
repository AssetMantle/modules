// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/constansts"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/lists/base"
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
		// TODO: Add test cases.
		{"+ve for can Mint", args{true, false, false, false, false, false}, idList.Add(constansts.Mint)},
		{"+ve for can Burn", args{false, true, false, false, false, false}, idList.Add(constansts.Burn)},
		{"+ve for can Remunerate", args{false, false, true, false, false, false}, idList.Add(constansts.Renumerate)},
		{"+ve for can Add", args{false, false, false, true, false, false}, idList.Add(constansts.Add)},
		{"+ve for can remove", args{false, false, false, false, true, false}, idList.Add(constansts.Remove)},
		{"+ve for can mutate", args{false, false, false, false, false, true}, idList.Add(constansts.Mutate)},
		{"+ve", args{true, true, true, true, true, true}, idList.Add([]ids.ID{constansts.Mint, constansts.Burn, constansts.Renumerate, constansts.Add, constansts.Remove, constansts.Mutate}...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetPermissions(tt.args.mint, tt.args.burn, tt.args.renumerate, tt.args.add, tt.args.remove, tt.args.mutate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
