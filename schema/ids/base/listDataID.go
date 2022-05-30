// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"strings"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/traits"
)

type listDataID struct {
	lists.IDList `json:"idList"`
}

var _ ids.ListDataID = (*listDataID)(nil)

func (listID listDataID) String() string {
	idStringList := make([]string, listID.Size())

	for i, id := range listID.IDList.GetList() {
		idStringList[i] = id.String()
	}

	return strings.Join(idStringList, constants.ListDataStringSeparator)
}
func (listID listDataID) Bytes() []byte {
	var byteList []byte

	for _, id := range listID.IDList.GetList() {
		byteList = append(byteList, id.Bytes()...)
	}

	return byteList
}
func (listID listDataID) Compare(listable capabilities.Listable) int {
	if listID, err := listDataIDFromInterface(listable); err != nil {
		panic(err)
	} else {
		return bytes.Compare(listID.Bytes(), listID.Bytes())
	}
}

func listDataIDFromInterface(i interface{}) (listDataID, error) {
	switch value := i.(type) {
	case listDataID:
		return value, nil
	default:
		return listDataID{}, errors.MetaDataError
	}
}
