// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package docs

import (
	"github.com/AssetMantle/modules/helpers"
)

type response struct {
	Success bool  `json:"success"`
	Error   error `json:"error"`
	ID      string
}

var _ helpers.Response = response{}

func (response response) IsSuccessful() bool {
	return response.Success
}
func (response response) GetError() error {
	return response.Error
}
func newResponse(id string, error error) helpers.Response {
	success := true
	if error != nil {
		success = false
	}

	return response{
		Success: success,
		Error:   error,
		ID:      id,
	}
}
