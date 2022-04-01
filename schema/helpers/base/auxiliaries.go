// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import "github.com/persistenceOne/persistenceSDK/schema/helpers"

type auxiliaries struct {
	auxiliaryList []helpers.Auxiliary
}

var _ helpers.Auxiliaries = (*auxiliaries)(nil)

func (auxiliaries auxiliaries) Get(name string) helpers.Auxiliary {
	for _, auxiliary := range auxiliaries.auxiliaryList {
		if auxiliary.GetName() == name {
			return auxiliary
		}
	}

	return nil
}
func (auxiliaries auxiliaries) GetList() []helpers.Auxiliary {
	return auxiliaries.auxiliaryList
}

func NewAuxiliaries(auxiliaryList ...helpers.Auxiliary) helpers.Auxiliaries {
	return auxiliaries{
		auxiliaryList: auxiliaryList,
	}
}
