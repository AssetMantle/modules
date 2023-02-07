// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"reflect"
	"testing"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
)

func Test_auxiliaryResponse_GetError(t *testing.T) {
	type fields struct {
		Success bool
		Error   error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{true, nil}, false},
		{"+ve", fields{false, errorConstants.EntityNotFound}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryResponse := auxiliaryResponse{
				Success: tt.fields.Success,
				Error:   tt.fields.Error,
			}
			if err := auxiliaryResponse.GetError(); (err != nil) != tt.wantErr {
				t.Errorf("GetError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_auxiliaryResponse_IsSuccessful(t *testing.T) {
	type fields struct {
		Success bool
		Error   error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{true, nil}, true},
		{"+ve", fields{false, errorConstants.EntityNotFound}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryResponse := auxiliaryResponse{
				Success: tt.fields.Success,
				Error:   tt.fields.Error,
			}
			if got := auxiliaryResponse.IsSuccessful(); got != tt.want {
				t.Errorf("IsSuccessful() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newAuxiliaryResponse(t *testing.T) {
	type args struct {
		error error
	}
	tests := []struct {
		name string
		args args
		want helpers.AuxiliaryResponse
	}{
		{"+ve", args{nil}, auxiliaryResponse{true, nil}},
		{"-ve", args{errorConstants.EntityNotFound}, auxiliaryResponse{false, errorConstants.EntityNotFound}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newAuxiliaryResponse(tt.args.error); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAuxiliaryResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
