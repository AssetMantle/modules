/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package revoke

import (
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Transaction = base.NewTransaction(
	"revoke",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	flags.FromID,
	flags.ToID,
	flags.ClassificationID,
	flags.MaintainedProperties,
	flags.AddMaintainer,
	flags.RemoveMaintainer,
	flags.MutateMaintainer,
)
