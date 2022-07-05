// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestGenerateRandomData(t *testing.T) {
	type args struct {
		r *rand.Rand
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		// TODO: check for nil case
		{"+ve case", args{rand.New(rand.NewSource(7))}, false},
		{"nil case", args{nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("GenerateRandomData() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := GenerateRandomData(tt.args.r); reflect.TypeOf(got).String() != "base.idData" {
				t.Errorf("GenerateRandomData() = %v, want base.idData", got)
			}
		})
	}
}
