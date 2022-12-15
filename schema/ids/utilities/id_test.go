// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"reflect"
	"testing"
)

func TestJoinIDStrings(t *testing.T) {
	type args struct {
		idStrings []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"+ve empty", args{[]string{}}, ""},
		{"+ve single", args{[]string{"a"}}, "a"},
		{"+ve double", args{[]string{"a", "b"}}, "a.b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinIDStrings(tt.args.idStrings...); got != tt.want {
				t.Errorf("JoinIDStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitCompositeIDString(t *testing.T) {
	type args struct {
		idString string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"+ve empty", args{""}, []string{""}},
		{"+ve single", args{"a"}, []string{"a"}},
		{"+ve double", args{"a.b"}, []string{"a", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitCompositeIDString(tt.args.idString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitCompositeIDString() = %v, want %v", got, tt.want)
			}
		})
	}
}
