/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package module

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/keys"
	"github.com/persistenceOne/persistenceSDK/constants/names"
)

var (
	OrderExpired = errors.Register(constants.ProjectRoute+"/"+names.Orders, uint32(keys.Orders), "OrderExpired")
)
