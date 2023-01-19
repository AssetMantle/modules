// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/modules/splits/internal/transactions/send"
	"github.com/AssetMantle/modules/modules/splits/internal/transactions/unwrap"
	"github.com/AssetMantle/modules/modules/splits/internal/transactions/wrap"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Transactions
	}{
		{"+ve", baseHelpers.NewTransactions(
			send.Transaction,
			unwrap.Transaction,
			wrap.Transaction,
		)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
