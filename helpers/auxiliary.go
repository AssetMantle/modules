// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

type Auxiliary interface {
	GetName() string
	GetKeeper() AuxiliaryKeeper
	Initialize(Mapper, ParameterManager, ...interface{}) Auxiliary
}
