// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"strings"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/traits"
)

type listID struct {
	lists.IDList `json:"idList"`
}

var _ ids.ListID = (*listID)(nil)

func (listID listID) String() string {
	idStringList := make([]string, listID.Size())

	for i, id := range listID.IDList.Get() {
		idStringList[i] = id.String()
	}

	return strings.Join(idStringList, constants.ListDataStringSeparator)
}
func (listID listID) Bytes() []byte {
	var byteList []byte

	for _, id := range listID.IDList.Get() {
		byteList = append(byteList, id.Bytes()...)
	}

	return byteList
}
func (listID listID) Compare(listable traits.Listable) int {
	return bytes.Compare(listID.Bytes(), listIDFromInterface(listable).Bytes())
}

func listIDFromInterface(i interface{}) listID {
	// TODO implement
	panic("")
}
