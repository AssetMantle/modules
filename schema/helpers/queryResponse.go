// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

type QueryResponse interface {
	Response
	Encode() ([]byte, error)
	Decode([]byte) (QueryResponse, error)
}
