/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import "github.com/persistenceOne/persistenceSDK/schema/helpers"

type queries struct {
	queriesList []helpers.Query
}

var _ helpers.Queries = (*queries)(nil)

func (queries queries) Get(name string) helpers.Query {
	for _, query := range queries.GetList() {
		if query.GetName() == name {
			return query
		}
	}

	return nil
}

func (queries queries) GetList() []helpers.Query {
	return queries.queriesList
}

func NewQueries(queriesList ...helpers.Query) helpers.Queries {
	return queries{
		queriesList: queriesList,
	}
}
