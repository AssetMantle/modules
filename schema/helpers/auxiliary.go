/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

type Auxiliary interface {
	GetName() string
	GetKeeper() AuxiliaryKeeper
	InitializeKeeper(Mapper, ...interface{})
}
