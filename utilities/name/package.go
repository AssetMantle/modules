// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package name

import (
	"reflect"
	"strings"
)

func GetPackageName(i interface{}) string {
	packagePath := strings.Split(reflect.TypeOf(i).PkgPath(), "/")
	if len(packagePath) > 0 {
		return packagePath[len(packagePath)-1]
	}

	return ""
}

func GetSuperPackageName(i interface{}) string {
	packagePath := strings.Split(reflect.TypeOf(i).PkgPath(), "/")
	if len(packagePath) > 1 {
		return packagePath[len(packagePath)-2]
	}

	return ""
}
