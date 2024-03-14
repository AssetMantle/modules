// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/helpers"
)

type auxiliaries struct {
	auxiliaryList []helpers.Auxiliary
}

var _ helpers.Auxiliaries = (*auxiliaries)(nil)

func (auxiliaries auxiliaries) GetAuxiliary(name string) helpers.Auxiliary {
	for _, auxiliary := range auxiliaries.auxiliaryList {
		if auxiliary.GetName() == name {
			return auxiliary
		}
	}

	return nil
}
func (auxiliaries auxiliaries) Get() []helpers.Auxiliary {
	return auxiliaries.auxiliaryList
}

func NewAuxiliaries(auxiliaryList ...helpers.Auxiliary) helpers.Auxiliaries {
	for i, auxiliary := range auxiliaryList {

		if auxiliary == nil {
			panic("nil auxiliary")
		}

		for j, checkAuxiliary := range auxiliaryList {
			if i != j && auxiliary.GetName() == checkAuxiliary.GetName() {
				panic("repeated auxiliary")
			}
		}
	}

	return auxiliaries{
		auxiliaryList: auxiliaryList,
	}
}
