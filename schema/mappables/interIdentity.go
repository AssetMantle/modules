/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappables

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type InterIdentity interface {
	GetAuthentication() types.Property
	GetExpiry() types.Property

	Document
}
