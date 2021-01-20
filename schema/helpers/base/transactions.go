/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import "github.com/persistenceOne/persistenceSDK/schema/helpers"

type transactions struct {
	transactionList []helpers.Transaction
}

var _ helpers.Transactions = (*transactions)(nil)

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

func NewTransactions(transactionList ...helpers.Transaction) helpers.Transactions {
	return transactions{
		transactionList: transactionList,
	}
}
