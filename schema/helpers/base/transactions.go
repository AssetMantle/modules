/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import "github.com/persistenceOne/persistenceSDK/schema/helpers"

type transactions struct {
	route           string
	transactionList []helpers.Transaction
}

var _ helpers.Transactions = (*transactions)(nil)

func (transactions transactions) GetRoute() string {
	return transactions.route
}

func (transactions transactions) Get(name string) helpers.Transaction {
	for _, transaction := range transactions.transactionList {
		if transaction.GetName() == name {
			return transaction
		}
	}
	return nil
}

func (transactions transactions) GetList() []helpers.Transaction {
	return transactions.transactionList
}

func NewTransactions(route string, transactionList ...helpers.Transaction) helpers.Transactions {
	return transactions{
		route:           route,
		transactionList: transactionList,
	}
}
