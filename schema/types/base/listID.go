// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"strings"

	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/lists"
	"github.com/persistenceOne/persistenceSDK/schema/traits"

	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type listID struct {
	IDList lists.IDList `json:"idList"`
}

var _ types.ID = (*listID)(nil)
var _ types.List = (*listID)(nil)

func (listID listID) String() string {
	idStringList := make([]string, listID.Size())

	for i, id := range listID.IDList.GetList() {
		idStringList[i] = id.String()
	}

	return strings.Join(idStringList, constants.ListDataStringSeparator)
}
func (listID listID) Bytes() []byte {
	var byteList []byte

	for _, id := range listID.IDList.GetList() {
		byteList = append(byteList, id.Bytes()...)
	}

	return byteList
}
func (listID listID) Compare(compareID types.ID) int {
	return bytes.Compare(listID.Bytes(), compareID.Bytes())
}
func (listID listID) GetList() []traits.Listable {
	// TODO implement me
	panic("implement me")
}
func (listID listID) Size() int {
	// TODO implement me
	panic("implement me")
}
func (listID listID) Search(listable traits.Listable) int {
	// TODO implement me
	panic("implement me")
}
func (listID listID) Add(listableList ...traits.Listable) types.List {
	// TODO implement me
	panic("implement me")
}
func (listID listID) Remove(listableList ...traits.Listable) types.List {
	// TODO implement me
	panic("implement me")
}
