/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

type QueryResponse interface {
	Response
	LegacyAminoEncode() ([]byte, error)
	LegacyAminoDecode([]byte) (QueryResponse, error)
}
