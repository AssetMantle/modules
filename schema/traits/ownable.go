/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package traits

import "github.com/persistenceOne/persistenceSDK/schema/types"

type Ownable interface {
	GetOwnerID() types.ID
	GetOwnableID() types.ID
}
