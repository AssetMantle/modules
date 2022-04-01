// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"strings"

	"github.com/persistenceOne/persistenceSDK/constants"

	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type listID struct {
	IDList []types.ID `json:"idList"`
}

var _ types.ID = (*listID)(nil)
var _ types.List = (*listID)()

func (listID listID) String() string {
	idStringList := make([]string, len(listID.IDList))

	for i, id := range listID.IDList {
		idStringList[i] = id.String()
	}

	return strings.Join(idStringList, constants.ListDataStringSeparator)
}
func (listID listID) Bytes() []byte {
	var byteList []byte

	for _, id := range listID.IDList {
		byteList = append(byteList, id.Bytes()...)
	}

	return byteList
}
func (listID listID) Compare(compareID types.ID) int {
	return bytes.Compare(listID.Bytes(), compareID.Bytes())
}

func (listID listID) GetList() []interface{} {
	// TODO
}
func (listID listID) Search(f func()) int {
	// TODO
}
func (listID listID) Apply(f func()) types.List {
	// TODO
}
func (listID listID) Add(i ...interface{}) types.List {
	// TODO
}
func (listID listID) Remove(i ...interface{}) types.List {
	// TODO
}
func (listID listID) Mutate(i ...interface{}) types.List {
	// TODO
}
