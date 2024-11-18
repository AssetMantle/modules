// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"fmt"
	"reflect"
)

type Keeper interface {
	Initialize(Mapper, ParameterManager, []interface{}) Keeper
}

func PanicOnUninitializedKeeperFields[T Keeper](keeper T) {
	elem := reflect.ValueOf(&keeper).Elem()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		if field.IsNil() {
			panic(fmt.Sprintf("field %s of %s uninitialized", elem.Type().Field(i).Name, elem.Type().PkgPath()))
		}
	}
}
