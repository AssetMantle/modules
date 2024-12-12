// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import "github.com/AssetMantle/modules/helpers"

type queries struct {
	queriesList []helpers.Query
}

var _ helpers.Queries = (*queries)(nil)

func (queries queries) GetQuery(servicePath string) helpers.Query {
	for _, query := range queries.Get() {
		if query.GetServicePath() == servicePath {
			return query
		}
	}

	return nil
}

func (queries queries) Get() []helpers.Query {
	return queries.queriesList
}

func NewQueries(queriesList ...helpers.Query) helpers.Queries {
	return queries{
		queriesList: queriesList,
	}
}
