// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"context"
	"fmt"
	"reflect"
)

type TransactionKeeper interface {
	Transact(context.Context, Message) (TransactionResponse, error)
	Keeper
}

func PanicOnUninitializedTransactionKeeperFields[T TransactionKeeper](transactionKeeper T) {
	elem := reflect.ValueOf(&transactionKeeper).Elem()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		if field.IsNil() {
			panic(fmt.Sprintf("field %s of %s uninitialized", elem.Type().Field(i).Name, elem.Type().PkgPath()))
		}
	}
}
