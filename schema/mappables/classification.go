// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappables

import "github.com/persistenceOne/persistenceSDK/schema/helpers"

type Classification interface {
	Document
	helpers.Mappable
}
