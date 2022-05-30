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
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/utilities/meta"
)

type listDataID struct {
	lists.IDList `json:"idList"`
}

var _ ids.ListDataID = (*listDataID)(nil)

// TODO ambigous implementation, recheck
func (listDataID listDataID) GetHash() types.ID {
	idStringList := make([]string, listDataID.IDList.Size())

	for i, id := range listDataID.IDList.GetList() {
		idStringList[i] = id.String()
	}

	return NewID(meta.Hash(idStringList...))
}
func (listDataID listDataID) String() string {
	idStringList := make([]string, listDataID.IDList.Size())

	for i, id := range listDataID.IDList.GetList() {
		idStringList[i] = id.String()
	}

	return strings.Join(idStringList, constants.ListDataStringSeparator)
}
func (listDataID listDataID) Bytes() []byte {
	var byteList []byte

	for _, id := range listDataID.IDList.GetList() {
		byteList = append(byteList, id.Bytes()...)
	}

	return byteList
}
func (listDataID listDataID) Compare(listable traits.Listable) int {
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
